# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.17+
* [zip](https://linux.die.net/man/1/zip)

# INSTALL

```console
$ go install ./...
```

# UNINSTALL

```console
$ rm "$GOPATH/src/factorio"
```

# TEST

```console
$ factorio
```

# PORT

```console
$ FACTORIO_BANNER=factorio-0.0.1 factorio

$ cd bin

$ zip -r factorio-0.0.1.zip factorio-0.0.1
```
