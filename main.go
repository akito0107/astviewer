// +build js,wasm

package main

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"syscall/js"

	"github.com/k0kubun/pp"
)

func main() {
	gosrc := js.Global().Get("document").Call("getElementById", "gosrc")
	src := gosrc.Get("value").String()

	f, err := parser.ParseFile(token.NewFileSet(), "", src, parser.ParseComments)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	goout := js.Global().Get("document").Call("getElementById", "goout")
	var buf bytes.Buffer
	// var depth int
	scheme := pp.ColorScheme{
		Bool:            pp.NoColor,
		Integer:         pp.NoColor,
		Float:           pp.NoColor,
		String:          pp.NoColor,
		StringQuotation: pp.NoColor,
		EscapedChar:     pp.NoColor,
		FieldName:       pp.NoColor,
		PointerAdress:   pp.NoColor,
		Nil:             pp.NoColor,
		Time:            pp.NoColor,
		StructName:      pp.NoColor,
		ObjectLength:    pp.NoColor,
	}
	pp.SetColorScheme(scheme)
	pp.Fprintf(&buf, "%+v", f)
	// ast.Inspect(f, func(n ast.Node) bool {
	// 	if n != nil {
	// 		depth--
	// 	} else {
	// 		depth++
	// 		return true
	// 	}
	// 	if _, err := fmt.Fprintf(&buf, "%*s%#v\n", depth*2, "", n); err != nil {
	// 		log.Printf("%+v", err)
	// 		return false
	// 	}
	// 	return true
	// })
	goout.Set("innerHTML", buf.String())
}
