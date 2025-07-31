# BUILDTIME REQUIREMENTS

* [Go](https://go.dev/) 1.24.5+
* [POSIX](https://pubs.opengroup.org/onlinepubs/9799919799/) compatible [make](https://en.wikipedia.org/wiki/Make_(software))
* [Rust](https://www.rust-lang.org/) 1.87.0+
* [Snyk](https://snyk.io/)
* Provision additional dev tools with `make [-j 4]`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10 (run `asdf reshim` after provisioning)
* [direnv](https://direnv.net/) 2
* [GNU](https://www.gnu.org/)/[BSD](https://en.wikipedia.org/wiki/Berkeley_Software_Distribution) [make](https://en.wikipedia.org/wiki/Make_(software))
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
