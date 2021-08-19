# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.17+
* a POSIX compatible shell (e.g., `bash`, `ksh`, `sh`, `zsh`)
* Go development tools (`sh acquire`)

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
