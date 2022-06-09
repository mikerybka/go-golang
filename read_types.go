package golang

import (
	"go/ast"
	"go/token"
)

func ReadTypes(pkg *ast.Package) (map[string]*ast.TypeSpec, error) {
	types := make(map[string]*ast.TypeSpec)
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			switch decl := decl.(type) {
			case *ast.GenDecl:
				if decl.Tok == token.TYPE {
					for _, spec := range decl.Specs {
						switch spec := spec.(type) {
						case *ast.TypeSpec:
							types[spec.Name.Name] = spec
						}
					}
				}
			}
		}
	}
	return types, nil
}
