.PHONY: build

build:
	go build -o dist/seimei cmd/seimei/main.go

.PHONY: test
test:
	go test -v ./...