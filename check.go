package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
)

func recursivePrintChild(scope *types.Scope) {
	numChildren := scope.NumChildren()
	if numChildren == 0 {
		return
	}
	for i := 0; i < numChildren; i++ {
		child := scope.Child(i)
		fmt.Printf("i: %d child: %v\n", i, child)
		recursivePrintChild(child)
	}
}

func check(enums []enum, filepath string) bool {
	fileSet := token.NewFileSet()
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	astFile, err := parser.ParseFile(fileSet, filepath, bytes, 0)
	if err != nil {
		panic(err)
	}
	info := types.Info{
		Types:  make(map[ast.Expr]types.TypeAndValue),
		Defs:   make(map[*ast.Ident]types.Object),
		Uses:   make(map[*ast.Ident]types.Object),
		Scopes: make(map[ast.Node]*types.Scope),
	}

	var conf types.Config
	conf.Importer = importer.Default()
	pkg, err := conf.Check(filepath, fileSet, []*ast.File{astFile}, &info)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pkg: %v\n", pkg)

	for node, scope := range info.Scopes {
		if switchNode, ok := node.(*ast.SwitchStmt); ok {
			fmt.Printf("node init: %v, tag %v, block pos start: %d end: %d, scope: %v \n", switchNode.Init, switchNode.Tag, switchNode.Body.Lbrace, switchNode.Body.Rbrace, scope)
			for i, stmt := range switchNode.Body.List {
				if caseClause, ok := stmt.(*ast.CaseClause); ok {
					for j, expr := range caseClause.List {
						if identifier, ok := expr.(*ast.Ident); ok {
							fmt.Printf("i: %d, j: %d, switch node statements: %v, identifier: %v \n", i, j, caseClause, identifier)
						}
					}
				}
			}
			recursivePrintChild(scope)
		}
	}

	for identifier, use := range info.Uses {
		pkg := use.Pkg()
		pkgs := ""
		if pkg != nil {
			pkgs = pkg.Name()
		}
		fmt.Printf("identifier: %+v, pos: %d,  use: %+v, name: %s, package name: %s\n", identifier, identifier.Pos(), use, use.Name(), pkgs)
	}
	return true
}
