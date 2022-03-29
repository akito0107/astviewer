package main

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestParse(t *testing.T) {
	fset := token.NewFileSet()

	src := `
	package main

import "fmt"

func main() {
	fmt.Println("hello, 世界")
}

func f[T Stringer](xs []T) []string {
	var result []string
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}
	`
	_, err := parser.ParseFile(fset, "", src, parser.AllErrors)
	if err != nil {
		t.Errorf("parse error: %v", err)
	}
}
