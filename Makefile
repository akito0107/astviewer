GOOS=js
GOARCH=wasm

build: tools statik/statik.go
	go build -o bin/astviewer cmd/astviewer/main.go

statik/statik.go: wasm public/bundle.js
	tools/statik -src=./public

wasm:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o public/main.wasm ./wasmmain/main.go

tools: tools/goexec tools/statik

tools/goexec: vendor/github.com/shurcooL/goexec
	go build -o tools/goexec vendor/github.com/shurcooL/goexec/main.go

tools/statik: vendor/github.com/rakyll/statik
	go build -o tools/statik vendor/github.com/rakyll/statik/statik.go

vendor/github.com/shurcooL/goexec: vendor

vendor/github.com/rakyll/statik: vendor

vendor: Gopkg.toml Gopkg.lock
	dep ensure

public/bundle.js: front/src/index.js front/node_modules
	cd front; NODE_ENV=production yarn webpack

front/node_modules: front/package.json front/yarn.lock
	cd front; yarn install
