package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
)

type enum struct {
	name        string
	packageName string
	packagePath string
	patterns    []string
}

func parse(filepaths []string) []enum {
	enums := []enum{}
	for _, path := range filepaths {
		e := enum{}

		astFile := parseASTFile(path)
		e.packageName = astFile.Name.Name
		e.packagePath = filepath.Dir(path)

		ast.Inspect(astFile, func(node ast.Node) bool {
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

var cachedSourceASTFile = map[string]*ast.File{}

func parseASTFile(path string) *ast.File {
	if f, ok := cachedSourceASTFile[path]; ok {
		return f
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, path, buf, 0)
	if err != nil {
		panic(err)
	}
	cachedSourceASTFile[path] = astFile
	return astFile
}
