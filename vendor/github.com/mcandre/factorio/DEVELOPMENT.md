# BUILDTIME REQUIREMENTS

* [Go](https://go.dev/) 1.24.0+
* [GNU](https://www.gnu.org/software/make/) / [BSD](https://man.freebsd.org/cgi/man.cgi?make(1)) make
* [Rust](https://www.rust-lang.org/) 1.75.0+
* [Snyk](https://snyk.io/)
* POSIX compatible [tar](https://pubs.opengroup.org/onlinepubs/7908799/xcu/tar.html)
* Provision additional dev tools with `make [-j 4]`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10 (run `asdf reshim` after provisioning)
* [direnv](https://direnv.net/) 2
* a UNIX environment, such as macOS, Linux, BSD, [WSL](https://learn.microsoft.com/en-us/windows/wsl/), etc.

Non-UNIX environments may produce subtle adverse effects when linting or generating application ports.

## Windows

Apply a user environment variable `GODEBUG=modcacheunzipinplace=1` per [access denied resolution](https://github.com/golang/go/wiki/Modules/e93463d3e853031af84204dc5d3e2a9a710a7607#go-115), for native Windows development environments (Command Prompt / PowerShell, not WLS, not Cygwin, not MSYS2, not MinGW, not msysGit, not Git Bash, not etc).

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

# LINT

```console
$ mage lint
```

# TEST

```console
$ mage test
```

# PORT

```console
$ mage port
```
