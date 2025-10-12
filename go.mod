module github.com/mcandre/stank

go 1.24.6

tool (
	github.com/alexkohler/nakedret/v2/cmd/nakedret
	github.com/kisielk/errcheck
	github.com/magefile/mage
	github.com/mcandre/factorio/cmd/factorio
	honnef.co/go/tools/cmd/staticcheck
)

require (
	github.com/magefile/mage v1.15.0
	github.com/mcandre/mage-extras v0.0.26
	mvdan.cc/sh/v3 v3.12.0
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/alexkohler/nakedret/v2 v2.0.6 // indirect
	github.com/kisielk/errcheck v1.9.0 // indirect
	github.com/mcandre/factorio v0.0.13 // indirect
	golang.org/x/exp/typeparams v0.0.0-20250408133849-7e4ce0ab07d0 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/tools v0.32.0 // indirect
	honnef.co/go/tools v0.6.1 // indirect
)
