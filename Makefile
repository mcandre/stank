VERSION=0.0.1

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
	errcheck

# stank would be a fantastic candidate for replacing the find exclusions in bashate, shlint, checkbashisms, and shellcheck. However, this would result in the development setup for stank depending on itself, which is no bueno. Other projects are encouraged to run stank | xargs bashate, etc. but the stank project itself will continue to leverage more traditional tools for build system portability.

editorconfig:
	flcl . | xargs -n 100 editorconfig-cli check

lint: govet golint gofmt goimports errcheck editorconfig

port: archive-ports

archive-ports: bin
	zipc -chdir bin "stank-$(VERSION).zip" "stank-$(VERSION)"

bin:
	gox -output="bin/stank-$(VERSION)/{{.OS}}/{{.Arch}}/{{.Dir}}" ./cmd...

clean: clean-ports

clean-ports:
	rm -rf bin
