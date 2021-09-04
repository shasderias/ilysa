package gradient

import (
	"go/ast"
	"go/parser"
	"go/token"
	"runtime"
)

var DeclaredGradients = map[string]Table{}

func registerGrad(grad Table) {
	defer func() {
		recover()
	}()
	_, f, l, _ := runtime.Caller(2)

	fset := token.NewFileSet()
	parserMode := parser.ParseComments

	a, err := parser.ParseFile(fset, f, nil, parserMode)
	if err != nil {
		return
	}

	for _, d := range a.Decls {
		gd, ok := d.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range gd.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			name := vs.Names[0].Name

			pos := fset.Position(vs.Pos())

			if pos.Line == l {
				DeclaredGradients[name] = grad
			}
		}
	}
}
