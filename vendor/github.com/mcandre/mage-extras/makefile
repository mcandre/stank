.POSIX:
.SILENT:
.PHONY: \
	all \
	go \
	rust

all: go rust

go:
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install tool
	go mod tidy

rust:
	cargo install --force unmake@0.0.17
