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
const Version = "0.0.14"

// LOWEREXTENSIONS2POSIXyNESS is a fairly exhaustive map of lowercase file extensions to whether or not they represent POSIX shell scripts.
// Newly minted extensions can be added by stank contributors.
var LOWEREXTENSIONS2POSIXyNESS = map[string]bool{
	".sh":           true,
	".tsh":          false,
	".etsh":         false,
	".bash":         true,
	".bash4":        true,
	".bosh":         true,
	".yash":         true,
	".zsh":          true,
	".lksh":         false,
	".ksh":          true,
	".ksh88":        true,
	".pdksh":        true,
	".ksh93":        true,
	".mksh":         true,
	".oksh":         true,
	".rksh":         true,
	".dash":         true,
	".posh":         true,
	".ash":          true,
	".shrc":         true,
	".shinit":       true,
	".bash_profile": true,
	".bashrc":       true,
	".bash_login":   true,
	".bash_logout":  true,
	".kshrc":        true,
	".zshenv":       true,
	".zprofile":     true,
	".zshrc":        true,
	".zlogin":       true,
	".zlogout":      true,
	".csh":          false,
	".cshrc":        false,
	".tcsh":         false,
	".tcshrc":       false,
	".fish":         false,
	".rc":           false,
	".ionrc":        false,
	".expect":       false,
	".py":           false,
	".pyw":          false,
	".pl":           false,
	".rb":           false,
	".php":          false,
	".lua":          false,
	".js":           false,
	".lisp":         false,
	".mf":           false,
	".exe":          false,
	".bin":          false,
	".cmd":          false,
	".bat":          false,
	".psh":          false,
	".vbs":          false,
	".ada":          false,
	".c":            false,
	".cl":           false,
	".e":            false,
	".erl":          false,
	".escript":      false,
	".fth":          false,
	".groovy":       false,
	".j":            false,
	".pike":         false,
	".rkt":          false,
	".scala":        false,
	".elv":          false,
	".sf":           false,
	".txr":          false,
	".zkl":          false,
	".txt":          false,
	".md":           false,
	".markdown":     false,
	".doc":          false,
	".docx":         false,
	".pdf":          false,
	".log":          false,
	".gitignore":    false,
	".gitmodules":   false,
	".gitkeep":      false,
	".xml":          false,
	".json":         false,
	".yml":          false,
	".yaml":         false,
	".conf":         false,
	".properties":   false,
	".svg":          false,
	".gif":          false,
	".jpg":          false,
	".jpeg":         false,
	".png":          false,
	".bmp":          false,
	".tiff":         false,
	".mp3":          false,
	".wav":          false,
	".mp4":          false,
	".mov":          false,
	".flv":          false,
	".swp":          false,
	".ds_store":     false,
}

// LOWEREXTENSIONS2CONFIG is a fairly exhaustive map of lowercase file extensions to whether or not they represent shell script configurations.
// Newly minted extensions can be added by stank contributors.
var LOWEREXTENSIONS2CONFIG = map[string]bool{
	".shrc":         true,
	".shinit":       true,
	".profile":      true,
	".bash_profile": true,
	".bashrc":       true,
	".bash_login":   true,
	".bash_logout":  true,
	".ashrc":        true,
	".dashrc":       true,
	".kshrc":        true,
	".zshenv":       true,
	".zprofile":     true,
	".zshrc":        true,
	".zlogin":       true,
	".zlogout":      true,
	".cshrc":        true,
	".tcshrc":       true,
	".fishrc":       true,
	".rcrc":         true,
	".ionrc":        true,
}

// LOWERFILENAMES2POSIXyNESS is a fairly exhaustive map of lowercase filenames to whether or not they represent POSIX shell scripts.
// Newly minted config filenames can be added by stank contributors.
var LOWERFILENAMES2POSIXyNESS = map[string]bool{
	"shrc":                      true,
	"shinit":                    true,
	".profile":                  true,
	"profile":                   true,
	"login":                     true,
	"logout":                    true,
	"bash_login":                true,
	"bash_logout":               true,
	"zshenv":                    true,
	"zprofile":                  true,
	"zshrc":                     true,
	"zlogin":                    true,
	"zlogout":                   true,
	"csh.login":                 false,
	"csh.logout":                false,
	"tcsh.login":                false,
	"tcsh.logout":               false,
	"rcrc":                      false,
	"makefile":                  false,
	"readme":                    false,
	"changelog":                 false,
	"applypatch-msg.sample":     false,
	"commit-msg.sample":         false,
	"post-update.sample":        false,
	"pre-applypatch.sample":     false,
	"pre-commit.sample":         false,
	"pre-push.sample":           false,
	"pre-rebase.sample":         false,
	"pre-receive.sample":        false,
	"prepare-commit-msg.sample": false,
	"update.sample":             false,
	"rc.elv":                    false,
	"thumbs.db":                 false,
}

