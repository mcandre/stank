# factorio: mo' platforms, mo' problems!

# EXAMPLE

```console
$ cd example

$ factorio

$ tree bin/artifact-dev
bin/artifact-dev
├── darwin
│   ├── amd64
│   │   └── factorio
│   └── arm64
│       └── factorio
...
```

# ABOUT

factorio accelerates Go application development, by automating the process of generating binaries for a multitude of platforms. Your time is valuable. Spend it developing software, not tinkering with toolchains.

# LICENSE

FreeBSD

# DOCUMENTATION

https://godoc.org/github.com/mcandre/factorio

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/factorio/cmd/factorio
```

# RUNTIME REQUIREMENTS

* [Go](https://golang.org/) 1.17+

## Recommended

* [tree](https://linux.die.net/man/1/tree)

# CONTRIBUTING

For more information on developing factorio itself, see [DEVELOPMENT.md](DEVELOPMENT.md).

# CONFIGURATION

The default subdirectory can be customized with a `FACTORIO_BANNER` environment variable, e.g. `FACTORIO_BANNER=hello-0.0.1`. Then artifacts will appear in `bin/hello-0.0.1/`. This is helpful when structuring file paths in prepraration for compressed archives, for example.

factorio primarily assists conventional PC (desktop/laptop/workstation/server) applications. factorio enables the standard platforms expected to work out of the box for `go build`, particularly any pure Go project.

factorio will exclude mobile platforms by default. You can customize the platform blocklist by supplying a Go [Regexp](https://godoc.org/regexp) to a `FACTORIO_PLATFORM_BLOCKLIST` environment variable, e.g. `FACTORIO_PLATFORM_BLOCKLIST=//`.

factorio plugs directly into the standard Go system, compatible with most `go build` flags and environment variables. Any Extra environment variables or flags you pass to `factorio` will propagate to `go build`.

## SEE ALSO

* [cartel](https://github.com/mcandre/cartel) ports native C/C++ applications to Linux variants.
* [gox](https://github.com/mitchellh/gox), an older Go cross-compiler wrapper.
* [tonixxx](https://github.com/mcandre/tonixxx) ports applications of any programming language.
* [xgo](https://github.com/karalabe/xgo) supports Go projects with native cgo dependencies.
