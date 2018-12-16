import * as monaco from "monaco-editor";
import * as $ from "jquery";

const gocode = `package main

import "fmt"

func main() {
	fmt.Println("hello, 世界")
}
`;

const editor = monaco.editor.create(document.getElementById("monaco-editor"), {
  value: gocode,
  language: "go",
  lineNumbers: "on",
  theme: "vs-dark"
});
const input = document.getElementById("gosrc");

const go = new Go();

const model = editor.getModel();
let rows = document.getElementsByClassName("astline");

model.onDidChangeContent(() => {
  input.value = model.getValue();
  WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
    result => {
      go.run(result.instance);
      rows = document.getElementsByClassName("astline");
    }
  );
});

editor.onDidChangeCursorSelection(event => {
  const start = event.selection.startLineNumber;
  const end = event.selection.endLineNumber;
  $(".astline").removeClass("selected");

  for (let i = start; i <= end; i++) {
    $(`.lineno${i}`).addClass("selected");
  }
});
