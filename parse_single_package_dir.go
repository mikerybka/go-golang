package golang

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
)

func ParseSinglePackageDir(dir string, filter func(fs.FileInfo) bool) (*ast.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, filter, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parse dir: %v", err)
	}

	if len(pkgs) != 1 {
		return nil, errors.New("expected one package in dir")
	}

	var pkg *ast.Package
	for _, pkg = range pkgs {
		break
	}
	return pkg, nil
}
