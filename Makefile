all: lint test build start

start:
	./bin/InfobaeAPI

dev: path
	gowatch

fmt:
	@if [ -n "$$(go fmt ./src)" ]; then \
        echo "Code is not properly formatted"; \
        exit 1; \
    fi

lint:
	golangci-lint run

clear:
	rm -rf ./bin

## https://github.com/golang-standards/project-layout/issues/113#issuecomment-1336514449
build: clear fmt
	GOARCH=amd64 go build -o ./bin/InfobaeAPI ./src/main.go

build-arm: clear fmt
	GOARCH=arm64 go build -o ./bin/InfobaeAPI ./src/main.go

test:
	go test ./src/...

path:
	export PATH=$PATH:$HOME/go/bin