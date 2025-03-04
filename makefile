.POSIX:
.SILENT:
.PHONY: \
	all \
	go \
	rust

all: go rust

go:
	go install golang.org/x/tools/cmd/deadcode@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go mod tidy
	go install tool

rust:
	cargo install --force unmake@0.0.17
