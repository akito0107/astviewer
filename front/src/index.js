import * as monaco from "monaco-editor";

const gocode = `package main

import "fmt"

func main() {
	fmt.Println("hello, 世界")
}
`

console.log('hoge')

monaco.editor.create(document.getElementById("gosrc"), {
    value: gocode,
    language: "go",
    lineNumbers: "on",
    theme: "vs-dark",
});

const go = new Go();
const gosrc = document.getElementById("gosrc")
gosrc.addEventListener('change', () => {
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
        PR.prettyPrint();
    });
});
