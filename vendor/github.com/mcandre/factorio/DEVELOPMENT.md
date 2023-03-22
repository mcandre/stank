# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.20.2+ with `go install github.com/mcandre/accio/cmd/accio@v0.0.3` and `accio -install`
* [Node.js](https://nodejs.org/en) 16.14.2+ with `npm install -g snyk@1.996.0`
* [zip](https://linux.die.net/man/1/zip)

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10
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
