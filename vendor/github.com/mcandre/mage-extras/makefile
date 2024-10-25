.POSIX:
.SILENT:
.PHONY: \
	all \
	go \
	rust

all: go rust

go:
	go install github.com/alexkohler/nakedret/v2/cmd/nakedret@v2.0.3
	go install github.com/kisielk/errcheck@v1.7.0
	go install github.com/magefile/mage@v1.14.0
	go install github.com/mgechev/revive@v1.4.0
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install honnef.co/go/tools/cmd/staticcheck@2024.1
	go mod tidy

rust:
	cargo install --force unmake@0.0.17
