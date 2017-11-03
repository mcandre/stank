# stank: analyzers for determining whether files smell like rotten POSIX shell scripts, or faintly rosy like Ruby and Python scripts

# ABOUT

stank is a library and collection of command line utilities for sniffing files to identify shell scripts like bash, sh, zsh, ksh and so on, those funky farmfresh gobs of garbaggio; versus other more palatable files like rb, py, pl. Believe it or not, shell scripts are notoriously difficult to write well, so it behooves a developer to either write shell scripts in safer languages, or else wargame your scripts with an armada of linters. Trouble is, in large projects one can never be too sure which files are honest to dog POSIX compliant shell scripts, and which are pretenders. csh, tcsh, fish, ion, rc, and most other nonderivatives of bash tend to be NOT POSIX compatible. If you're geeky enough to have followed thus far, let's get crackalackin with some fruity examples dammit!

# EXAMPLES

The stank system includes the stank Go library as well as several command line utilities for convenience. The `stank` application scans directories and files for POSIX-derived shell scripts and prints their paths, designed as a convenient standalone filter for linting large collections of source code. For example, use `stank` in combination with `xargs` to help per-file shell linters process large projects.

```console
$ stank examples
examples/.profile
examples/.zshrc
examples/badconfigs/.bash_profile
examples/badconfigs/zprofile
examples/blank.bash
examples/derp.zsh
examples/globs.bash
examples/goodbye.sh
examples/greetings.bash
examples/hello
examples/hello.sh
examples/hooks/post-update
examples/hooks/pre-applypatch
examples/hooks/pre-commit
examples/hooks/pre-push
examples/hooks/pre-rebase
examples/hooks/update
examples/howdy
examples/howdy.zsh
examples/i-should-have-an-extension
examples/just-eol.bash
examples/just-shebang.bash
examples/pipefail
examples/salutations.bash
examples/salutations.sh
examples/salutations4.bash
examples/sample.envrc
examples/wednesday
examples/welcome
examples/welcome.sh

$ stank examples/hooks | xargs checkbashisms
error: examples/hooks/pre-rebase: Unterminated quoted string found, EOF reached. Wanted: <'>, opened in line 133

$ stank -help
  -alt
        Limit results to specifically alternative, non-POSIX lowlevel shell scripts
  -help
        Show usage information
  -sh
        Limit results to specifically bare POSIX sh scripts
  -version
        Show version information
```

`rosy` recommends scripts to be rewritten in other languages, such as porting bash scripts to ksh for speed; porting sh scripts to bash for robustness; or porting zsh scripts to sh for portability. By default, Rose mode is applied, encouraging shell scripts to be rewritten in non-shell languages for significant improvements in robustness and speed.

```console
$ rosy -kame examples
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/blank.bash
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/derp.zsh
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/globs.bash
Clarify interpreter with a shebang line: examples/goodbye.sh
Clarify interpreter with a shebang line: examples/greetings.bash
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/hello-legacy
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/hello.bosh
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/hello.lksh
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/hello.osh
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/hello.yash
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/howdy
Clarify interpreter with a shebang line: examples/howdy.zsh
Clarify interpreter with a shebang line: examples/just-eol.bash
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/just-shebang.bash
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/lo
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/lo-cr.csh
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/lo.csh
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/pipefail
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/salutations.bash
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/salutations.sh
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/salutations4.bash
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/wednesday
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/wednesday-bom
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/welcome
Rewrite script in sh, ksh, posh, dash, etc. for performance boost: examples/welcome.sh

$ rosy -help
  -ahiru
        Recommend sh for portability
  -help
        Show usage information
  -kame
        Recommend faster shells
  -usagi
        Recommend more robust shells
  -version
        Show version information
```

The `funk` linter reports strange odors emanating from scripts, such as improper line endings, the presence of Byte Order Marker's in some Unicode scripts.

