# mage-extras: some predefined tasks for common mage workflows

# EXAMPLE

```console
$ mage noVendor
/Users/andrew/go/src/github.com/mcandre/mage-extras/test.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/mageextras_test.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/errcheck.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/gofmt.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/sources.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/gox.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/nakedret.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/vet.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/goimports.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/mageextras.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/packages.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/pathseparator.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/binaries_test.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/sources_test.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/binaries.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/golint.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/install.go
```

# ABOUT

mage-extras defines some reusable task predicates for common workflows, in a platform-agnostic way:

* checking that Go source code actually compiles
* running unit tests
* generating code coverage reports
* linting with assorted Go linting tools
* formatting Go code
* installing and uninstall Go applications
* collecting Go source file paths
* obtaining the GOPATH/bin directory
* referencing all local Go packages
* referencing all local Go commands
* cross-compiling applications with factorio, gox, goxcart, and xgo
* archiving artifacts
* manipulating the path separator as a string

Mage is highly agnostic about workflows. mage-extras is a little more opinionated, introducing some useful conventions on top, such as reliably obtaining a list of non-vendored Go files paths, while allowing developers to customize builds to suit their project needs.

# DOCUMENTATION

https://godoc.org/github.com/mcandre/mage-extras

# RUNTIME REQUIREMENTS

* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)

## Recommended

* [karp](https://github.com/mcandre/karp) (e.g. `go get https://github.com/mcandre/karp/...`) for conveniently browsing coverage reports.

# CONTRIBUTING

For more details on developing mage-extras itself, see [DEVELOPMENT.md](DEVELOPMENT.md).
