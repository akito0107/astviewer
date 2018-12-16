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

type ASTView struct {
	Label  string
	Value  string
	LineNo int
}

func (a *ASTView) String() string {
	return fmt.Sprintf("%s: { %s }", a.Label, a.Value)
}

func main() {
	gosrc := js.Global().Get("document").Call("getElementById", "gosrc")
	src := gosrc.Get("value").String()

	goout := js.Global().Get("document").Call("getElementById", "goout")
	var buf bytes.Buffer
	var depth int

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	ast.Inspect(f, func(n ast.Node) bool {
		if n != nil {
			depth--
		} else {
			depth++
			return true
		}
		lineno := fset.Position(n.Pos()).Line
		view := &ASTView{
			LineNo: lineno,
		}
		switch x := n.(type) {
		case *ast.ArrayType:
			view.Label = "ArrayType"
			view.Value = fmt.Sprintf("Len: %s, Elt: %v", x.Len, x.Elt)
		case *ast.AssignStmt:
			view.Label = "AssignStmt"
			view.Value = fmt.Sprintf("Lhs: %v, Tok: %d, Rhs: %v", x.Lhs, x.Tok, x.Rhs)
		case *ast.BadDecl:
			view.Label = "BadDecl"
		case *ast.BadExpr:
			view.Label = "BadExpr"
		case *ast.BadStmt:
			view.Label = "BadStmt"
		case *ast.BasicLit:
			view.Label = "BasicLit"
			view.Value = fmt.Sprintf("Kind: %d, Value: %s", x.Kind, x.Value)
		case *ast.BinaryExpr:
			view.Label = "BinaryExpr"
			view.Value = fmt.Sprintf("X: %v, Op: %d, Y: %v", x.X, x.Op, x.Y)
		case *ast.BlockStmt:
			view.Label = "BlockStmt"
			view.Value = fmt.Sprintf("List: %v", x.List)
		case *ast.BranchStmt:
			view.Label = "BranchStmt"
			view.Value = fmt.Sprintf("Tok: %d, Label: %v", x.Tok, x.Label)
		case *ast.CallExpr:
			view.Label = "CallExpr"
			view.Value = fmt.Sprintf("Fun: %v, Args: %v", x.Fun, x.Args)
		case *ast.CaseClause:
			view.Label = "CaseClause"
			view.Value = fmt.Sprintf("List: %v, Body: %v", x.List, x.Body)
		case *ast.ChanType:
			view.Label = "ChanType"
			view.Value = fmt.Sprintf("Dir: %v, Value: %v", x.Dir, x.Value)
		case *ast.CommClause:
			view.Label = "CommClause"
			view.Value = fmt.Sprintf("Comm: %v, Body: %v", x.Comm, x.Body)
		case *ast.Comment:
			view.Label = "Comment"
			view.Value = fmt.Sprintf("Text: %s", x.Text)
		case *ast.CommentGroup:
			view.Label = "CommentGroup"
			view.Value = fmt.Sprintf("List: %v", x.List)
		case *ast.CompositeLit:
			view.Label = "CompositeLit"
			view.Value = fmt.Sprintf("Type: %v, Elts: %v, Incomplete: %v", x.Type, x.Elts, x.Incomplete)
		case *ast.DeclStmt:
			view.Label = "DeclStmt"
			view.Value = fmt.Sprintf("Decl: %v", x.Decl)
		case *ast.DeferStmt:
			view.Label = "DeferStmt"
			view.Label = fmt.Sprintf("Call: %+v", x.Call)
		case *ast.Ellipsis:
			view.Label = "Ellipsis"
			view.Value = fmt.Sprintf("Elt: %v", x.Elt)
		case *ast.EmptyStmt:
			view.Label = "EmptyStmt"
		case *ast.ExprStmt:
			view.Label = "ExprStmt"
			view.Value = fmt.Sprintf("X: %v", x.X)
		case *ast.Field:
			view.Label = "Field"
			view.Value = fmt.Sprintf("Doc: %v, Names: %s, Type: %s, Tag: %v, Comment: %v", x.Doc, x.Names, x.Type, x.Tag, x.Comment)
		case *ast.FieldList:
			view.Label = "FieldList"
			view.Value = fmt.Sprintf("List: %v", x.List)
		case *ast.File:
			view.Label = "File"
		case *ast.ForStmt:
			view.Label = "ForStmt"
			view.Value = fmt.Sprintf("Init: %v, Cond: %v, Post: %v, Body: %v", x.Init, x.Cond, x.Post, x.Body)
		case *ast.FuncDecl:
			view.Label = "FuncDecl"
			view.Value = fmt.Sprintf("Doc: %v, Recv: %v, Name: %s, Type: %v, Body: %v", x.Doc, x.Recv, x.Name, x.Type, x.Body)
		case *ast.FuncLit:
			view.Label = "FuncLit"
			view.Value = fmt.Sprintf("Type: %v, Body: %v", x.Type, x.Body)
		case *ast.FuncType:
			view.Label = "FuncType"
			view.Value = fmt.Sprintf("Param: %+v, Results: %+v", x.Params, x.Results)
		case *ast.GenDecl:
			view.Label = "GenDecl"
			view.Value = fmt.Sprintf("Doc: %+v, Tok: %d, Specs: %+v", x.Doc, x.Tok, x.Specs)
		case *ast.GoStmt:
			view.Label = "GoStmt"
			view.Value = fmt.Sprintf("Call: %+v", x.Call)
		case *ast.Ident:
			view.Label = "Ident"
			view.Value = fmt.Sprintf("Name: %s, Obj: %+v", x.Name, x.Obj)
		case *ast.IfStmt:
			view.Label = "IfStmt"
			view.Value = fmt.Sprintf("Init: %#v, Cond: %#v, Body: %#v, Else: %#v", x.Init, x.Cond, x.Body, x.Else)
		case *ast.ImportSpec:
			view.Label = "ImportSpec"
			view.Value = fmt.Sprintf("Doc: %#v, Name: %#v, Path: %#v, Comment: %#v", x.Doc, x.Name, x.Path, x.Comment)
		case *ast.IncDecStmt:
			view.Label = "IncDecStmt"
			view.Value = fmt.Sprintf("X: %#v, Tok: %d", x.X, x.Tok)
		case *ast.IndexExpr:
			view.Label = "IndexExpr"
			view.Value = fmt.Sprintf("X: %#v, Index: %#v", x.X, x.Index)
		case *ast.InterfaceType:
			view.Label = "InterfaceType"
			view.Value = fmt.Sprintf("Methods: %#v, Incomplete: %v", x.Methods, x.Incomplete)
		case *ast.KeyValueExpr:
			view.Label = "KeyValueExpr"
			view.Value = fmt.Sprintf("Key: %#v, Value: %#v", x.Key, x.Value)
		case *ast.LabeledStmt:
			view.Label = "LabeledStmt"
			view.Value = fmt.Sprintf("Label: %#v, Stmt: %#v", x.Label, x.Stmt)
		case *ast.MapType:
			view.Label = "MapType"
			view.Label = fmt.Sprintf("Key: %#v, Value: %#v", x.Key, x.Value)
		case *ast.Package:
			view.Label = "Package"
			view.Value = fmt.Sprintf("Name: %s, Scope: %#v, Imports: %#v, Files: %#v", x.Name, x.Scope, x.Imports, x.Files)
		case *ast.ParenExpr:
			view.Label = "ParenExpr"
			view.Value = fmt.Sprintf("X: %#v", x.X)
		case *ast.RangeStmt:
			view.Label = "RangeStmt"
			view.Value = fmt.Sprintf("Key: %#v, Value: %#v, Tok: %d, X: %#v, Body: %#v", x.Key, x.Value, x.Tok, x.X, x.Body)
		case *ast.ReturnStmt:
			view.Label = "ReturnStmt"
			view.Value = fmt.Sprintf("Results: %#v", x.Results)
		case *ast.SelectStmt:
			view.Label = "SelectStmt"
			view.Value = fmt.Sprintf("Body: %#v", x.Body)
		case *ast.SelectorExpr:
			view.Label = "SelectorExpr"
			view.Value = fmt.Sprintf("X: %#v, Sel: %#v", x.X, x.Sel)
		case *ast.SendStmt:
			view.Label = "SendStmt"
			view.Value = fmt.Sprintf("Chan: %#v, Arrow: %d, Value: %#v", x.Chan, x.Arrow, x.Value)
		case *ast.SliceExpr:
			view.Label = "SliceExpr"
			view.Value = fmt.Sprintf("X: %#v, Low: %#v, High: %#v, Max: %#v, Slice3: %v", x.X, x.Low, x.High, x.Max, x.Slice3)
		case *ast.StarExpr:
			view.Label = "StarExpr"
			view.Value = fmt.Sprintf("X: %#v", x.X)
		case *ast.StructType:
			view.Label = "StructType"
			view.Value = fmt.Sprintf("Fields: %#v, Incomplete: %v", x.Fields, x.Incomplete)
		case *ast.SwitchStmt:
			view.Label = "SwitchStmt"
			view.Value = fmt.Sprintf("Init: %#v, Tag: %#v, Body: %#v", x.Init, x.Tag, x.Body)
		case *ast.TypeAssertExpr:
			view.Label = "TypeAssertExpr"
			view.Value = fmt.Sprintf("X: %#v, Type: %#v", x.X, x.Type)
		case *ast.TypeSpec:
			view.Label = "TypeSpec"
			view.Value = fmt.Sprintf("Doc: %#v, Name: %#v, Type: %#v, Comment: %#v", x.Doc, x.Name, x.Type, x.Comment)
		case *ast.TypeSwitchStmt:
			view.Label = "TypeSwitchStmt"
			view.Value = fmt.Sprintf("Init: %#v, Assign: %#v, Body: %#v", x.Init, x.Assign, x.Body)
		case *ast.UnaryExpr:
			view.Label = "UnaryExpr"
			view.Value = fmt.Sprintf("Op: %d, X: %#v", x.Op, x.X)
		case *ast.ValueSpec:
			view.Label = "ValueSpec"
			view.Value = fmt.Sprintf("Doc: %#v, Names: %#v, Type: %#v, Values: %#v, Comment: %#v", x.Doc, x.Names, x.Type, x.Values, x.Comment)
		default:
			view.Label = "Unknown"
		}

		if _, err := fmt.Fprintf(&buf, "<span class=\"astline lineno%d\">%*s%s</span>\n", view.LineNo, depth*2, "", view); err != nil {
			log.Printf("%+v", err)
			return false
		}
		return true
	})
	goout.Set("innerHTML", buf.String())
}