```console
$ funk examples
Configuration features shebang: examples/badconfigs/.bash_profile
Configuration features executable permissions: examples/badconfigs/zprofile
Missing final end of line sequence: examples/blank.bash
Missing shebang: examples/blank.bash
Interpreter mismatch between shebang and extension: examples/derp.zsh
Missing shebang: examples/greetings.bash
Missing final end of line sequence: examples/hello-crlf.sh
CR/CRLF line ending detected: examples/hello-crlf.sh
Missing shebang: examples/howdy.zsh
Missing shebang: examples/just-eol.bash
Missing final end of line sequence: examples/lo-cr.csh
CR/CRLF line ending detected: examples/lo-cr.csh
Leading BOM reduces portability: examples/wednesday-bom

$ funk -modulino examples
Configuration features shebang: examples/badconfigs/.bash_profile
Configuration features executable permissions: examples/badconfigs/zprofile
Missing final end of line sequence: examples/blank.bash
Missing shebang: examples/blank.bash
Interpreter mismatch between shebang and extension: examples/derp.zsh
Missing shebang: examples/greetings.bash
Missing final end of line sequence: examples/hello-crlf.sh
CR/CRLF line ending detected: examples/hello-crlf.sh
Modulino ambiguity. Either have owner executable permissions with no extension, or else remove executable bits and use an extension like .lib.sh: examples/hello-crlf.sh
Modulino ambiguity. Either have owner executable permissions with no extension, or else remove executable bits and use an extension like .lib.sh: examples/howdy
Missing shebang: examples/howdy.zsh
Missing shebang: examples/just-eol.bash
Modulino ambiguity. Either have owner executable permissions with no extension, or else remove executable bits and use an extension like .lib.sh: examples/lo
Missing final end of line sequence: examples/lo-cr.csh
CR/CRLF line ending detected: examples/lo-cr.csh
Modulino ambiguity. Either have owner executable permissions with no extension, or else remove executable bits and use an extension like .lib.sh: examples/pipefail
Modulino ambiguity. Either have owner executable permissions with no extension, or else remove executable bits and use an extension like .lib.sh: examples/shout.sh
Modulino ambiguity. Either have owner executable permissions with no extension, or else remove executable bits and use an extension like .lib.sh: examples/wednesday
Modulino ambiguity. Either have owner executable permissions with no extension, or else remove executable bits and use an extension like .lib.sh: examples/wednesday-bom
Leading BOM reduces portability: examples/wednesday-bom
Modulino ambiguity. Either have owner executable permissions with no extension, or else remove executable bits and use an extension like .lib.sh: examples/welcome

$ funk -help
  -cr
        Report presence/absence of final end of line sequence (default true)
  -eol
        Report presence/absence of final end of line sequence (default true)
  -help
        Show usage information
  -modulino
        Enforce strict separation of application scripts vs. library scripts
  -version
        Show version information
```

Each of `stank`, `funk`, and `rosy` have the ability to select lowlevel, nonPOSIX scripts as well, such as csh/tcsh scripts used in FreeBSD.

Note that funk cannot reliably warn for missing shebangs if the extension is also missing; typically, script authors use one or the other to mark files as shell scripts. Lacking both a shebang and a file extension, means that a file could contain code for many languages, making it difficult to determine the POSIXy nature of the code. Even if an exhaustive set of ASTs are applied to test the file contents for syntactical validity across the dozens of available shell languages, there is a strong possibility in shorter files that the contents are merely incidentally valid script syntax, though the intent of the file is not to operate as a POSIX shell script. Short, nonPOSIX scripts such as for csh/tcsh could easily trigger a "POSIX" syntax match. In any case, know that the shebang is requisite for ensuring your scripts are properly interpreted.

Note that funk may fail to present permissions warnings if the scripts are housed on non*nix file systems such as NTFS, where executable bits are often missing from the file metadata altogether. When storing shell scripts, be sure to set the appropriate file permissions, and transfer files as a bundle in a tarball or similar to safeguard against dropped permissions.

Note that funk may warn of interpreter mismatches for scripts with extraneous dots in the filename. Rather than `.envrc.sample`, name the file `sample.envrc`. Rather than `wget-google.com`, name the file `wget-google-com`. Appending `.sh` is also an option, so `update.es.cluster` renames to `update.es.cluster.sh`.

The optional `-modulino` flag to funk enables strict separation of script duties, into distinct application scripts vs. library scripts. Application scripts are generally executed by invoking the path, such as `./hello` or `~/bin/hello` or simply `hello` when `$PATH` is appropriately modified. Application scripts feature owner executable permissions, and perhaps group and other as well depending on system configuration needs. In contrast, library scripts are intended to be imported with dot (`.`) or `source` into user shells or other scripts, and should feature a file extension like `.lib.sh`, `.sh`, `.bash`, etc. By using separate naming conventions, we more quickly communicate to downstream users how to interact with a shell script. In particular, by dropping file extensions for shell script applications, we encourage authors to choose more meaningful script names. Instead of the generic `build.sh`, choose `build-docker`. Instead of `kafka.sh`, choose `start-kafka`, `kafka-entrypoint`, etc.

Finally, `stink` prints a record of each file's POSIXyness, including any interesting fields it identified along the way. Note that some fields may be zero valued if the stench of POSIX or rosy waft of nonPOSIX is overwhelming, short-circuiting analysis. This short-circuiting feature dramatically speeds up how `stank` and `rosy` search large projects.

