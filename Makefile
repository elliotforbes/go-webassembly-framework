GOCMD=go

build:
	GO111MODULE=on
	go get ./...
	GOOS=js GOARCH=wasm go build -o lib.wasm cmd/oak/main.go
build-cli:
	GO111MODULE=on
	go get ./...
	GOOS=darwin GOARCH=amd64 go build -o oak cmd/oak-cli/main.go
	PATH=bin:${PATH}
run:
	$(GOCMD) build 
