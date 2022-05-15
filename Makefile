

Version=$(shell git describe --tags --abbrev=0)
Revision=$(shell git rev-parse --short HEAD)

.PHONY: build
build:
	go build -o dist/seimei -ldflags "-X main.Version=$(Version) -X main.Revision=$(Revision)"  cmd/seimei/main.go

.PHONY: release
release:
	goreleaser --snapshot --skip-publish --rm-dist

.PHONY: test
test:
	go test -v $(go list ./... | grep -v /benchmark/)

.PHONY: test-coverage
test-coverage:
	go test -cover -v ./... -coverprofile=dist/cover.out
	go tool cover -html=dist/cover.out -o dist/cover.html
	open dist/cover.html


.PHONY: lint
lint:
	 golangci-lint run

.PHONY: lint-fix
lint-fix:
	 golangci-lint run --fix

 .PHONY: version
version:
	@echo "$(Version)-$(Revision)" > version.txt