// LOWERFILENAMES2CONFIG is a fairly exhaustive map of lowercase filenames to whether or not they represent shell script configurations.
// Newly minted config filenames can be added by stank contributors.
var LOWERFILENAMES2CONFIG = map[string]bool{
	"shrc":        true,
	"shinit":      true,
	"profile":     true,
	"login":       true,
	"logout":      true,
	"bash_login":  true,
	"bash_logout": true,
	"zshenv":      true,
	"zprofile":    true,
	"zshrc":       true,
	"zlogin":      true,
	"zlogout":     true,
	"csh.login":   true,
	"csh.logout":  true,
	"tcsh.login":  true,
	"tcsh.logout": true,
	"rcrc":        true,
	"rc.elv":      true,
}

// LOWEREXTENSIONS2INTERPRETER is a fairly exhaustive map of lowercase file extensions to their corresponding interpreters.
// Newly minted config extensions can be added by stank contributors.
var LOWEREXTENSIONS2INTERPRETER = map[string]string{
	".shrc":         "sh",
	".shinit":       "sh",
	".bashrc":       "bash",
	".zshrc":        "zsh",
	".lkshrc":       "lksh",
	".kshrc":        "ksh",
	".pdkshrc":      "pdksh",
	".ksh93rc":      "ksh93",
	".mkshrc":       "mksh",
	".dashrc":       "dash",
	".poshrc":       "posh",
	".ashrc":        "ash",
	".zshenv":       "zsh",
	".zprofile":     "zsh",
	".zlogin":       "zsh",
	".zlogout":      "zsh",
	".cshrc":        "csh",
	".tcshrc":       "tcsh",
	".fishrc":       "fish",
	".rcrc":         "rc",
	".ionrc":        "ion",
	".profile":      "sh",
	".bash_profile": "bash",
	".bash_login":   "bash",
	".bash_logout":  "bash",
	".zshprofile":   "zsh",
	".elv":          "elvish",
	".php":          "php",
	".lua":          "lua",
}

