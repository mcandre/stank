// Package stank provides primitives for analyzing programming scripts,
// especially shell scripts.
package stank

import (
	"mvdan.cc/sh/v3/syntax"

	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// Ignores is a poor man's gitignore.
//
// TODO: https://github.com/mcandre/stank/issues/1
var Ignores = []string{
	".git",
	".venv",
	"node_modules",
	"vendor",
}

// Ignore is a poor man's gitignore.
//
// TODO: https://github.com/mcandre/stank/issues/1
func Ignore(pth string) bool {
	for _, part := range strings.Split(pth, string(os.PathSeparator)) {
		for _, ignore := range Ignores {
			if part == ignore {
				return true
			}
		}
	}

	return false
}

// LOWEREXTENSIONS2POSIXyNESS is a fairly exhaustive map of lowercase file extensions to whether or not they represent POSIX shell scripts.
// Newly minted extensions can be added by stank contributors.
var LOWEREXTENSIONS2POSIXyNESS = map[string]bool{
	".ada":          false,
	".ash":          true,
	".bash":         true,
	".bash4":        true,
	".bash_login":   true,
	".bash_logout":  true,
	".bash_profile": true,
	".bashrc":       true,
	".bat":          false,
	".bin":          false,
	".bmp":          false,
	".bosh":         true,
	".c":            false,
	".cl":           false,
	".cmd":          false,
	".conf":         false,
	".csh":          false,
	".cshrc":        false,
	".dash":         true,
	".doc":          false,
	".docx":         false,
	".ds_store":     false,
	".e":            false,
	".elv":          false,
	".erl":          false,
	".escript":      false,
	".etsh":         false,
	".exe":          false,
	".expect":       false,
	".fish":         false,
	".flv":          false,
	".fth":          false,
	".gif":          false,
	".gitignore":    false,
	".gitkeep":      false,
	".gitmodules":   false,
	".groovy":       false,
	".hsh":          true,
	".ionrc":        false,
	".j":            false,
	".jpeg":         false,
	".jpg":          false,
	".js":           false,
	".json":         false,
	".ksh":          true,
	".ksh88":        true,
	".ksh93":        true,
	".kshrc":        true,
	".lisp":         false,
	".lksh":         false,
	".log":          false,
	".lua":          false,
	".markdown":     false,
	".md":           false,
	".mf":           false,
	".mksh":         true,
	".mov":          false,
	".mp3":          false,
	".mp4":          false,
	".oksh":         true,
	".pdf":          false,
	".pdksh":        true,
	".php":          false,
	".pike":         false,
	".pl":           false,
	".png":          false,
	".posh":         true,
	".properties":   false,
	".psh":          false,
	".py":           false,
	".pyw":          false,
	".rb":           false,
	".rc":           false,
	".rksh":         true,
	".rkt":          false,
	".scala":        false,
	".sf":           false,
	".sh":           true,
	".shinit":       true,
	".shrc":         true,
	".svg":          false,
	".swp":          false,
	".tcsh":         false,
	".tcshrc":       false,
	".tiff":         false,
	".tsh":          false,
	".txr":          false,
	".txt":          false,
	".vbs":          false,
	".wav":          false,
	".xml":          false,
	".yaml":         false,
	".yash":         true,
	".yml":          false,
	".zkl":          false,
	".zlogin":       true,
	".zlogout":      true,
	".zprofile":     true,
	".zsh":          true,
	".zshenv":       true,
	".zshrc":        true,
}

// LOWEREXTENSIONS2CONFIG is a fairly exhaustive map of lowercase file extensions to whether or not they represent shell script configurations.
// Newly minted extensions can be added by stank contributors.
var LOWEREXTENSIONS2CONFIG = map[string]bool{
	".ashrc":        true,
	".bash_login":   true,
	".bash_logout":  true,
	".bash_profile": true,
	".bashrc":       true,
	".cshrc":        true,
	".dashrc":       true,
	".fishrc":       true,
	".ionrc":        true,
	".kshrc":        true,
	".profile":      true,
	".rcrc":         true,
	".shinit":       true,
	".shrc":         true,
	".tcshrc":       true,
	".zlogin":       true,
	".zlogout":      true,
	".zprofile":     true,
	".zshenv":       true,
	".zshrc":        true,
}

// LOWERFILENAMES2POSIXyNESS is a fairly exhaustive map of lowercase filenames to whether or not they represent POSIX shell scripts.
// Newly minted config filenames can be added by stank contributors.
var LOWERFILENAMES2POSIXyNESS = map[string]bool{
	".profile":    true,
	"bash_login":  true,
	"bash_logout": true,
	"changelog":   false,
	"csh.login":   false,
	"csh.logout":  false,
	"login":       true,
	"logout":      true,
	"makefile":    false,
	"oilrc":       true,
	"oshrc":       true,
	"profile":     true,
	"rc.elv":      false,
	"rcrc":        false,
	"readme":      false,
	"shinit":      true,
	"shrc":        true,
	"tcsh.login":  false,
	"tcsh.logout": false,
	"thumbs.db":   false,
	"yshrc":       false,
	"zlogin":      true,
	"zlogout":     true,
	"zprofile":    true,
	"zshenv":      true,
	"zshrc":       true,
}

// LOWERFILENAMES2CONFIG is a fairly exhaustive map of lowercase filenames to whether or not they represent shell script configurations.
// Newly minted config filenames can be added by stank contributors.
var LOWERFILENAMES2CONFIG = map[string]bool{
	"bash_login":  true,
	"bash_logout": true,
	"csh.login":   true,
	"csh.logout":  true,
	"login":       true,
	"logout":      true,
	"oilrc":       true,
	"oshrc":       true,
	"profile":     true,
	"rc.elv":      true,
	"rcrc":        true,
	"shinit":      true,
	"shrc":        true,
	"tcsh.login":  true,
	"tcsh.logout": true,
	"yshrc":       true,
	"zlogin":      true,
	"zlogout":     true,
	"zprofile":    true,
	"zshenv":      true,
	"zshrc":       true,
}

// LOWEREXTENSIONS2INTERPRETER is a fairly exhaustive map of lowercase file extensions to their corresponding interpreters.
// Newly minted config extensions can be added by stank contributors.
var LOWEREXTENSIONS2INTERPRETER = map[string]string{
	".ashrc":        "ash",
	".awk":          "awk",
	".bash":         "bash",
	".bash_login":   "bash",
	".bash_logout":  "bash",
	".bash_profile": "bash",
	".bashrc":       "bash",
	".bsdmakefile":  "bmake",
	".csh":          "csh",
	".cshrc":        "csh",
	".dash":         "dash",
	".dashrc":       "dash",
	".elv":          "elvish",
	".fish":         "fish",
	".fishrc":       "fish",
	".gawk":         "gawk",
	".gnumakefile":  "gmake",
	".hsh":          "hsh",
	".ion":          "ion",
	".ionrc":        "ion",
	".ksh":          "ksh",
	".ksh88":        "ksh",
	".ksh93":        "ksh93",
	".ksh93rc":      "ksh93",
	".kshrc":        "ksh",
	".lkshrc":       "lksh",
	".lua":          "lua",
	".makefile":     "make",
	".mf":           "make",
	".mksh":         "mksh",
	".mkshrc":       "mksh",
	".osh":          "osh",
	".pdksh":        "pdksh",
	".pdkshrc":      "pdksh",
	".php":          "php",
	".pmakefile":    "pmake",
	".poshrc":       "posh",
	".profile":      "sh",
	".rc":           "rc",
	".rcrc":         "rc",
	".sed":          "sed",
	".sh":           "sh",
	".shinit":       "sh",
	".shrc":         "sh",
	".tcsh":         "tcsh",
	".tcshrc":       "tcsh",
	".ysh":          "ysh",
	".zlogin":       "zsh",
	".zlogout":      "zsh",
	".zprofile":     "zsh",
	".zsh":          "zsh",
	".zshenv":       "zsh",
	".zshprofile":   "zsh",
	".zshrc":        "zsh",
	"ash":           "ash",
}

// LOWERFILENAMES2INTERPRETER is a fairly exhaustive map of lowercase filenames to their corresponding interpreters.
// Newly minted config filenames can be added by stank contributors.
var LOWERFILENAMES2INTERPRETER = map[string]string{
	".ashrc":      "ash",
	".bashrc":     "bash",
	".cshrc":      "csh",
	".dashrc":     "dash",
	".fishrc":     "fish",
	".ionrc":      "ion",
	".ksh93rc":    "ksh93",
	".kshrc":      "ksh",
	".lkshrc":     "lksh",
	".login":      "sh",
	".logout":     "sh",
	".mkshrc":     "mksh",
	".pdkshrc":    "pdksh",
	".poshrc":     "posh",
	".rcrc":       "rc",
	".shinit":     "sh",
	".shrc":       "sh",
	".tcshrc":     "tcsh",
	".zlogin":     "zsh",
	".zlogout":    "zsh",
	".zprofile":   "zsh",
	".zshenv":     "zsh",
	".zshrc":      "zsh",
	"bsdmakefile": "bmake",
	"csh.login":   "csh",
	"csh.logout":  "csh",
	"gnumakefile": "gmake",
	"makefile":    "make",
	"oilrc":       "osh",
	"oshrc":       "osh",
	"pmakefile":   "pmake",
	"profile":     "sh",
	"rc.elv":      "elvish",
	"tcsh.login":  "tcsh",
	"tcsh.logout": "tcsh",
	"yshrc":       "ysh",
	"zlogin":      "zsh",
	"zlogout":     "zsh",
	"zprofile":    "zsh",
	"zshenv":      "zsh",
	"zshrc":       "zsh",
}

// BOMS acts as a registry set of known Byte Order mark sequences.
// See https://en.wikipedia.org/wiki/Byte_order_mark for more information.
var BOMS = [][]byte{
	{0x00, 0x00, 0xFE, 0xFF},
	{0x2B, 0x2F, 0x76, 0x2B},
	{0x2B, 0x2F, 0x76, 0x2F},
	{0x2B, 0x2F, 0x76, 0x38},
	{0x2B, 0x2F, 0x76, 0x39},
	{0xFE, 0xFF},
	{0xFF, 0xBB, 0xBF},
	{0xFF, 0xFE},
	{0xFF, 0xFE, 0x00, 0x00},
	{0x0E, 0xFE, 0xFF},
	{0x2B, 0x2F, 0x76, 0x38, 0x3D},
	{0x84, 0x31, 0x95, 0x33},
	{0xDD, 0x73, 0x66, 0x73},
	{0xF7, 0x64, 0x4C},
	{0xFB, 0xEE, 0x28},
}

// IsBOM checks whether a byte sequence is a BOM.
func IsBOM(bs []byte) bool {
	for _, bom := range BOMS {
		if bytes.Equal(bs, bom) {
			return true
		}
	}

	return false
}

// INTERPRETERS2POSIXyNESS is a fairly exhaustive map of interpreters to whether or not the interpreter is a POSIX compatible shell.
// Newly minted interpreters can be added by stank contributors.
var INTERPRETERS2POSIXyNESS = map[string]bool{
	"ash":    true,
	"awk":    false,
	"bash":   true,
	"bash4":  true,
	"bosh":   true,
	"csh":    false,
	"dash":   true,
	"elvish": false,
	"etsh":   false,
	"expect": false,
	"fish":   false,
	"gawk":   false,
	"hsh":    true,
	"ion":    false,
	"jruby":  false,
	"jython": false,
	"ksh":    true,
	"ksh88":  true,
	"ksh93":  true,
	"lksh":   false,
	"lua":    false,
	"mksh":   true,
	"node":   false,
	"oksh":   true,
	"oil":    true,
	"osh":    true,
	"pdksh":  true,
	"perl":   false,
	"perl6":  false,
	"php":    false,
	"posh":   true,
	"python": false,
	"rc":     false,
	"rksh":   true,
	"ruby":   false,
	"sed":    false,
	"sh":     true,
	"stash":  false,
	"swift":  false,
	"tclsh":  false,
	"tcsh":   false,
	"tsh":    false,
	"yash":   true,
	"ysh":    false,
	"zsh":    true,
}

// FullBashInterpreters note when a shell has the basic modern bash features,
// as opposed to subsets such as ash, dash, posh, ksh, zsh.
var FullBashInterpreters = map[string]bool{
	"bash":  true,
	"bash4": true,
}

// KshInterpreters note when a shell is a member of the modern ksh family.
var KshInterpreters = map[string]bool{
	"ksh":   true,
	"ksh88": true,
	"ksh93": true,
	"mksh":  true,
	"oksh":  true,
	"pdksh": true,
	"rksh":  true,
}

// SniffConfig bundles together the various options when sniffing files for POSIXyNESS.
type SniffConfig struct {
	EOLCheck bool
	CRCheck  bool
}

// ALTINTERPRETERS collects some alternative shell interpreters.
var ALTINTERPRETERS = map[string]bool{
	"csh":    true,
	"elvish": true,
	"etsh":   true,
	"fish":   true,
	"ion":    true,
	"lksh":   true,
	"rc":     true,
	"tcsh":   true,
	"tsh":    true,
}

// ALTEXTENSIONS collects some alternative shell script file extensions.
var ALTEXTENSIONS = map[string]bool{
	".csh":    true,
	".cshrc":  true,
	".elv":    true,
	".etsh":   true,
	".fish":   true,
	".fishrc": true,
	".ion":    true,
	".ionrc":  true,
	".lksh":   true,
	".rc":     true,
	".rcrc":   true,
	".tcsh":   true,
	".tcshrc": true,
	".tsh":    true,
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

// POSIXShCheckSyntax validates syntax for strict POSIX sh compliance.
func POSIXShCheckSyntax(smell Smell) error {
	parser := syntax.NewParser(syntax.Variant(syntax.LangPOSIX))

	fd, err := os.Open(smell.Path)

	if err != nil {
		return err
	}

	_, err = parser.Parse(bufio.NewReader(fd), smell.Path)
	return err
}

// UnixCheckSyntax validates syntax for the wider UNIX shell family.
func UnixCheckSyntax(smell Smell) error {
	cmd := exec.Command(smell.Interpreter, "-n", smell.Path)
	return cmd.Run()
}

// PerlishCheckSyntax validates syntax for Perl, Ruby, and Node.js.
func PerlishCheckSyntax(smell Smell) error {
	cmd := exec.Command(smell.Interpreter, "-c", smell.Path)
	return cmd.Run()
}

// PHPCheckSyntax validates syntax for PHP.
func PHPCheckSyntax(smell Smell) error {
	cmd := exec.Command(smell.Interpreter, "-l", smell.Path)
	return cmd.Run()
}

// PythonCheckSyntax validates syntax for Python.
func PythonCheckSyntax(smell Smell) error {
	cmd := exec.Command(smell.Interpreter, "-m", "py_compile", smell.Path)
	return cmd.Run()
}

// GoCheckSyntax validates syntax for Go.
func GoCheckSyntax(smell Smell) error {
	cmd := exec.Command("gofmt", "-e", smell.Path)
	return cmd.Run()
}

// GNUAwkCheckSyntax validates syntax for GNU awk files.
func GNUAwkCheckSyntax(smell Smell) error {
	cmd := exec.Command(smell.Interpreter, "--lint", "-f", smell.Path)
	return cmd.Run()
}

// Interpreter2SyntaxValidator provides syntax validator delegates, if one is available.
var Interpreter2SyntaxValidator = map[string]func(Smell) error{
	"ash":        UnixCheckSyntax,
	"bash":       UnixCheckSyntax,
	"bash4":      UnixCheckSyntax,
	"bmake":      UnixCheckSyntax,
	"bosh":       UnixCheckSyntax,
	"csh":        UnixCheckSyntax,
	"dash":       UnixCheckSyntax,
	"elvish":     UnixCheckSyntax,
	"fish":       UnixCheckSyntax,
	"gawk":       GNUAwkCheckSyntax,
	"generic-sh": POSIXShCheckSyntax,
	"gmake":      UnixCheckSyntax,
	"go":         GoCheckSyntax,
	"iojs":       PerlishCheckSyntax,
	"ksh":        UnixCheckSyntax,
	"ksh88":      UnixCheckSyntax,
	"ksh93":      UnixCheckSyntax,
	"lksh":       UnixCheckSyntax,
	"make":       UnixCheckSyntax,
	"mksh":       UnixCheckSyntax,
	"node":       PerlishCheckSyntax,
	"oksh":       UnixCheckSyntax,
	"oil":        UnixCheckSyntax,
	"osh":        UnixCheckSyntax,
	"pdksh":      UnixCheckSyntax,
	"perl":       PerlishCheckSyntax,
	"perl6":      PerlishCheckSyntax,
	"php":        PHPCheckSyntax,
	"pmake":      UnixCheckSyntax,
	"posh":       UnixCheckSyntax,
	"python":     PythonCheckSyntax,
	"python3":    PythonCheckSyntax,
	"rc":         UnixCheckSyntax,
	"rksh":       UnixCheckSyntax,
	"ruby":       PerlishCheckSyntax,
	"sh":         POSIXShCheckSyntax,
	"tcsh":       UnixCheckSyntax,
	"yash":       UnixCheckSyntax,
	"ysh":        UnixCheckSyntax,
	"zsh":        UnixCheckSyntax,
}

// LOWERMACHINEEXTENSIONS collects a rather truncated survey of
// machine-generated file extensions likely to not be edited directly
// by most shell script authors.
var LOWERMACHINEEXTENSIONS = map[string]bool{
	".sample": true, // git hooks, which violate every known shell script virtue
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
	fi, err := os.Lstat(pth)

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

	if _, extensionMachineOK := LOWERMACHINEEXTENSIONS[smell.Extension]; extensionMachineOK {
		smell.MachineGenerated = true
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

	smell.Symlink = fi.Mode()&os.ModeSymlink != 0

	if smell.Symlink {
		return smell, nil
	}

	extensionInterpreter, extensionInterpreterOK := LOWEREXTENSIONS2INTERPRETER[strings.ToLower(smell.Extension)]

	if extensionInterpreterOK {
		smell.Interpreter = extensionInterpreter
	}

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
		if IsBOM(bs[:i]) {
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
		if smell.POSIXy && !extensionInterpreterOK {
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

	filenameInterpreter, filenameInterpreterOK := LOWERFILENAMES2INTERPRETER[strings.ToLower(interpreterFilename)]

	// Identify the interpreter, or mark as generic, unknown sh interpreter.
	if interpreterFilename == "" {
		if filenameInterpreterOK {
			smell.Interpreter = filenameInterpreter
		} else if !extensionInterpreterOK {
			smell.Interpreter = "generic-sh"
		}
	} else {
		smell.Interpreter = interpreterFilename
		smell.InterpreterFlags = commandParts[1:]
	}

	smell.Bash = FullBashInterpreters[smell.Interpreter]
	smell.Ksh = KshInterpreters[smell.Interpreter]

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
