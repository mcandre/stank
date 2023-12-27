# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.21.5+
* [Node.js](https://nodejs.org/en) 16.14.2+
* [Rust](https://www.rust-lang.org/) 1.68.2+
* a POSIX compliant [make](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/make.html) implementation (e.g. GNU make, BSD make, etc.)
* Provision additional dev tools with `make`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10 (run `asdf reshim` after provisioning)
* [direnv](https://direnv.net/) 2

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
