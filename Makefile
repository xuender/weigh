PACKAGE = github.com/xuender/weigh
default: lint test

tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/cespare/reflex@latest
	go install github.com/xuender/go-cli@latest

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...

lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...

test:
	go test -race -v ./... -gcflags=all=-l -cover

watch-test:
	reflex -t 50ms -s -- sh -c 'gotest -v ./...'

clean:
	rm -rf dist

build:
	go build -o dist/weigh -tags="sonic avx" cmd/weigh/main.go

dev:
	go-cli watch go run cmd/weigh/main.go

wire:
	wire gen ${PACKAGE}/app

proto:
	protoc --go_out=. pb/*.proto
