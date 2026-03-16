# DEVELOPMENT GUIDE

stank follows standard, cargo based operations for compiling and unit testing Go code.

For advanced operations, such as linting, we further supplement with some software industry tools.

# BUILDTIME REQUIREMENTS

* [Go](https://go.dev/)
* POSIX compliant [make](https://pubs.opengroup.org/onlinepubs/9799919799/utilities/make.html)
* Provision additional dev tools with `make`

## Recommended

* a UNIX-like environment (e.g. [WSL](https://learn.microsoft.com/en-us/windows/wsl/))
* [ASDF](https://asdf-vm.com/) 0.18 (run `asdf reshim` after provisioning)

# AUDIT

```sh
mage audit
```

# INSTALL

```sh
mage install
```

# UNINSTALL

```sh
mage uninstall
```

# TEST

```sh
mage test
```

# LINT

```sh
mage lint
```

# CLEAN

```sh
mage clean
```
