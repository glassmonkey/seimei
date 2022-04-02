.PHONY: build

Version=dev
Revision=$(shell git rev-parse --short HEAD)

build:
	go build -o dist/seimei -ldflags "-X main.Version=$(Version) -X main.Revision=$(Revision)"  cmd/seimei/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test -cover -v ./... -coverprofile=dist/cover.out
	go tool cover -html=dist/cover.out -o dist/cover.html
	open dist/cover.html


.PHONY: lint
lint:
	 golangci-lint run

lint-fix:
	 golangci-lint run --fix
