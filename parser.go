package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

type enum struct {
	name        string
	packageName string
	patterns    []string
}

func parse(filepaths []string) []enum {
	enums := []enum{}
	for _, filepath := range filepaths {
		e := enum{}

		astFile := parseASTFile(filepath)
		e.packageName = astFile.Name.Name

		ast.Inspect(astFile, func(node ast.Node) bool {
			// end inspect
			if node == nil {
				return false
			}
			// parsed enum declearation
			decl, ok := node.(*ast.GenDecl)
			if !ok {
				return true
			}
			if decl.Tok != token.CONST {
				return true
			}

			patterns := []string{}
			for _, spec := range decl.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}

				// e.g) golang language = iota
				isFirstDeclearation := valueSpec.Type != nil && len(valueSpec.Values) > 0
				if isFirstDeclearation {
					ident, ok := valueSpec.Type.(*ast.Ident)
					if !ok {
						continue
					}
					e.name = ident.Name
				}

				for _, name := range valueSpec.Names {
					if name.Name == "_" {
						continue
					}
					patterns = append(patterns, name.Name)
				}
			}
			e.patterns = patterns

			enums = append(enums, e)

			return true
		})
	}
	return enums
}

func parseASTFile(filepath string) *ast.File {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, filepath, buf, 0)
	if err != nil {
		panic(err)
	}
	return astFile
}
