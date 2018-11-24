// +build js,wasm

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"syscall/js"
)

func main() {
	gosrc := js.Global().Get("document").Call("getElementById", "gosrc")
	src := gosrc.Get("value").String()

	f, err := parser.ParseFile(token.NewFileSet(), "", src, parser.AllErrors)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	goout := js.Global().Get("document").Call("getElementById", "goout")
	var buf bytes.Buffer
	var depth int
	ast.Inspect(f, func(n ast.Node) bool {
		if n != nil {
			depth--
		} else {
			depth++
			return true
		}
		if _, err := fmt.Fprintf(&buf, "%*s%#v\n", depth*2, "", n); err != nil {
			log.Printf("%+v", err)
			return false
		}
		return true
	})
	goout.Set("innerHTML", buf.String())
}
