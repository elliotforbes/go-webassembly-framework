build:
	GO111MODULE=on
	GOOS=js GOARCH=wasm go get
	GOOS=js GOARCH=wasm go build -o lib.wasm main.go
build-cli:
	GO111MODULE=on
	GOOS=darwin GOARCH=amd64 go build -o oak cmd/oak-cli/main.go
run:
	$(GOCMD) build 