// LOWERFILENAMES2INTERPRETER is a fairly exhaustive map of lowercase filenames to their corresponding interpreters.
// Newly minted config filenames can be added by stank contributors.
var LOWERFILENAMES2INTERPRETER = map[string]string{
	".shrc":       "sh",
	".shinit":     "sh",
	".bashrc":     "bash",
	".zshrc":      "zsh",
	".lkshrc":     "lksh",
	".kshrc":      "ksh",
	".pdkshrc":    "pdksh",
	".ksh93rc":    "ksh93",
	".mkshrc":     "mksh",
	".dashrc":     "dash",
	".poshrc":     "posh",
	".ashrc":      "ash",
	".zshenv":     "zsh",
	".zprofile":   "zsh",
	".zlogin":     "zsh",
	".zlogout":    "zsh",
	".cshrc":      "csh",
	".tcshrc":     "tcsh",
	".fishrc":     "fish",
	".rcrc":       "rc",
	".ionrc":      "ion",
	"profile":     "sh",
	".login":      "sh",
	".logout":     "sh",
	"zshenv":      "zsh",
	"zprofile":    "zsh",
	"zshrc":       "zsh",
	"zlogin":      "zsh",
	"zlogout":     "zsh",
	"csh.login":   "csh",
	"csh.logout":  "csh",
	"tcsh.login":  "tcsh",
	"tcsh.logout": "tcsh",
	"rc.elv":      "elvish",
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

// INTERPRETERS2POSIXyNESS is a fairly exhaustive map of interpreters to whether or not the interpreter is a POSIX compatible shell.
// Newly minted interpreters can be added by stank contributors.
var INTERPRETERS2POSIXyNESS = map[string]bool{
	"sh":     true,
	"tsh":    false,
	"etsh":   false,
	"bash":   true,
	"bash4":  true,
	"bosh":   true,
	"yash":   true,
	"zsh":    true,
	"lksh":   false,
	"ksh":    true,
	"ksh88":  true,
	"pdksh":  true,
	"ksh93":  true,
	"mksh":   true,
	"oksh":   true,
	"rksh":   true,
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
	"php":    false,
	"lua":    false,
	"node":   false,
	"awk":    false,
	"sed":    false,
	"swift":  false,
	"tclsh":  false,
	"ion":    false,
	"elvish": false,
	"expect": false,
	"stash":  false,
}

// FULL_BASH_INTERPRETERS note when a shell has the basic modern bash features,
// as opposed to subsets such as ash, dash, posh, ksh, zsh.
var FULL_BASH_INTERPRETERS = map[string]bool {
	"bash": true,
	"bash4": true,
}

// KSH_INTERPRETERS note when a shell is a member of the modern ksh family.
var KSH_INTERPRETERS = map[string]bool {
	"ksh":          true,
	"ksh88":        true,
	"pdksh":        true,
	"ksh93":        true,
	"mksh":         true,
	"oksh":         true,
	"rksh":         true,
}

// SniffConfig bundles together the various options when sniffing files for POSIXyNESS.
type SniffConfig struct {
	EOLCheck bool
	CRCheck  bool
}

// ALTINTERPRETERS collects some alternative shell interpreters.
var ALTINTERPRETERS = map[string]bool{
	"osh":    true,
	"lksh":   true,
	"csh":    true,
	"tcsh":   true,
	"fish":   true,
	"ion":    true,
	"rc":     true,
	"tsh":    true,
	"etsh":   true,
	"elvish": true,
}

// ALTEXTENSIONS collets some alternative shell script file extensions.
var ALTEXTENSIONS = map[string]bool{
	".osh":    true,
	".lksh":   true,
	".csh":    true,
	".cshrc":  true,
	".tcsh":   true,
	".tcshrc": true,
	".fish":   true,
	".fishrc": true,
	".ion":    true,
	".ionrc":  true,
	".rc":     true,
	".rcrc":   true,
	".tsh":    true,
	".etsh":   true,
	".elv":    true,
}

// ALTFILENAMES matches some alternative shell script profile filenames.
var ALTFILENAMES = map[string]bool{
	"csh.login":  true,
	"csh.logout": true,
	"rc.elv":     true,
}

// IsAltShellScript returns whether a smell represents a non-POSIX, but nonetheless similar kind of lowlevel shell script language.
func IsAltShellScript(smell Smell) bool {
	return ALTINTERPRETERS[smell.Interpreter] || ALTEXTENSIONS[smell.Extension] || ALTFILENAMES[smell.Filename]
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
func Sniff(pth string, config SniffConfig) (Smell, error) {
	// Attempt to short-circuit for directories
	fi, err := os.Stat(pth)

	smell := Smell{Path: pth}

	if err != nil {
		return smell, err
	}

	mode := fi.Mode()

	if mode.IsDir() {
		smell.Directory = true
		return smell, nil
	}

	smell.Permissions = mode.Perm()
	smell.OwnerExecutable = smell.Permissions&0100 != 0
	smell.Filename = path.Base(pth)
	smell.Basename = filepath.Base(smell.Filename)
	smell.Extension = filepath.Ext(smell.Filename)

	// Attempt to short-circuit for Emacs swap files
	if strings.HasSuffix(smell.Filename, "~") {
		return smell, nil
	}

	extensionPOSIXy, extensionPOSIXyOK := LOWEREXTENSIONS2POSIXyNESS[strings.ToLower(smell.Extension)]

	if extensionPOSIXyOK {
		smell.POSIXy = extensionPOSIXy
	}

	filenamePOSIXy, filenamePOSIXyOK := LOWERFILENAMES2POSIXyNESS[strings.ToLower(smell.Filename)]

	if filenamePOSIXyOK {
		smell.POSIXy = filenamePOSIXy
	}

	smell.CoreConfiguration = LOWEREXTENSIONS2CONFIG[strings.ToLower(smell.Extension)] ||
		LOWERFILENAMES2CONFIG[strings.ToLower(smell.Filename)]

	smell.Library = (smell.CoreConfiguration || smell.Extension != "") && !smell.OwnerExecutable

	fd, err := os.Open(pth)

	if err != nil {
		return smell, err
	}

	defer func() {
		err = fd.Close()

		if err != nil {
			log.Panic(err)
		}
	}()

	//
	// Check for BOMs
	//

	br := bufio.NewReader(fd)

	maxBOMCheckLength := 5

	if fi.Size() < 5 {
		maxBOMCheckLength = int(fi.Size())
	}

	bs, err := br.Peek(maxBOMCheckLength)

	if err != nil {
		return smell, err
	}

	for i := 2; i < 6 && i < maxBOMCheckLength; i++ {
		if BOMS[string(bs[:i])] {
			smell.BOM = true

			if _, err = br.Discard(i); err != nil {
				return smell, err
			}

			break
		}
	}

	LF := byte('\n')

	// Attempt to find the first occurence of a line feed.
	// CR-ended files and binary files will be read in their entirety.
	line, err := br.ReadString(LF)

	if err != nil {
		return smell, err
	}

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
	//
	// Note that stank currently ignores mixed line ending styles within a file.
	//

	if strings.HasSuffix(line, "\r\n") {
		smell.LineEnding = "\r\n"
	} else if strings.HasSuffix(line, "\n") {
		smell.LineEnding = "\n"
	} else if strings.HasSuffix(line, "\r") {
		smell.LineEnding = "\r"
	}

	filenameInterpreter, filenameInterpreterOK := LOWERFILENAMES2INTERPRETER[strings.ToLower(smell.Filename)]

	if filenameInterpreterOK {
		smell.Interpreter = filenameInterpreter
	}

	//
	// Read the entire script in order to assess the presence/absence of a final POSIX end of line (\n) sequence.
	//
	if config.EOLCheck && fi.Size() > 0 {
		fd2, err := os.Open(pth)

		if err != nil {
			log.Print(err)
			return smell, nil
		}

		defer func() {
			err := fd2.Close()

			if err != nil {
				log.Panic(err)
			}
		}()

		maxEOLSequenceLength := int64(2)

		if fi.Size() < 2 {
			maxEOLSequenceLength = 1
		}

		eolBuf := make([]byte, maxEOLSequenceLength)

		if _, err := fd2.ReadAt(eolBuf, fi.Size()-maxEOLSequenceLength); err != nil {
			return smell, err
		}

		if eolBuf[maxEOLSequenceLength-1] == byte('\n') && (maxEOLSequenceLength < 2 || eolBuf[0] != byte('\r')) {
			b := true
			smell.FinalEOL = &b
		}
	}

	// Recognize poorly written shell scripts that feature
	// a POSIXy filename but lack a proper shebang line.
	if !strings.HasPrefix(line, "#!") && !strings.HasPrefix(line, "!#") {
		if smell.POSIXy && !filenameInterpreterOK {
			smell.Interpreter = "generic-sh"
		}

		return smell, nil
	}

	smell.Shebang = strings.TrimRight(line, "\r\n")

	// shebang minus the #! prefix.
	command := strings.TrimSpace(smell.Shebang[2:])

	// At this point, we have a script that is not obviously filenamed either a POSIX shell script file, nor obviously a nonPOSIX file. We have read the first line of the file, and determined that it is some sort of POSIX-style shebang.
	// Example commonly encountered shebang forms:
	//
	// * #!/bin/bash
	// * #!/usr/local/bin/bash
	// * #!/usr/bin/env python
	// * #!/usr/bin/env MathKernel -script
	// * #!/bin/busybox python
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

	// Strip /bin/busybox, if present
	if commandParts[0] == "/bin/busybox" {
		commandParts = commandParts[1:]
	}

	interpreterPath := commandParts[0]

	// Strip out directory path, if any
	interpreterFilename := filepath.Base(interpreterPath)

	extensionInterpreter, extensionInterpreterOK := LOWEREXTENSIONS2INTERPRETER[strings.ToLower(smell.Extension)]

	// Identify the interpreter, or mark as generic, unknown sh interpreter.
	if interpreterFilename == "" {
		if extensionInterpreterOK {
			smell.Interpreter = extensionInterpreter
		} else if filenameInterpreterOK {
			smell.Interpreter = filenameInterpreter
		} else {
			smell.Interpreter = "generic-sh"
		}
	} else {
		smell.Interpreter = interpreterFilename
		smell.InterpreterFlags = commandParts[1:]
	}

	smell.Bash = FULL_BASH_INTERPRETERS[smell.Interpreter]
	smell.Ksh = KSH_INTERPRETERS[smell.Interpreter]

	// Compare interpreter against common POSIX and nonPOSIX names.
	interpreterPOSIXy := INTERPRETERS2POSIXyNESS[interpreterFilename]

	if interpreterPOSIXy && (!extensionPOSIXyOK || extensionPOSIXy) && (!filenamePOSIXyOK || filenamePOSIXy) {
		smell.POSIXy = true
	} else if IsAltShellScript(smell) {
		smell.AltShellScript = true
	}

	if (smell.POSIXy || smell.AltShellScript) && config.CRCheck {
		fd3, err := os.Open(pth)

		defer func() {
			err = fd3.Close()

			if err != nil {
				log.Panic(err)
			}
		}()

		if err != nil {
			return smell, err
		}

		br2 := bufio.NewReader(fd3)

		CR := byte('\r')

		_, err = br2.ReadString(CR)

		smell.ContainsCR = err == nil
	}

	return smell, nil
}
