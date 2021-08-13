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
* cross-compiling applications with gox or goxcart
* archiving artifacts
* manipulating the path separator as a string

Mage is highly agnostic about workflows. mage-extras is a little more opinionated, introducing some useful conventions on top, such as reliably obtaining a list of non-vendored Go files paths, while allowing developers to customize builds to suit their project needs.

# DOCUMENTATION

https://godoc.org/github.com/mcandre/mage-extras

# RUNTIME REQUIREMENTS

* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)

## Recommended

* [karp](https://github.com/mcandre/karp) (e.g. `go get https://github.com/mcandre/karp/...`) for conveniently browsing coverage reports.

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.11+
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [shadow](golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow) (e.g. `go get -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow`)

# INSTALL

```console
$ go get github.com/mcandre/mage-extras
```

# UNIT TEST

```console
$ go test
```

# COVERAGE

```console
$ mage coverageHTML
$ karp cover.html
```

# LINT

```console
$ mage lint
```

# CLEAN ALL ARTIFACTS

```console
$ mage clean; mage -clean
```
