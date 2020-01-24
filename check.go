package main

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
)

type checkInfo struct {
	enum
	startPosition token.Pos
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

	infos := []checkInfo{}

	for _, enum := range enums {
		for identifier, use := range info.Uses {
			packageName := ""
			if pkg := use.Pkg(); pkg != nil {
				packageName = pkg.Name()
			}
			if enum.packageName != packageName {
				continue
			}
			t := use.Type()
			if t == nil {
				continue
			}

			// NOTE: it expected namedType.Obj().Name() is "language"
			namedType, ok := t.(*types.Named)
			if !ok {
				continue
			}
			if enum.name != namedType.Obj().Name() {
				continue
			}
			infos = append(infos, checkInfo{enum: enum, startPosition: identifier.Pos()})
		}
	}

	for node, scope := range info.Scopes {
		for _, info := range infos {
			if !scope.Contains(info.startPosition) {
				continue
			}

			switchNode, ok := node.(*ast.SwitchStmt)
			if !ok {
				continue
			}

			// FIXME: It is difficult to tell about `switch x` ast.SwitchStmt or  `case xyz:` ast.SwitchStmt.
			patternContainer := map[string]struct{}{}
			for _, stmt := range switchNode.Body.List {
				if caseClause, ok := stmt.(*ast.CaseClause); ok {
					for _, expr := range caseClause.List {
						if caseValue, ok := expr.(*ast.Ident); ok {
							for _, pattern := range info.enum.patterns {
								if pattern == caseValue.Name {
									patternContainer[pattern] = struct{}{}
								}
							}
						}
					}
				}
			}

			for _, pattern := range info.enum.patterns {
				if _, ok := patternContainer[pattern]; !ok {
					return false
				}
			}
		}
	}

	return true
}
