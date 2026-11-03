module github.com/mcandre/stank

go 1.27.3

require (
	github.com/magefile/mage v1.17.2
	github.com/mcandre/mx v0.0.47
	mvdan.cc/sh/v3 v3.14.1
)

tool (
	github.com/alexkohler/nakedret/v2/cmd/nakedret
	github.com/kisielk/errcheck
	github.com/magefile/mage
	honnef.co/go/tools/cmd/staticcheck
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/alexkohler/nakedret/v2 v2.0.6 // indirect
	github.com/kisielk/errcheck v1.30.0 // indirect
	golang.org/x/exp/typeparams v0.0.0-20250408133849-7e4ce0ab07d0 // indirect
	golang.org/x/mod v0.41.0 // indirect
	golang.org/x/sync v0.23.0 // indirect
	golang.org/x/tools v0.51.0 // indirect
	honnef.co/go/tools v0.8.0-rc.1 // indirect
)
