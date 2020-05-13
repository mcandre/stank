package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mcandre/stank"
	"mvdan.cc/sh/syntax"
)

var flagEOL = flag.Bool("eol", true, "Report presence/absence of final end of line sequence")
var flagCR = flag.Bool("cr", true, "Report presence/absence of final end of line sequence")
var flagModulino = flag.Bool("modulino", false, "Enforce strict separation of application scripts vs. library scripts")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// Funk holds configuration for a funky walk.
type Funk struct {
	EOLCheck      bool
	CRCheck       bool
	ModulinoCheck bool
	FoundOdor     bool
}

// CheckEOL analyzes POSIXy scripts for the presence/absence of a final end of line sequence such as \n at the end of a file, \r\n, etc.
func CheckEOL(smell stank.Smell) bool {
	if smell.FinalEOL != nil && !(*smell.FinalEOL) {
		fmt.Printf("Missing final end of line sequence: %s\n", smell.Path)
		return true
	}

	return false
}

// CheckCR analyzes POSIXy scripts for the presence/absence of a CR/CRLF line ending sequence.
func CheckCR(smell stank.Smell) bool {
	if smell.ContainsCR {
		fmt.Printf("CR/CRLF line ending detected: %s\n", smell.Path)
		return true
	}

	return false
}

// CheckBOMs analyzes POSIXy scripts for byte order markers. If a BOM is found, CheckBOMs prints a warning and returns true.
// Otherwise, CheckBOMs returns false.
func CheckBOMs(smell stank.Smell) bool {
	if smell.BOM {
		fmt.Printf("Leading BOM reduces portability: %s\n", smell.Path)

		return true
	}

	return false
}

// CheckShebangs analyzes POSIXy scripts for some shebang oddities. If an oddity is found, CheckShebangs prints a warning and returns true.
// Otherwise, CheckShebangs returns false.
//
// Note: While shell safety flags are risky when placed in shebangs,
// Unfortunately many non-POSIXy languages unfortunately require such flags:
// sed, awk, Emacs Lisp, Fourth, Octave, Mathematica, ...
// Therefore, CheckShebangs may trigger unactionable warnings when run on non-POSIXy files.
func CheckShebangs(smell stank.Smell) bool {
	// Shebangs are ill advised for configuration files.
	if smell.CoreConfiguration {
		if smell.Shebang != "" {
			fmt.Printf("Configuration features shebang: %s\n", smell.Path)
			return true
		}

		return false
	}

	if smell.Shebang == "" {
		fmt.Printf("Missing shebang: %s\n", smell.Path)
		return true
	}

	if !strings.HasPrefix(smell.Shebang, "#!") {
		fmt.Printf("Shebang appears to be flipped: %v\n", smell.Path)
		return true
	}

	if !strings.HasPrefix(smell.Shebang, "#!/") {
		fmt.Printf("Shebang application should be absolute and non-nested: %v\n", smell.Path)
		return true
	}

	if strings.Contains(smell.Shebang[2:], "#") {
		fmt.Printf("Commented shebangs may be unparsable: %v\n", smell.Path)
		return true
	}

	if len(smell.InterpreterFlags) != 0 {
		fmt.Printf("Risk of parse error for interpreter space / secondary argument. Any safety flags will be ignored on `%v <script>` launch: %v\n", smell.Interpreter, smell.Path)
		return true
	}

	return false
}

// CheckPermissions analyzes POSIXy scripts for some file permission oddities. If an oddity is found, CheckPermissions prints a warning and returns true.
// Otherwise, CheckPermissions returns false.
func CheckPermissions(smell stank.Smell) bool {
	if smell.Library && smell.Permissions&0111 != 0 {
		fmt.Printf("Sourceable script features executable mode bits: %s\n", smell.Path)
		return true
	}

	if (smell.Extension == "" && smell.Permissions&0100 == 0) ||
		(smell.Extension != "" && smell.Permissions&0111 != 0) {
		fmt.Printf("Ambiguous launch style. Either feature a file extensions, or else feature executable bits: %v\n", smell.Path)
		return true
	}

	return false
}

// CheckModulino warns when a smell features some aspects of an application, such as executable bits, and simultaneously some aspects of a library, such as a non-empty file extension.
// If the file is a pure application or library, CheckModulino returns false.
// Otherwise, CheckModulino returns true.
func CheckModulino(smell stank.Smell) bool {
	if stank.LOWEREXTENSIONS2CONFIG[strings.ToLower(smell.Extension)] || stank.LOWERFILENAMES2CONFIG[strings.ToLower(smell.Filename)] {
		return false
	}

	if (smell.Extension == "" && !smell.OwnerExecutable) || (smell.Extension != "" && (smell.Permissions&0100 != 0 || smell.Permissions&0010 != 0 || smell.Permissions&0001 != 0)) {
		fmt.Printf("Modulino ambiguity. Either have owner executable permissions with no extension, or else remove executable bits and use an extension like .lib.sh: %s\n", smell.Path)
		return true
	}

	return false
}

// CheckSlick passes any sh interpreted scripts through a strict POSIX sh parser.
func CheckSlick(smell stank.Smell) bool {
	if !smell.POSIXy || (smell.Interpreter != "generic-sh" && smell.Interpreter != "sh") {
		return false
	}

	parser := syntax.NewParser(syntax.Variant(syntax.LangPOSIX))

	fd, er := os.Open(smell.Path)

	if er != nil {
		fmt.Printf("%v\n", er)
		return true
	}

	br := bufio.NewReader(fd)
	_, er = parser.Parse(br, smell.Path)

	if er != nil {
		fmt.Printf("%v\n", er)
		return true
	}

	return false
}

