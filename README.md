# stank: analyzers for determining whether files smell like rotten POSIX shell scripts, or faintly rosy like Ruby and Python scripts

# ABOUT

stank is a library and collection of command line utilities for sniffing files to identify shell scripts like bash, sh, zsh, ksh and so on, those funky farmfresh gobs of garbaggio; versus other more palatable files like rb, py, pl. Believe it or not, shell scripts are notoriously difficult to write well, so it behooves a developer to either write shell scripts in safer languages, or else wargame your scripts with an armada of linters. Trouble is, in large projects one can never be too sure which files are honest to dog POSIX compliant shell scripts, and which are pretenders. csh, tcsh, fish, ion, rc, and most other nonderivatives of bash tend to be NOT POSIX compatible. If you're geeky enough to have followed thus far, let's get crackalackin with some fruity examples dammit!

# EXAMPLES

The included `stink` application accepts file and directory names as argument inputs, recursively reporting analytics per-path. Note that some fields are zero valued, not filled in, if one of the fields stands out as obviously a POSIX shell script.

```console
$ cat examples/hello
#!/bin/sh
echo "Hello"

$ stink examples/hello
{"Path":"examples/hello","Filename":"hello","Basename":"hello","Extension":"","BOM":false,"Shebang":"#!/bin/sh","Interpreter":"/bin/sh","LineEnding":"\n","POSIXy":true}

$ stink -pp examples/hello
{
  "Path": "examples/hello",
  "Filename": "hello",
  "Basename": "hello",
  "Extension": "",
  "BOM": false,
  "Shebang": "#!/bin/sh",
  "Interpreter": "/bin/sh",
  "LineEnding": "\n",
  "POSIXy": true
}

$ cat examples/hello.py
#!/usr/bin/env python
print "Hello"

$ stink -pp examples/hello.py
{
  "Path": "examples/hello.py",
  "Filename": "hello.py",
  "Basename": "hello.py",
  "Extension": ".py",
  "BOM": false,
  "Shebang": "",
  "Interpreter": "",
  "LineEnding": "",
  "POSIXy": false
}

$ cat examples/wednesday
#!/bin/bash
echo -e "Óðinn!\n(I am Odin.)"

$ stink -pp examples/wednesday
{
  "Path": "examples/wednesday",
  "Filename": "wednesday",
  "Basename": "wednesday",
  "Extension": "",
  "BOM": false,
  "Shebang": "#!/bin/bash",
  "Interpreter": "/bin/bash",
  "LineEnding": "\n",
  "POSIXy": true
}

$ cat examples/lo
#!/bin/csh
echo "Lo"

$ stink -pp examples/lo
{
  "Path": "examples/lo",
  "Filename": "lo",
  "Basename": "lo",
  "Extension": "",
  "BOM": false,
  "Shebang": "#!/bin/csh",
  "Interpreter": "/bin/csh",
  "LineEnding": "\n",
  "POSIXy": false
}

$ stink -help
  -help
        Show usage information
  -pp
        Prettyprint smell records
  -version
        Show version information
```

And so on. Basically, `#!.*sh.*` or `\..*sh` tends to yield "POSIXy" true, whereas csh, tcsh, fish, python, ruby, lua, node.js crap, and animated GIFs tend to yeidl "POSIXy" false.

A second included application, `stank`, accepts file and directory names, recursively selecting just the POSIX shell script paths and printing these. nonPOSIX paths are omitted.

```console
$ ls examples | head
CHANGELOG
README
badhello.lua
badhello.py
blank.bash
globs.bash
goodbye
goodbye.sh
greetings.bash
hello

$ stank examples | head
examples/.profile
examples/.zshrc
examples/blank.bash
examples/globs.bash
examples/goodbye.sh
examples/greetings.bash
examples/hello
examples/hello.sh
examples/hooks/post-update
examples/hooks/post-update.sample
```

The included `examples/` directory demonstrates many edge cases, such as empty scripts, shebang-less scripts, extensioned and extensionless scripts, and various Hello World applications in across many programming languages. Some files, such as `examples/goodbye` may contain 100% valid POSIX shell script content, but fail to self-identify with either shebangs or relevant file extensions. In a large project, such files may be mistakenly treated as whoknowswhat format, or simply plain text. Perhaps statistical methods could help identify POSIX grammars, but even an empty file is technically POSIX, which is unhelpful from a reliable classification standpoint. In any case, `examples/` hopefully covers the more common edge cases.

# DOWNLOADS

https://github.com/mcandre/stank/releases

# DOCUMENTATION

http://godoc.org/github.com/mcandre/stank

# REQUIREMENTS

The `stink` and `stank` applications have no special runtime requirements.

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
* [flcl](https://github.com/mcandre/flcl)
* [editorconfig-cli](https://github.com/amyboyd/editorconfig-cli)

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

Honorable mention for [linguist](https://github.com/github/linguist), GitHub's extraordinary effort to identify which language each of its millions of repositories are written in. While this stanky Go project does not employ linguist in automated analysis, it's worth mentioning for forensic purposes, if you ever come across a strange, unidentified (or misidentified!) source code file.
