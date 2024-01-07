package stank

import (
	"os"
)

// Smell describes the overall impression of a file's POSIXyness,
// using several factors to determine with a reasonably high accuracy
// whether or not the file is a POSIX compatible shell script.
//
// An idiomatic shebang preferably leads the file, such as #!/bin/bash, #!/bin/zsh, #!/bin/sh, etc.
// This represents good form when writing shell scripts, in particular ensuring that
// the script will be evaluated by the right interpreter, even if the extension is omitted or a generic ".sh".
// Shell scripts, whether executable applications or source'able libraries, should include a shebang.
// One attribute not analyzed by this library is unix file permission bits. Application shell scripts should set the executable bit(s) to 1,
// while shell scripts intended to be sourced or imported should not set these bits.
// Either way, the bits have hardly any correlation with the POSIXyness of a file, as the false positives and false negatives are too frequent.
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
// Note that some filenames such as ".profile" may be logically considered to have basename "" (blank) and extension ".profile",
// or basename ".profile" with extension ".profile", or else basename ".profile" and extension "" (blank).
// In practice, Go treats both the basename and extension for these kinds of files as containing ".profile", and Smell will behave accordingly.
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
// For performance reasons, the stank library will not attempt to discern the exact encoding of a file,
// but merely report whether the file leads with a byte order marker such as 0xEFBBBF (UTF-8) or 0xFEFF (UTF-16, UTF-32).
// If BOM, then the file is Unicode, which may lead to a stank warning, as POSIX shell scripts are best written in pure ASCII,
// for maximum cross-platform compatibliity. BOMs are outside of the 127 max integer range for ASCII values,
// so a file with a BOM is likely not a POSIX shell script, while a file without a BOM may be a POSIX shell script.
//
// Line endings for POSIX shell scripts should LF="\n" in C-style notation. Alternative line endings such as CRLF="\r\n",
// ancient Macintosh CR="\r", and bizarre forms like vertical tab (ASCII code 0x0B) or form feed (ASCII code 0x0C)
// are possible in a fuzzing sense, but may lead to undefined behavior depending on the particular shell interpreter.
// For the purposes of identifying POSIX vs nonPOSIX scripts, a Smell will look for LF, CRLF, and CR; and ignore the presence
// or absence of these other exotic whitespace separators.
// NonPOSIX scripts written in Windows, such as Python and Ruby scripts, are ideally written with LF line endings,
// though it is common to observe CRLF endings, as Windows users more frequently invoke these as "python script.py",
// "ruby script.rb", rather than the bare "script" or dot slash "./script" forms typically used by unix administrators.
// For performance, the stank library will not report possible multiple line ending styles,
// such as poorly formatted text files featuring both CRLF and LF line endings.
// The library will simply report the first confirmed line ending style.
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
//
// Scripts referencing "sh" are generally considered to be POSIX sh. Ignoring unmarked legacy Thompson sh scripts.
//
// Unknown, even more obscure languages are assumed to be non-POSIXY.
//
// Languages with duplicate names (e.g. oil shell osh vs. OpenSolaris oil shell) are generally assumed not to be POSIXy.
// Unable to disambiguate without more specific information (shebang names, file extentions).
type Smell struct {
	Path              string
	Filename          string
	Basename          string
	Extension         string
	Symlink           bool
	Shebang           string
	Interpreter       string
	InterpreterFlags  []string
	LineEnding        string
	FinalEOL          *bool
	ContainsCR        bool
	Permissions       os.FileMode
	Directory         bool
	OwnerExecutable   bool
	Library           bool
	BOM               bool
	POSIXy            bool
	Bash              bool
	Ksh               bool
	AltShellScript    bool
	CoreConfiguration bool
	MachineGenerated  bool
}
