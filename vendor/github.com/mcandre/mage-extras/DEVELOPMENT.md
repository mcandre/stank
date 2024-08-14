# BUILDTIME REQUIREMENTS

* [Go](https://go.dev/) 1.23.0+
* POSIX compatible [make](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/make.html)
* [Node.js](https://nodejs.org/en) 20.10.0+
* [Rust](https://www.rust-lang.org/) 1.75.0+
* Provision additional dev tools with `make`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10 (run `asdf reshim` after provisioning)
* [direnv](https://direnv.net/) 2
* macOS [open](https://ss64.com/mac/open.html) or equivalent alias

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
$ open cover.html
```

# LINT

```console
$ mage lint
```

# CLEAN

```console
$ mage clean
```
