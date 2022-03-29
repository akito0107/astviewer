GOOS=js
GOARCH=wasm

build: wasm public/bundle.js
	go build -o bin/astviewer cmd/astviewer/main.go

wasm:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o public/main.wasm ./wasmmain/main.go

public/bundle.js: front/src/index.js front/node_modules
	cd front; NODE_ENV=production yarn webpack

front/node_modules: front/package.json front/yarn.lock
	cd front; yarn install
