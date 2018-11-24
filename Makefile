GOOS=js
GOARCH=wasm

build: tools
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o public/main.wasm

tools: tools/goexec

tools/goexec: vendor/github.com/shurcooL/goexec
	go build -o tools/goexec vendor/github.com/shurcooL/goexec/main.go

vendor/github.com/shurcooL/goexec: vendor

vendor: Gopkg.toml
	dep ensure

