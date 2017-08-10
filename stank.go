package stank

import (
	"bufio"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Version is semver.
const Version = "0.0.3"

// Smell describes the overall impression of a file's POSIXyness,
// using several factors to determine with a reasonably high accuracy
// whether or not the file is a POSIX compatible shell script.
//
// An idiomatic shebang preferably leads the file, such as #!/bin/bash, #!/bin/zsh, #!/bin/sh, etc.
// This represents good form when writing shell scripts, in particular ensuring that
// the script will be evaluated by the right interpreter, even if the extension is omitted or a generic ".sh".
// Shell scripts, whether executable applications or source'able libraries, should include a shebang.
// One attribute not analyzed by this library is unix file permission bits. Application shell scripts should set the executable bit(s) to 1, while shell scripts intended to be sourced or imported should not set these bits. Either way, the bits have hardly any correlation with the POSIXyness of a file, as the false positives and false negatives are too frequent.
//
// Common filenames for POSIX compatible scripts include .profile, .login, .bashrc, .bash_profile,
// .zshrc, .kshrc, .envrc*, and names for git hooks. The stank library will catalog some of these
// standard names, though application-specific filenames are various and sundry. Ultimately,
// all files containing POSIX compatible shell content should include a shebang, to help
// interpreters, editors, and linters identify POSIX shell content.
//
// File extension is another way to estimate a script's POSIXyness. For example,
// ".bash", ".ksh", ".posh", ".sh", etc. would each indicate a POSIX compatible shell script,
// whereas ".py", ".pl", ".rb", ".csh", ".rc", and so on would indicate nonPOSIX script.
// File extensions are often omitted or set to a generic ".sh" for command line applications,
// in which case the extension is insufficient for establishing the POSIX vs. nonPOSIX nature
// of the script. This is why shebangs are so important; while file extensions can be helpful,
// shell scripts really rely moreso on the shebang for self identification, and extensions aren't always
// desirable, as unix CLI applications prefer to omit the extension from the filename for brevity.
// Note that some filenames such as ".profile" may be logically considered to have basename "" (blank) and extension ".profile", or basename ".profile" with extension ".profile", or else basename ".profile" and extension "" (blank). In practice, Go treats both the basename and extension for these kinds of files as containing ".profile", and Smell will behave accordingly.
//
// File encoding also sensitive for shell scripts. Generally, ASCII subset is recommended for maximum portability.
// If your terminal supports it, the LANG environment variable can be altered to accept UTF-8 and other encodings,
// enabling raw UTF-8 data to be used in script contents. However, this restricts your scripts to running only
// on systems explicitly configured to match the encoding/locale of your script; and tends to furter limit
// the platforms for your script to specifically GNU libc Linux distributions, so using nonASCII content in your scripts
// is inadvisable. Shell scripts conforming to POSIX should really use pure ASCII characters.
// NonUTF-8 encodings such as UTF-16, UTF-32, and even nonUnicode encodings like EBCDIC, Latin1, and KOI8-R
// usually indicate a nonPOSIX shell script, even a localization file or other nonscript. These encodings
// are encountered less often than ASCII and UTF-8, and are generally considered legacy formats.
// For performance reasons, the stank library will not attempt to discern the exact encoding of a file, but merely report whether the file leads with a byte order marker such as 0xEFBBBF (UTF-8) or 0xFEFF (UTF-16, UTF-32). If BOM, then the file is Unicode, which may lead to a stank warning, as POSIX shell scripts are best written in pure ASCII, for maximum cross-platform compatibliity. BOMs are outside of the 127 max integer range for ASCII values, so a file with a BOM is likely not a POSIX shell script, while a file without a BOM may be a POSIX shell script.
//
// Line endings for POSIX shell scripts should LF="\n" in C-style notation. Alternative line endings such as CRLF="\r\n",
// ancient Macintosh CR="\r", and bizarre forms like vertical tab (ASCII code 0x0B) or form feed (ASCII code 0x0C)
// are possible in a fuzzing sense, but may lead to undefined behavior depending on the particular shell interpreter.
// For the purposes of identifying POSIX vs nonPOSIX scripts, a Smell will look for LF, CRLF, and CR; and ignore the presence
// or absence of these other exotic whitespace separators.
// NonPOSIX scripts written in Windows, such as Python and Ruby scripts, are ideally written with LF line endings,
// though it is common to observe CRLF endings, as Windows users more frequently invoke these as "python script.py",
// "ruby script.rb", rather than the bare "script" or dot slash "./script" forms typically used by unix administrators.
// For performance, the stank library will not report possible multiple line ending styles, such as poorly formatted text files featuring both CRLF and LF line endings. The library will simply report the first confirmed line ending style.
//
// Moreover, POSIX line ending LF is expected at the end of a text file, so a final end of line character "\n" is good form.
// Common unix utilities such as cat expect this final EOL, and will misrender the successive shell prompt when processing
// files that omit the final EOL. Make expects a final EOL, and gcc may produce malformed .c code if the .h header files neglect
// to include a final EOL.
// For performance reasons, the stank library will not attempt to read the entire file to
// report on the presence/absence of a final EOL. Shell script authors should nonetheless configure their text editors
// to consistently include a final EOL in the vast majority of text file formats.
//
// A POSIXy flag indicates that, to the best of the stank library's ability,
// a file is identified as either very likely a POSIX shell script, or something else.
// Something else encompasses nonPOSIX shell scripts such as Csh, Tcsh, Python, Ruby, Lua scripts;
// also encompasses nonscript files such as multimedia images, audio, rich text documents,
// machine code, and other nonUTF-8, nonASCII content.
type Smell struct {
	Path        string
	Filename    string
	Basename    string
	Extension   string
	BOM         bool
	Shebang     string
	Interpreter string
	LineEnding  string
	POSIXy      bool
}

// LOWEREXTENSIONS2POSIXyNESS is a fairly exhaustive map of lowercase file extensions to whether or not they represent POSIX shell scripts.
// Newly minted extensions can be added by stank contributors.
var LOWEREXTENSIONS2POSIXyNESS = map[string]bool{
	".sh":         true,
	".bash":       true,
	".zsh":        true,
	".ksh":        true,
	".pdksh":      true,
	".ksh93":      true,
	".mksh":       true,
	".dash":       true,
	".posh":       true,
	".ash":        true,
	".shrc":       true,
	".bashrc":     true,
	".kshrc":      true,
	".zshenv":     true,
	".zprofile":   true,
	".zshrc":      true,
	".zlogin":     true,
	".zlogout":    true,
	".csh":        false,
	".cshrc":      false,
	".tcsh":       false,
	".tcshrc":     false,
	".fish":       false,
	".rc":         false,
	".ionrc":      false,
	".py":         false,
	".pyw":        false,
	".pl":         false,
	".rb":         false,
	".lua":        false,
	".js":         false,
	".mf":         false,
	".exe":        false,
	".bin":        false,
	".cmd":        false,
	".bat":        false,
	".psh":        false,
	".vbs":        false,
	".txt":        false,
	".md":         false,
	".markdown":   false,
	".doc":        false,
	".docx":       false,
	".pdf":        false,
	".log":        false,
	".gitignore":  false,
	".gitmodules": false,
	".gitkeep":    false,
	".xml":        false,
	".json":       false,
	".yml":        false,
	".yaml":       false,
	".conf":       false,
	".properties": false,
	".svg":        false,
	".gif":        false,
	".jpg":        false,
	".jpeg":       false,
	".png":        false,
	".bmp":        false,
	".tiff":       false,
	".mp3":        false,
	".wav":        false,
	".mp4":        false,
	".mov":        false,
	".flv":        false,
	".swp":        false,
	".ds_store":   false,
}

// LOWERFILENAMES2POSIXyNESS is a fairly exhaustive map of lowercase filenames to whether or not they represent POSIX shell scripts.
// Newly minted config filenames can be added by stank contributors.
var LOWERFILENAMES2POSIXyNESS = map[string]bool{
	"profile":       true,
	".profile":      true,
	".login":        true,
	".logout":       true,
	".bash_profile": true,
	".bash_login":   true,
	".bash_logout":  true,
	"zshenv":        true,
	"zprofile":      true,
	"zshrc":         true,
	"zlogin":        true,
	"zlogout":       true,
	"csh.login":     false,
	"csh.logout":    false,
	"tcsh.login":    false,
	"tcsh.logout":   false,
	"makefile":      false,
	"readme":        false,
	"changelog":     false,
	"thumbs.db":     false,
}

// BOMS acts as a registry set of known Byte Order mark sequences.
// See https://en.wikipedia.org/wiki/Byte_order_mark for more information.
var BOMS = map[string]bool{
	"\uFFBBBF":   true,
	"\uFEFF":     true,
	"\uFFFE":     true,
	"\u0000FEFF": true,
	"\uFFFE0000": true,
	"\u2B2F7638": true,
	"\u2B2F7639": true,
	"\u2B2F762B": true,
	"\u2B2F762F": true,
	// {byte(0x2B), byte(0x2F), byte(0x76), byte(0x38), byte(0x3D)}: true,
	// {byte(0xF7), byte(0x64), byte(0x4C)}:                         true,
	// {byte(0xDD), byte(0x73), byte(0x66), byte(0x73)}:             true,
	// {byte(0x0E), byte(0xFE), byte(0xFF)}:                         true,
	// {byte(0xFB), byte(0xEE), byte(0x28)}:                         true,
	// {byte(0x84), byte(0x31), byte(0x95), byte(0x33)}:             true,
}

// INTERPRETERS2POSIXyNESS is a fairly exhaustivemap of interpreters to whether or not the interpreter is a POSIX compatible shell.
// Newly minted interpreters can be added by stank contributors.
var INTERPRETERS2POSIXyNESS = map[string]bool{
	"sh":     true,
	"bash":   true,
	"zsh":    true,
	"ksh":    true,
	"pdksh":  true,
	"ksh93":  true,
	"mksh":   true,
	"dash":   true,
	"posh":   true,
	"ash":    true,
	"csh":    false,
	"tcsh":   false,
	"fish":   false,
	"rc":     false,
	"python": false,
	"jython": false,
	"perl":   false,
	"perl6":  false,
	"ruby":   false,
	"jruby":  false,
	"lua":    false,
	"node":   false,
	"awk":    false,
	"sed":    false,
	"swift":  false,
	"tclsh":  false,
	"ion":    false,
}

// Sniff analyzes the holistic smell of a given file path,
// returning a Smell record of key indicators tending towards either POSIX compliance or noncompliance,
// including a flag for the final "POSIXy" trace scent of the file.
//
// For performance, if the scent of one or more attributes obviously indicates POSIX or nonPOSIX,
// Sniff() may short-circuit, setting the POSIXy flag and returning a record
// with some attributes set to zero value.
//
// Polyglot and multiline shebangs are technically possible in languages that do not support native POSIX-style shebang comments ( see https://rosettacode.org/wiki/Multiline_shebang ). However, Sniff() can reliably identify only ^#!.+$ POSIX-style shebangs, and will populate the Shebang field accordingly.
//
// If an I/O problem occurs during analysis, an error value will be set.
// Otherwise, the error value will be nil.
func Sniff(pth string) (Smell, error) {
	// Attempt to short-circuit for directories
	fi, err := os.Stat(pth)

	smell := Smell{Path: pth}

	if err != nil {
		return smell, err
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return smell, nil
	}

	smell.Filename = path.Base(pth)
	smell.Basename = filepath.Base(smell.Filename)
	smell.Extension = filepath.Ext(smell.Filename)

	// Attempt to short-circuit for Emacs swap files
	if strings.HasSuffix(smell.Filename, "~") {
		return smell, nil
	}

	// Attempt to short-circuit by extension
	if posixy, ok := LOWEREXTENSIONS2POSIXyNESS[strings.ToLower(smell.Extension)]; ok {
		smell.POSIXy = posixy
		return smell, nil
	}

	// Attempt to short-circuit by filename
	if posixy, ok := LOWERFILENAMES2POSIXyNESS[strings.ToLower(smell.Filename)]; ok {
		smell.POSIXy = posixy
		return smell, nil
	}

	fd, err := os.Open(pth)

	if err != nil {
		return smell, err
	}

	defer func() {
		err := fd.Close()

		if err != nil {
			log.Panic(err)
		}
	}()

	// Check for BOMs
	br := bufio.NewReader(fd)

	bs, err := br.Peek(5)

	if err != nil {
		return smell, err
	}

	for i := 2; i < 6; i++ {
		if BOMS[string(bs[:i])] {
			smell.BOM = true
			break
		}
	}

	if smell.BOM {
		return smell, nil
	}

	LF := byte('\n')

	// Attempt to find the first occurence of a line feed.
	// CR-ended files and binary files will be read in their entirety.
	line, err := br.ReadString(LF)

	// An error occurred while attempting to find the first occurence of a line feed in the file.
	// This could mean one of several things:
	//
	// * The connection to the file was lost (network disruption, file movement, file deletion, etc.)
	// * The file is completely empty.
	// * The file is binary.
	// * The file is CR-ended.
	// * The file consists of a single line, without a line ending sequence.
	//
	// Only the cases of an empty file or single line without an ending could reasonably considered candidates for POSIX shell scripts. The former can only be evidenced as POSIX if a POSIXy extension is present, in which case the previous analysis instructions above would have short-circuited POSIXy: true. So we can now ignore the former and only check the latter.

	if strings.HasSuffix(line, "\r\n") {
		smell.LineEnding = "\r\n"
	} else if strings.HasSuffix(line, "\n") {
		smell.LineEnding = "\n"
	}

	// Absent a shebang, at this point we have no evidence at all
	// that the file would be a POSIX shell script. We could almost
	// run statistical or grammatical analysis on the remainder of
	// the file contents (a la GitHub Linguist), but that kind of
	// processing is slow, unreliable, and there appears to be no
	// Go library for this, but merely external Ruby processes.
	// If your shell script omits both an extension and a shebang,
	// you fucked up. If your shebang line begins !#, you fucked up.
	// In short, shebang lines are the primary mechanism for operating
	// with shell scripts, and the above analysis is merely a polite
	// undertaking to account for minor infractions in POSIX shell scripts,
	// quick tests for obvious, but honestly secondary, signs of POSIXyness.
	if !strings.HasPrefix(line, "#!") {
		return smell, err
	}

	smell.Shebang = strings.TrimRight(line, "\r\n")

	// shebang minus the #! prefix.
	command := smell.Shebang[2:]

	// At this point, we have a script that is not obviously filenamed either a POSIX shell script file, nor obviously a nonPOSIX file. We have read the first line of the file, and determined that it is some sort of POSIX-style shebang.
	// Example commonly encountered shebang forms:
	//
	// * #!/bin/bash
	// * #!/usr/local/bin/bash
	// * #!/usr/bin/env python
	// * #!/usr/bin/env MathKernel -script
	// * #!someapplication
	//
	// Let's break these down.
	//
	// #!/bin/someinterpreter is the idiomatic way to shebang most POSIX shell scripts, especially those depending on very standard, established shells like bash, zsh, ksh, and so on, that are expected to be installed in /bin.
	// #!/usr/local/bin/bash is acceptable for interpreters installed in custom locations, such as macOS users using Homebrew to provide bash v4 in /usr/local/bin.
	// #!/usr/bin/env python is preferred for general purpose scripting languages like Python, Perl, Ruby, and Lua, that are installed somewhere on the system, but not necessarily in /bin on all systems. For example, rvm may place ruby in $HOME/.rvm/rubies/ruby-$RUBY_VERSION/bin. So the /usr/bin/env command prefix helps these languages interoperate with POSIX sh standards, allowing the interpreter to be used in the shebang without hardcoding any particular absolute path to the interpreter; the interpreter simply needs to be available somewhere in $PATH. When identifying the interpreter, We will need to be careful to strip out /usr/bin/env, if present.
	// #!/usr/bin/env MathKernel -script and #!/bin/bash -euo pipefail constitute shebangs with flags to be passed to the interpreters. When identifying the interpreter, We will need to be careful to strip out flags meant for the interpreter, if present.
	//
	// Finally, #!bash, #!fish, #!python, etc. are technically allowed, though some systems may balk on the interpreter being relative to $PATH rather than an absolute file path. This form is no problem for identifying the stinky interpreter for our purposes, but the stank linter may emit a warning to use the more idiomatic shebangs #!/bin/bash, #!/usr/bin/env fish, #!/usr/bin/env python, etc.

	commandParts := strings.Split(command, " ")

	// Strip /usr/bin/env, if present
	if commandParts[0] == "/usr/bin/env" {
		commandParts = commandParts[1:]
	}

	// Identify the interpreter, or blank
	smell.Interpreter = commandParts[0]

	// Strip out directory path, if any
	interpreterFilename := filepath.Base(smell.Interpreter)

	// Compare interpreter against common POSIX and nonPOSIX names
	smell.POSIXy = INTERPRETERS2POSIXyNESS[interpreterFilename]

	return smell, nil
}
