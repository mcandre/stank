# stank: analyzers for determining whether files smell like rotten POSIX shell scripts, or faintly rosy like Ruby and Python scripts

# ABOUT

stank is a library and collection of command line utilities for sniffing files to identify shell scripts like bash, sh, zsh, ksh and so on, those funky farmfresh gobs of garbaggio; versus other more palatable files like rb, py, pl. Believe it or not, shell scripts are notoriously difficult to write well, so it behooves a developer to either write shell scripts in safer languages, or else wargame your scripts with an armada of linters. Trouble is, in large projects one can never be too sure which files are honest to dog POSIX compliant shell scripts, and which are pretenders. csh, tcsh, fish, ion, rc, and most other nonderivatives of bash tend to be NOT POSIX compatible. If you're geeky enough to have followed thus far, let's get crackalackin with some fruity examples dammit!

# EXAMPLES

The stank system includes the stank Go library as well as three command line utilities for convenience. `rosy` recursively searches directory trees for POSIX shell scripts, recommending that they be rewritten in safer general purpose languages like Ruby, Python, Node.js, etc.

```console
$ rosy examples/hooks
Rewrite POSIX script in Ruby or other safer general purpose scripting language: examples/hooks/post-update
Rewrite POSIX script in Ruby or other safer general purpose scripting language: examples/hooks/post-update.sample
Rewrite POSIX script in Ruby or other safer general purpose scripting language: examples/hooks/pre-applypatch
Rewrite POSIX script in Ruby or other safer general purpose scripting language: examples/hooks/pre-applypatch.sample
Rewrite POSIX script in Ruby or other safer general purpose scripting language: examples/hooks/pre-commit
Rewrite POSIX script in Ruby or other safer general purpose scripting language: examples/hooks/pre-commit.sample
Rewrite POSIX script in Ruby or other safer general purpose scripting language: examples/hooks/pre-push
Rewrite POSIX script in Ruby or other safer general purpose scripting language: examples/hooks/pre-push.sample
Rewrite POSIX script in Ruby or other safer general purpose scripting language: examples/hooks/pre-rebase
Rewrite POSIX script in Ruby or other safer general purpose scripting language: examples/hooks/pre-rebase.sample
Rewrite POSIX script in Ruby or other safer general purpose scripting language: examples/hooks/update
Rewrite POSIX script in Ruby or other safer general purpose scripting language: examples/hooks/update.sample
$ echo "$?"
1

$ rosy examples/just-python
$ echo "$?"
0

$ rosy -help
  -help
        Show usage information
  -version
        Show version information
```

The `stank` application prints paths to POSIX shell scripts, designed for use in combination with `xargs` to help per-file shell static analysis applications lint large projects.

```console
$ stank examples
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

$ stank examples/hooks | xargs checkbashisms
error: examples/hooks/pre-rebase: Unterminated quoted string found, EOF reached. Wanted: <'>, opened in line 133
error: examples/hooks/pre-rebase.sample: Unterminated quoted string found, EOF reached. Wanted: <'>, opened in line 133

$ stank -help
  -help
        Show usage information
  -version
        Show version information
```

Finally, `stink` prints a record of each file's POSIXyness, including any interesting fields it identified along the way. Note that some fields may be zero valued if the stench of POSIX or rosy waft of nonPOSIX is overwhelming, short-circuiting analysis. This short-circuiting feature dramatically speeds up how `stank` and `rosy` search large projects.

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

$ stink -help
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
$ git clone git@github.com:mcandre/stank.git $GOPATH/src/github.com/stank
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
