.POSIX:
.SILENT:
.PHONY: all

all:
	cargo install --force \
		chandler@0.0.9 \
		tuggy@0.0.29
	go install golang.org/x/tools/cmd/deadcode@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install tool
	go mod tidy
