# github.com/fritzing/fzp/Makefile

all: test build
	./bin/fzp/fzp --help

test:
	@go test -v -cover ./...

build:
	@cd bin/fzp && go build

clean:
	@go clean ./...
