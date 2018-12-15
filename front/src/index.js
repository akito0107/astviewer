import * as monaco from "monaco-editor";

const gocode = `package main

import "fmt"

func main() {
	fmt.Println("hello, 世界")
}
`

monaco.editor.create(document.getElementById("container"), {
    value: gocode,
    language: "go",
    lineNumbers: "on",
    theme: "vs-dark",
});
