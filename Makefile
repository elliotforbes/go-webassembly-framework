GOCMD=go

build:
	go get ./...
	GOOS=js GOARCH=wasm go build -o lib.wasm cmd/oak/main.go
build-cli:
	go get ./...
	GOOS=darwin GOARCH=amd64 go build -o oak cmd/oak-cli/main.go
	PATH=bin:${PATH}
run:
	$(GOCMD) build 
