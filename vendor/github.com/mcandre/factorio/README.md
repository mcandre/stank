# factorio: Go cross-compiler

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

factorio is fast. It has no intrinsic dependency on any containers or virtual machines. Factorio plugs directly into the standard `go` command line system.

# LICENSE

BSD-2-Clause

# DOCUMENTATION

https://pkg.go.dev/github.com/mcandre/factorio

# DOWNLOAD

https://github.com/mcandre/factorio/releases

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/factorio/cmd/factorio@latest
```

# RUNTIME REQUIREMENTS

* [Go](https://go.dev/) 1.24.4+

## Recommended

* [GNU](https://www.gnu.org/)/[BSD](https://en.wikipedia.org/wiki/Berkeley_Software_Distribution) [tar](https://en.wikipedia.org/wiki/Tar_(computing))
* [tree](https://en.wikipedia.org/wiki/Tree_(command))
* a [UNIX](https://en.wikipedia.org/wiki/Unix)-like environment

tar is a portable archiver suitable for creating `*.tgz` tarball archives. Users can then download the tarball and extract the executable relevant to their platform. Tarballs are especially well suited for use in Docker containers, as the tar command is more likely to be installed than unzip.

Note that non-UNIX file systems may not preserve crucial chmod ACL bits during port generation. This can corrupt downstream artifacts, such as compressed archives and installation procedures.

# CONTRIBUTING

For more information on developing factorio itself, see [DEVELOPMENT.md](DEVELOPMENT.md).

# CONFIGURATION

The default subdirectory can be customized with a `FACTORIO_BANNER` environment variable, e.g. `FACTORIO_BANNER=hello-0.0.1`. Then artifacts will appear in `bin/hello-0.0.1/`. This is helpful when structuring file paths in prepraration for compressed archives, for example.

factorio primarily assists conventional PC (desktop/laptop/workstation/server) applications. factorio enables the standard platforms expected to work out of the box for `go build`, particularly any pure Go project.

factorio will exclude mobile platforms by default. You can customize the platform blocklist by supplying a Go [Regexp](https://godoc.org/regexp) to a `FACTORIO_PLATFORM_BLOCKLIST` environment variable, e.g. `FACTORIO_PLATFORM_BLOCKLIST=//`.

factorio is essentially compatible with `go build` flags and environment variables. Any Extra environment variables or flags you pass to `factorio` will propagate to `go build`.

## SEE ALSO

* [crit](https://github.com/mcandre/crit) generates Rust ports
* [LLVM](https://llvm.org/) bitcode offers an abstract assembler format for C/C++ code.
* [tug](https://github.com/mcandre/tug) automates multi-platform Docker image builds.
* [WASM](https://webassembly.org/) provides a portable interface for C/C++ code.
* [xgo](https://github.com/techknowlogick/xgo) supports Go projects with native cgo dependencies.
