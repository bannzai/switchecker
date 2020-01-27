package main

import (
	"errors"
	"fmt"
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

func check(enums []enum, filepath string) error {
	fileSet := token.NewFileSet()
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("target filepath: %s read error %w", filepath, err)
	}

	astFile, err := parser.ParseFile(fileSet, filepath, bytes, 0)
	debugf("file set is %+v\n", fileSet)
	if err != nil {
		return err
	}
	info := types.Info{
		Uses:   make(map[*ast.Ident]types.Object),
		Scopes: make(map[ast.Node]*types.Scope),
	}

	var conf types.Config
	conf.Importer = importer.Default()
	pkg, err := conf.Check(filepath, fileSet, []*ast.File{astFile}, &info)
	debugf("types.Info.Uses of %+v. and scopes is %+v\n", info.Uses, info.Scopes)
	e := types.Error{}
	if errors.As(err, &e) {
		debugf("Maybe import is incomplete with %+v\n", e)
	} else if err != nil {
		return err
	}
	debugf("Package  %q\n", pkg.Path())
	debugf("Name:    %s\n", pkg.Name())
	debugf("Imports: %s\n", pkg.Imports())
	debugf("Scope:   %s\n", pkg.Scope())

	infos := []checkInfo{}

	for _, enum := range enums {
		for identifier, use := range info.Uses {
			debugf("use check info: %+v\n", use)
			packageName := ""
			if pkg := use.Pkg(); pkg != nil {
				packageName = pkg.Name()
			}
			debugf("enum.packageName is %s, use packageName is %s\n", enum.packageName, packageName)
			if enum.packageName != packageName {
				continue
			}
			t := use.Type()
			if t == nil {
				continue
			}

			// NOTE: it expected namedType.Obj().Name() is "language" if below statement.
			/*
				type language int
				const (
					golang language = iota
					swift
				)
			*/
			namedType, ok := t.(*types.Named)
			if !ok {
				continue
			}
			debugf("enum.name is %s, use enum type name is %s and start position %d\n", enum.name, namedType.Obj().Name(), identifier.Pos())
			if enum.name != namedType.Obj().Name() {
				continue
			}
			infos = append(infos, checkInfo{enum: enum, startPosition: identifier.Pos()})
		}
	}

	debugf("infos: %+v\n", infos)

	for node, scope := range info.Scopes {
		debugf("scope info: %+v\n", *scope)

		switchNode, ok := node.(*ast.SwitchStmt)
		if !ok {
			continue
		}
		debugf("switchNode is %+v\n", switchNode)
		for _, info := range infos {
			if !scope.Contains(info.startPosition) {
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
					position := switchNode.Switch
					file := fileSet.File(position)
					line := file.Line(position)
					return fmt.Errorf("missing enum pattern for %s.%s.%s. at %s:%d:%d", info.enum.packageName, info.enum.name, pattern, filepath, line, position)
				}
			}
		}
	}

	return nil
}
