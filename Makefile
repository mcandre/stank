VERSION=0.0.1

.PHONY: port clean clean-ports

all: port

govet:
	find . -path "*/vendor*" -prune -o -name "*.go" -type f -exec go tool vet -shadow {} \;

golint:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec golint {} \;

gofmt:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;

goimport:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec goimports -w {} \;

errcheck:
	errcheck

# stank would be a fantastic candidate for replacing the find exclusions in bashate, shlint, checkbashisms, and shellcheck. However, this would result in the development setup for stank depending on itself, which is no bueno. Other projects are encouraged to run stank | xargs bashate, etc. but the stank project itself will continue to leverage more traditional tools for build system portability.

editorconfig:
	flcl . | xargs -n 100 editorconfig-cli check

lint: govet golint gofmt goimport errcheck editorconfig
