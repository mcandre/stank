# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.17+
* a POSIX compatible shell (e.g., `bash`, `ksh`, `sh`, `zsh`)
* Go development tools (`sh acquire`)

## Recommended

* [snyk](https://www.npmjs.com/package/snyk) 1.893.0 (`npm install -g snyk@1.893.0`)

# SECURITY AUDIT

```console
$ snyk test
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