Note that permissions are relayed as decimals, due to constraints on JSON integer formatting (we didn't want to use a custom octal string field). Use `echo 'obase=8;<some integer> | bc` to display these values in octal.

Note that legacy systems, packages, and shell scripts referencing "sh" may refer to a plethora of pre-POSIX shells. Modern systems rename "sh" to "lksh", "tsh", "etsh", etc. to avoid confusion. In general, the stank suite will assume that the majority of scripts being scanned are targeting post-1971 technology, so use your human intuition and context to note any legacy Thompson UNIX v6 "sh", etc. scripts. Most modern linters will neither be able to parse such scripts of any complexity, nor will they recognize them for the legacy scripts that they are, unless the scripts' shebangs are rendered with the modern retro interpreters "lksh", "tsh", "etsh", etc. for deployment on modern *nix systems. One could almost use the fs stats for modification/change to try to identify these legacy outliers, but this is a practically unrealistic assumption except for the most obsessive archaeologist, diligently ensuring their legacy scripts continue to present 1970's metadata even after experimental content modifications. So the stank system will simply punt and assume sh -> POSIX sh, ksh -> ksh88 / ksh93 for the sake of modernity and balance.

Similarly, the old Bourne shell AKA "sh" AKA "bsh" presents language identification difficulties. Old Bourne shell scripts are most likely to present themselves with "sh" shebangs, which is okay as Bourne sh and ksh88/pdksh/ksh served as the bases for the POSIX sh standard. Some modern systems may present a Bourne shell as a "sh" or "bsh" binary. The former presents few problems for stank identification, though "bsh" is tricky, as the majority of its uses today are not associated with the Bourne shell but with the Java BeanShell. So stank may default to treating `bsh` scripts as non-POSIXy, and any such Bourne shell scripts are advised to feature either `bash` or `sh` shebangs, and perhaps `.sh` or `.bash` extensions, in order to self-identify as modern, POSIX compliant scripts.

```console
$ stink examples/hello
{"Path":"examples/hello","Filename":"hello","Basename":"hello","Extension":"","Shebang":"#!/bin/sh","Interpreter":"sh","LineEnding":"\n","FinalEOL":false,"ContainsCR":false
,"Permissions":509,"Directory":false,"OwnerExecutable":true,"BOM":false,"POSIXy":true,"AltShellScript":false}

$ stink -pp examples/hello
{
  "Path": "examples/hello",
  "Filename": "hello",
  "Basename": "hello",
  "Extension": "",
  "Shebang": "#!/bin/sh",
  "Interpreter": "sh",
  "LineEnding": "\n",
  "FinalEOL": false,
  "ContainsCR": false,
  "Permissions": 509,
  "Directory": false,
  "OwnerExecutable": true,
  "BOM": false,
  "POSIXy": true,
  "AltShellScript": false
}

$ stink -pp examples/hello.py
{
  "Path": "examples/hello.py",
  "Filename": "hello.py",
  "Basename": "hello.py",
  "Extension": ".py",
  "Shebang": "#!/usr/bin/env python",
  "Interpreter": "python",
  "LineEnding": "\n",
  "FinalEOL": false,
  "ContainsCR": false,
  "Permissions": 420,
  "Directory": false,
  "OwnerExecutable": false,
  "BOM": false,
  "POSIXy": false,
  "AltShellScript": false
}

$ stink -help
  -cr
        Report presence/absence of any CR/CRLF's
  -eol
        Report presence/absence of final end of line sequence
  -help
        Show usage information
  -pp
        Prettyprint smell records
  -version
        Show version information
```

The included `examples/` directory demonstrates many edge cases, such as empty scripts, shebang-less scripts, extensioned and extensionless scripts, and various Hello World applications in across many programming languages. Some files, such as `examples/goodbye` may contain 100% valid POSIX shell script content, but fail to self-identify with either shebangs or relevant file extensions. In a large project, such files may be mistakenly treated as whoknowswhat format, or simply plain text. Perhaps statistical methods could help identify POSIX grammars, but even an empty file is technically POSIX, which is unhelpful from a reliable classification standpoint. In any case, `examples/` hopefully covers the more common edge cases.

# DOWNLOADS

https://github.com/mcandre/stank/releases

# DOCUMENTATION

http://godoc.org/github.com/mcandre/stank

# REQUIREMENTS

Each of the applications in the stank suite are standalone, with no requirements other than deploying to a suitable operating system.

## Development

* [Go](https://golang.org) 1.7+
* [gox](https://github.com/mitchellh/gox)
* [zipc](https://github.com/mcandre/zipc)
* [coreutils](https://www.gnu.org/software/coreutils/coreutils.html)
* [findutils](https://www.gnu.org/software/findutils/)
* [make](https://www.gnu.org/software/make/)
* [golint](https://github.com/golang/lint)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)
* [errcheck](https://github.com/kisielk/errcheck)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [opennota/check](https://github.com/opennota/check) (e.g. `go get github.com/opennota/check/cmd/...`)
* [megacheck](https://github.com/dominikh/go-tools/tree/master/cmd/megacheck) (e.g. `go get github.com/dominikh/go-tools/cmd/megacheck`)
* [flcl](https://github.com/mcandre/flcl)
* [editorconfig-cli](https://github.com/amyboyd/editorconfig-cli)

# INSTALL FROM REMOTE GIT REPOSITORY

```console
$ go get github.com/mcandre/stank/...
```

# INSTALL FROM LOCAL GIT REPOSITORY

```console
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone https://github.com/mcandre/stank.git $GOPATH/src/github.com/stank
$ sh -c "cd $GOPATH/src/github.com/mcandre/stank/cmd/stink && go install"
$ sh -c "cd $GOPATH/src/github.com/mcandre/stank/cmd/stank && go install"
$ sh -c "cd $GOPATH/src/github.com/mcandre/stank/cmd/rosy && go install"
```

# WARNING ON FALSE POSITIVES

Some rather obscure files, such as Common Lisp source code with multiline, polyglot shebangs and no file extension, may falsely trigger the stank library, and the rosy, stink, and stank applications, which short-circuit on the first line of the hacky shebang. Such files may be falsely identified as "POSIX" code, which is actually the intended behavior! This is because the polyglot shebang is a hack to work around limitations in the Common Lisp language, which ordinarily does not accept POSIX shebang comments, in order to get Common Lisp scripts to be dot-slashable in bash. For this situation, it is best to supply a proper file extension to such files.

```console
$ head examples/i-should-have-an-extension
#!/usr/bin/env sh
#|
exec clisp -q -q $0 $0 ${1+"$@"}
|#

(defun hello-main (args)
  (format t "Hello from main!~%"))

;;; With help from Francois-Rene Rideau
;;; http://tinyurl.com/cli-args

$ stink -pp examples/i-should-have-an-extension
{
  "Path": "examples/i-should-have-an-extension",
  "Filename": "i-should-have-an-extension",
  "Basename": "i-should-have-an-extension",
  "Extension": "",
  "BOM": false,
  "Shebang": "#!/usr/bin/env sh",
  "Interpreter": "sh",
  "LineEnding": "\n",
  "POSIXy": true
}
```

Perhaps append a `.lisp` extension to such files. Or separate the modulino into clear library vs. command line modules. Or extract the shell interaction into a dedicated script. Or convince the language maintainers to treat shebangs as comments. Write your congressman. However you resolve this, know that the current situation is far outside the norm, and likely to break in a suitably arcane and dramatic fashion. With wyverns and flaming seas and portents of all ill manner.

# LINT

```console
$ make lint
```

# PORT

```console
$ make port
```

# Shell script linters

These bad bois help to shore up ur shell scripts. Though they're designed to work on individual files, so be sure to stank-ify larger projects and pipe the results to `xargs checkbashisms`, yo!

* [checkbashisms](https://sourceforge.net/projects/checkbaskisms/)
* [bashate](https://pypi.python.org/pypi/bashate)
* [shlint](https://rubygems.org/gems/shlint)
* [ShellCheck](https://hackage.haskell.org/package/ShellCheck)
* [editorconfig-cli](https://github.com/amyboyd/editorconfig-cli)
* [shfmt](https://github.com/mvdan/sh)
* [astyle](http://astyle.sourceforge.net)

## Honorable mentions

[ack](https://beyondgrep.com) offers `--shell [-f]` flags that act similarly to `stank`, with the caveat that ack includes nonPOSIX shells like csh, tcsh, and fish in these results; but as of this writing fails to include POSIX shells like ash, dash, posh, pdksh, ksh93, and mksh. ack also depends on Perl, making it more heavyweight for Docker microservices and other constrained platforms.

[linguist](https://github.com/github/linguist), GitHub's extraordinary effort to identify which language each of its millions of repositories are written in. While this stanky Go project does not employ linguist in automated analysis, it's worth mentioning for forensic purposes, if you ever come across a strange, unidentified (or misidentified!) source code file.

[slick](https://github.com/mcandre/slick) offers `sh -n` syntax checking against pure POSIX syntax, whereas actual `sh` on most systems symlinks to bash.
