package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/packages"
)

type keepUsesInfo struct {
	enum
	startPosition token.Pos
}

func check(enums []enum, filepaths []string) error {
	debugf("filepaths: %+v\n", filepaths)
	config := &packages.Config{Mode: packages.LoadSyntax}

	for _, filepath := range filepaths {
		debugf("start check %s\n", filepath)

		pkgs, err := packages.Load(config, filepath)
		if err != nil {
			return err
		}
		debugf("pkgs: %v,\n", pkgs)

		for _, pkg := range pkgs {
			debugf("package is: %+v, and path is %+v\n", pkg.Name, pkg.PkgPath)

			info := pkg.TypesInfo
			keepUsesInfos := []keepUsesInfo{}
			for identifier, use := range info.Uses {
				namedType, ok := use.Type().(*types.Named)
				if !ok {
					continue
				}

				for _, enum := range enums {
					debugf("enum.name is %s, use enum type name is %s and start position %d\n", enum.name, namedType.Obj().Name(), identifier.Pos())
					if enum.name != namedType.Obj().Name() {
						continue
					}
					keepUsesInfos = append(keepUsesInfos, keepUsesInfo{enum: enum, startPosition: identifier.Pos()})
				}
			}
			debugf("keep Uses infos %v \n", keepUsesInfos)

			for node, scope := range info.Scopes {
				debugf("scope info: %+v\n", *scope)

				switchNode, ok := node.(*ast.SwitchStmt)
				if !ok {
					continue
				}
				debugf("switchNode is %+v\n", switchNode)

				for _, info := range keepUsesInfos {
					if !scope.Contains(info.startPosition) {
						continue
					}

					patternContainer := map[string]struct{}{}
					for _, stmt := range switchNode.Body.List {
						debugf("stmt is %v\n", stmt)
						if caseClause, ok := stmt.(*ast.CaseClause); ok {
							debugf("caseClause is %v\n", *caseClause)
							for _, caseExpr := range caseClause.List {
								debugf("caseExpr is %v, type of %v\n", caseExpr, reflect.TypeOf(caseExpr))

								func() {
									// NOTE: this scope pack pattern names from external package
									// FIXME: more explicit condition for extarct pattern name from external package
									if caseValue, ok := caseExpr.(*ast.SelectorExpr); ok {
										if packageInfo, ok := caseValue.X.(*ast.Ident); ok {
											debugf("caseValue is %v, name of %v, x type is %v\n", caseValue, caseValue.Sel.Name, packageInfo.Name)
											if packageInfo.Name == info.enum.packageName {
												for _, pattern := range info.enum.patterns {
													if pattern == caseValue.Sel.Name {
														patternContainer[pattern] = struct{}{}
													}
												}
											}
										}
									}
								}()
								func() {
									// NOTE: this scope pack pattern names from internal package
									// FIXME: more explicit condition for extarct pattern name from internal package
									if caseValue, ok := caseExpr.(*ast.Ident); ok {
										debugf("caseValue is %v, name of %v\n", caseValue, caseValue.Name)
										for _, pattern := range info.enum.patterns {
											if pattern == caseValue.Name {
												patternContainer[pattern] = struct{}{}
											}
										}
									}
								}()
							}
						}
					}
					debugf("all patterns %v\n", patternContainer)

					for _, pattern := range info.enum.patterns {
						debugf("pattern: %v\n", pattern)

						if _, ok := patternContainer[pattern]; !ok {
							position := switchNode.Switch
							file := pkg.Fset.File(position)
							line := file.Line(position)
							filepath := pkg.GoFiles[0]
							return fmt.Errorf("missing enum pattern for %s.%s.%s. at %s:%d:%d", info.enum.packageName, info.enum.name, pattern, filepath, line, position)
						}
					}
				}
			}
		}
		debugf("end check %s\n", filepath)
	}
	return nil
}
