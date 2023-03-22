# BUILDTIME REQUIREMENTS

* a POSIX compatible shell (e.g., `bash`, `ksh`, `sh`, `zsh`)
* [Go](https://golang.org/) 1.20.2+ with `sh acquire`
* [Node.js](https://nodejs.org/en) 16.14.2+ with `npm install -g snyk@1.996.0`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10

# AUDIT

```console
$ mage audit
```

# UNIT TEST

```console
$ mage test
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

# CLEAN

```console
$ mage clean
```
