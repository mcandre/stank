VERSION=0.0.12

.PHONY: port clean clean-ports

all: install

install-stink:
	sh -c 'cd cmd/stink && go install'

install-stank:
	sh -c 'cd cmd/stank && go install'

install-rosy:
	sh -c 'cd cmd/rosy && go install'

install-funk:
	sh -c 'cd cmd/funk && go install'

install: install-stink install-stank install-rosy install-funk

uninstall:
	-rm "$$GOPATH/bin/funk"
	-rm "$$GOPATH/bin/rosy"
	-rm "$$GOPATH/bin/stank"
	-rm "$$GOPATH/bin/stink"

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

nakedret:
	nakedret -l 0 ./...

opennota-check:
	aligncheck
	structcheck
	varcheck

megacheck:
	megacheck

editorconfig:
	flcl . | xargs -n 100 editorconfig-cli check

lint: govet golint gofmt goimports errcheck nakedret opennota-check megacheck editorconfig

port: archive-ports

archive-ports: bin
	zipc -chdir bin "stank-$(VERSION).zip" "stank-$(VERSION)"

bin:
	gox -output="bin/stank-$(VERSION)/{{.OS}}/{{.Arch}}/{{.Dir}}" ./cmd/...

clean: clean-ports

clean-ports:
	rm -rf bin
