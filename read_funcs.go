package golang

import "go/ast"

func ReadFuncs(pkg *ast.Package) map[string]*ast.FuncDecl {
	funcs := make(map[string]*ast.FuncDecl)
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			switch decl := decl.(type) {
			case *ast.FuncDecl:
				if decl.Recv == nil {
					funcs[decl.Name.Name] = decl
				}
			}
		}
	}
	return funcs
}
