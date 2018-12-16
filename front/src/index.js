import * as monaco from "monaco-editor";

const gocode = `package main

import "fmt"

func main() {
	fmt.Println("hello, 世界")
}
`

console.log('hoge')

const editor = monaco.editor.create(document.getElementById("monaco-editor"), {
    value: gocode,
    language: "go",
    lineNumbers: "on",
    theme: "vs-dark",
});
const input = document.getElementById("gosrc");

const go = new Go();

const model = editor.getModel();

model.onDidChangeContent((event) => {
    input.value = model.getValue();
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
    });
});
