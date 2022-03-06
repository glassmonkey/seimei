.PHONY: build

build:
	go build -o dist/seimei cmd/seimei/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	 golangci-lint run

lint-fix:
	 golangci-lint run --fix