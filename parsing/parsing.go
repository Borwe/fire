package parsing

import (
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

	"golang.org/x/tools/go/packages"
)

func GetFunctions[Node ast.Node](node Node) []ast.FuncDecl{
	funcs := []ast.FuncDecl{}
	ast.Inspect(node, func(n ast.Node) bool {
		if n != nil {
			if fn, ok := n.(*ast.FuncDecl); ok {
				funcs = append(funcs, *fn)
			}
		}
		return true
	})
	return funcs
}

func GetFileName[Node ast.Node](fset *token.FileSet, file Node) string {
	return fset.Position(file.Pos()).Filename
}

func ParsePackage(dir string, fSet *token.FileSet) (fileAsts []*ast.File, errs[]packages.Error, err error){
	cfg := packages.Config{
		Mode: packages.NeedSyntax |
          packages.NeedDeps |
          packages.NeedImports |
          packages.NeedName |
          packages.NeedCompiledGoFiles |
          packages.NeedFiles,
		Fset: fSet,
		Tests: true,
	}

	pkgs, err := packages.Load(&cfg,dir)
	if err==nil {
		for _, pkg := range pkgs {
			fileAsts = append(fileAsts,pkg.Syntax...)
			errs = append(errs, pkg.Errors...)
		}
	}

	return fileAsts, errs, err
}

func ToBytes(astF *ast.File, fSet *token.FileSet) (result string, err error){
	var str strings.Builder
	err = printer.Fprint(&str, fSet, astF)
	return str.String(), err
}
