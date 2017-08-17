VERSION=0.0.4

.PHONY: port clean clean-ports

all: port

govet:
	find . -path "*/vendor*" -prune -o -name "*.go" -type f -exec go tool vet -shadow {} \;

golint:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec golint {} \;

gofmt:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;

goimports:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec goimports -w {} \;

errcheck:
	errcheck -blank

opennota-check:
	aligncheck
	structcheck
	varcheck

megacheck:
	megacheck

editorconfig:
	flcl . | xargs -n 100 editorconfig-cli check

lint: govet golint gofmt goimports errcheck opennota-check megacheck editorconfig

port: archive-ports

archive-ports: bin
	zipc -chdir bin "stank-$(VERSION).zip" "stank-$(VERSION)"

bin:
	gox -output="bin/stank-$(VERSION)/{{.OS}}/{{.Arch}}/{{.Dir}}" ./cmd...

clean: clean-ports

clean-ports:
	rm -rf bin