// UnsetIFSPattern matches unset IFS commands.
var UnsetIFSPattern = regexp.MustCompile(`^(\s)*unset(\s)+IFS(\s+(#.*)?)?$`)

// CheckIFSReset enforces IFS configured to '\n\t ' near the beginning of executable scripts,
// in order to reduce tokenization errors.
func CheckIFSReset(smell stank.Smell) bool {
	if !smell.POSIXy || smell.Library {
		return false
	}

	fd, err := os.Open(smell.Path)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return true
	}

	defer func() {
		err = fd.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	}()

	fi, err := os.Lstat(smell.Path)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return true
	}

	if fi.Size() == 0 {
		return false
	}

	scanner := bufio.NewScanner(fd)

	var candidateLine string

	for scanner.Scan() {
		line := scanner.Text()

		if UnsetIFSPattern.MatchString(line) {
			return false
		}

		if strings.HasPrefix(line, "#") ||
			strings.HasPrefix(line, "set") ||
			strings.HasPrefix(line, "unset") ||
			strings.HasPrefix(line, "trap") ||
			strings.TrimSpace(line) == "" {
			continue
		}

		candidateLine = line
		break
	}

	if candidateLine == "" {
		return false
	}

	if index := strings.Index(candidateLine, "#"); index != -1 {
		candidateLine = candidateLine[:index]
	}

	candidateLine = strings.TrimSpace(candidateLine)

	assignmentParts := strings.Split(candidateLine, "=")

	if len(assignmentParts) < 1 || strings.TrimSpace(assignmentParts[0]) != "IFS" {
		fmt.Printf("Tokenize like `unset IFS` at the top of executable scripts: %v\n", smell.Path)
		return true
	}

	return false
}

// CheckSafetyFlags warns on missing `set`... safety command from the beginning of executable scripts,
// in order to reduce runtime errors.
func CheckSafetyFlags(smell stank.Smell) bool {
	if !smell.POSIXy || smell.Library {
		return false
	}

	fd, err := os.Open(smell.Path)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return true
	}

	defer func() {
		err = fd.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	}()

	fi, err := os.Lstat(smell.Path)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return true
	}

	if fi.Size() == 0 {
		return false
	}

	scanner := bufio.NewScanner(fd)

	var candidateLine string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "#") ||
			strings.HasPrefix(line, "IFS") ||
			strings.HasPrefix(line, "unset") ||
			strings.HasPrefix(line, "trap") ||
			strings.TrimSpace(line) == "" {
			continue
		}

		candidateLine = line
		break
	}

	if candidateLine == "" {
		return false
	}

	if index := strings.Index(candidateLine, "#"); index != -1 {
		candidateLine = candidateLine[:index]
	}

	candidateLine = strings.TrimSpace(candidateLine)

	parts := strings.Split(candidateLine, " ")

	if len(parts) < 1 || strings.TrimSpace(parts[0]) != "set" {
		fmt.Printf("Control program flow like `set -euf` at the top of executable scripts: %v\n", smell.Path)
		return true
	}

	return false
}

// FunkyCheck analyzes POSIXy scripts for some oddities. If an oddity is found, FunkyCheck prints a warning and returns true.
// Otherwise, FunkyCheck returns false.
func (o Funk) FunkyCheck(smell stank.Smell) bool {
	var resEOL bool

	if o.EOLCheck {
		resEOL = CheckEOL(smell)
	}

	var resCR bool

	if o.CRCheck {
		resCR = CheckCR(smell)
	}

	var resModulino bool

	if o.ModulinoCheck {
		resModulino = CheckModulino(smell)
	}

	resBOM := CheckBOMs(smell)
	resShebang := CheckShebangs(smell)
	resPerms := CheckPermissions(smell)
	resSlick := CheckSlick(smell)
	resIFSReset := CheckIFSReset(smell)
	resSafetyFlags := CheckSafetyFlags(smell)

	return resEOL ||
		resCR ||
		resBOM ||
		resModulino ||
		resShebang ||
		resPerms ||
		resSlick ||
		resIFSReset ||
		resSafetyFlags
}

// Ignores is a poor man's gitignore.
//
// TODO: https://github.com/mcandre/stank/issues/1
var Ignores = []string{
	".git",
	"vendor",
	"node_modules",
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

// Walk is a callback for filepath.Walk to lint shell scripts.
func (o *Funk) Walk(pth string, info os.FileInfo, err error) error {
	if Ignore(pth) {
		return nil
	}

	smell, err := stank.Sniff(pth, stank.SniffConfig{EOLCheck: o.EOLCheck, CRCheck: o.CRCheck})

	if err != nil && err != io.EOF {
		fmt.Printf("%v\n", err)
	}

	if (smell.POSIXy || smell.AltShellScript) && o.FunkyCheck(smell) {
		o.FoundOdor = true
	}

	return nil
}

func main() {
	flag.Parse()

	funk := Funk{}

	if *flagEOL {
		funk.EOLCheck = true
	}

	if *flagCR {
		funk.CRCheck = true
	}

	if *flagModulino {
		funk.ModulinoCheck = true
	}

	switch {
	case *flagVersion:
		fmt.Println(stank.Version)
		os.Exit(0)
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(0)
	}

	paths := flag.Args()

	for _, pth := range paths {
		filepath.Walk(pth, funk.Walk)
	}

	if funk.FoundOdor {
		os.Exit(1)
	}
}
