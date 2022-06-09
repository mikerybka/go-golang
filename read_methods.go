package golang

import "go/ast"

func ReadMethods(pkg *ast.Package, typeID string) ([]*ast.FuncDecl, error) {
	methods := []*ast.FuncDecl{}
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok {
				if decl.Recv != nil && decl.Recv.List != nil && len(decl.Recv.List) == 1 {
					if ident, ok := decl.Recv.List[0].Type.(*ast.Ident); ok {
						if ident.Name == typeID {
							methods = append(methods, decl)
						}
					}
				}
			}
		}
	}
	return methods, nil
}
