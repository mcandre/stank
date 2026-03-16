# mx: mage extras

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/mcandre/mx) [![Test](https://github.com/mcandre/mx/actions/workflows/test.yml/badge.svg)](https://github.com/mcandre/mx/actions/workflows/test.yml) [![license](https://img.shields.io/badge/license-BSD-0)](LICENSE.md)

```txt
  _ __ ___ __  __
 | '_ ` _ \\ \/ /
 | | | | | |>  <
 |_| |_| |_/_/\_\
```

# SUMMARY

mx streamlines common Go development tasks for projects built with [mage](https://magefile.org/).

# ABOUT

[API Docs](https://pkg.go.dev/github.com/mcandre/mx)

mx provides utility functions for common Go development operations.

Examples:

* `GoEnv` - Query the `go env` toolchain configuration subsystem
* `Install` - Compile and install Go executables
* Lint Go projects recursively:
  * `GoImports`
  * `GoVet`
  * `GoVetShadow`
  * `Nakedret`
* `UnitTest` - trigger unit test suite

# SYSTEM REQUIREMENTS

* [Go](https://go.dev/)
* [Mage](https://magefile.org/) 1.16.1+

For details on developing mx, see our [development guide](DEVELOPMENT.md).

🐱
