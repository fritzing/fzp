# github.com/fritzing/fzp/Makefile

all: test build

test:
	@go test -v -cover ./...

build:
	@cd bin/fzp && go build && ./fzp --help
