test:
	GOOS=js GOARCH=wasm go test main.go
build-cli:
	GOOS=darwin GOARCH=amd64 go build -o oak cmd/oak-cli/main.go

