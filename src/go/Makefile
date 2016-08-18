# github.com/fritzing/fzp/src/go/Makefile

all: test lint

test:
	@go test -v -cover ./...

cover:
	go test -coverprofile cover.out
	go tool cover -html=cover.out -o cover.html

lint:
	@golint

clean:
	rm -f cover.out
	rm -f cover.html
