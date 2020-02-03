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

// parse traverse go file and these map to []enum.
// ast strcture can be confirmed from here: https://play.golang.org/p/aitWi-5RoHj
func parse(filepaths []string) []enum {

	type enumDeclInfo struct {
		packagePath string
		typeName    string
	}

	// NOTE: First, cache ast file
	astFiles := []*ast.File{}
	for _, path := range filepaths {
		astFile := parseASTFile(path)
		astFiles = append(astFiles, astFile)
	}

	// NOTE: Second, collect enum GenDecl Information
	declInfos := []enumDeclInfo{}
	for i, path := range filepaths {
		astFile := astFiles[i]
		ast.Inspect(astFile, func(node ast.Node) bool {
			f, ok := node.(*ast.File)
			if !ok {
				// NOTE: It is support only global definition enum
				return false
			}

			for _, decl := range f.Decls {
				decl, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}
				for _, spec := range decl.Specs {
					spec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					if identifier, ok := spec.Type.(*ast.Ident); ok {
						declInfos = append(declInfos, enumDeclInfo{
							packagePath: filepath.Dir(path),
							typeName:    identifier.Name,
						})
					}
				}
			}
			return true
		})
	}

	// NOTE: Third, collect enum information with matched GenDecl informations.
	enums := []enum{}
	for i, path := range filepaths {
		e := enum{}

		astFile := astFiles[i]
		e.packageName = astFile.Name.Name
		e.packagePath = filepath.Dir(path)

		ast.Inspect(astFile, func(node ast.Node) bool {
			f, ok := node.(*ast.File)
			if !ok {
				// NOTE: It is support only global definition enum
				return false
			}

			for _, decl := range f.Decls {
				decl, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}
				if decl.Tok != token.CONST {
					continue
				}

				patterns := []string{}
				for _, spec := range decl.Specs {
					valueSpec, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}

					// NOTE: e.g) golang language = iota
					isFirstDeclearation := valueSpec.Type != nil && len(valueSpec.Values) > 0
					if isFirstDeclearation {
						ident, ok := valueSpec.Type.(*ast.Ident)
						if !ok {
							continue
						}
						for _, declInfo := range declInfos {
							if declInfo.packagePath != e.packagePath {
								continue
							}
							if declInfo.typeName != ident.Name {
								continue
							}
						}
						e.name = ident.Name
					}

					if e.name == "" {
						debugf("e.name is empty. will patterns is %+v\n", valueSpec.Names)
						continue
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
			}
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
