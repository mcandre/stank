# BUILDTIME REQUIREMENTS

* [Go](https://go.dev/)
* POSIX compliant [make](https://pubs.opengroup.org/onlinepubs/9799919799/utilities/make.html)
* Provision additional dev tools with `make`

## Recommended

* a UNIX-like environment (e.g. [WSL](https://learn.microsoft.com/en-us/windows/wsl/))
* [ASDF](https://asdf-vm.com/) 0.18 (run `asdf reshim` after provisioning)
* [direnv](https://direnv.net/) 2
* macOS [open](https://ss64.com/mac/open.html) or equivalent alias

## Windows

Apply a user environment variable `GODEBUG=modcacheunzipinplace=1` per [access denied resolution](https://github.com/golang/go/wiki/Modules/e93463d3e853031af84204dc5d3e2a9a710a7607#go-115), for native Windows development environments (Command Prompt / PowerShell, not WLS, not Cygwin, not MSYS2, not MinGW, not msysGit, not Git Bash, not etc).

# AUDIT

```console
$ mage audit
```

# TEST

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
