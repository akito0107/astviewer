import * as monaco from "monaco-editor";
import * as $ from "jquery";

const gocode = `package main

import "fmt"

func main() {
	fmt.Println("hello, 世界")
}
`;

if (!WebAssembly.instantiateStreaming) {
  // polyfill
  WebAssembly.instantiateStreaming = async (resp, importObject) => {
    const source = await (await resp).arrayBuffer();
    return await WebAssembly.instantiate(source, importObject);
  };
}

const editor = monaco.editor.create(document.getElementById("monaco-editor"), {
  value: gocode,
  language: "go",
  lineNumbers: "on",
  theme: "vs-dark",
});
const input = document.getElementById("gosrc");

const model = editor.getModel();
let rows = document.getElementsByClassName("astline");

const go = new Go();
let mod, inst;

model.onDidChangeContent(() => {
  WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
    .then((result) => {
      mod = result.module;
      inst = result.instance;
      input.value = model.getValue();
      go.run(inst);
      rows = document.getElementsByClassName("astline");
    })
    .catch((err) => {
      console.error(err);
    });
});

editor.onDidChangeCursorSelection((event) => {
  const start = event.selection.startLineNumber;
  const end = event.selection.endLineNumber;
  $(".astline").removeClass("selected");

  for (let i = start; i <= end; i++) {
    $(`.lineno${i}`).addClass("selected");
  }
});
