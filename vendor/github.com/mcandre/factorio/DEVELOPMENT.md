# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.21.5+
* [Node.js](https://nodejs.org/en) 16.14.2+
* [Rust](https://www.rust-lang.org/) 1.68.2+
* [zip](https://linux.die.net/man/1/zip)
* a POSIX compliant [make](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/make.html) implementation (e.g. GNU make, BSD make, etc.)
* Provision additional dev tools with `make`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10 (run `asdf reshim` after provisioning)
* [direnv](https://direnv.net/) 2

# AUDIT

```console
$ mage audit
```

# INSTALL

```console
$ mage install
```

# UNINSTALL

```console
$ mage uninstall
```

# TEST

```console
$ mage [test]
```

# PORT

```console
$ mage port
```

# CLEAN

```console
$ rm -rf bin
```
