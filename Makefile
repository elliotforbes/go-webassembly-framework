GOCMD=go

build:
	GOOS=js GOARCH=wasm go build -o lib.wasm cmd/oak/main.go
build-cli:
	GOOS=darwin GOARCH=amd64 go build -o oak cmd/oak-cli/main.go
	PATH=bin:${PATH}
run:
	$(GOCMD) build 
