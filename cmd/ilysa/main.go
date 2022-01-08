package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	var (
		runCmd   = makeRunCmd()
		watchCmd = makeWatchCmd()
	)

	rootFlagSet := flag.NewFlagSet("root", flag.ExitOnError)

	root := &ffcli.Command{
		FlagSet:     rootFlagSet,
		ShortUsage:  "ilysa <subcommand>",
		Subcommands: []*ffcli.Command{runCmd, watchCmd},
		Exec: func(ctx context.Context, args []string) error {
			return flag.ErrHelp
		},
	}

	err := root.ParseAndRun(context.Background(), os.Args[1:])
	if err != nil && err != flag.ErrHelp {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//func doGradient(args []string) error {
//	path := "."
//	if len(args) >= 1 {
//		path = args[0]
//	}
//
//	fset := token.NewFileSet()
//	parserMode := parser.ParseComments
//
//	pkgs, err := parser.ParseDir(fset, path, nil, parserMode)
//	if err != nil {
//		return err
//	}
//
//	mainPkg, ok := pkgs["main"]
//	if !ok {
//		return fmt.Errorf("main package not found")
//	}
//
//	for _, f := range mainPkg.Files {
//		for _, d := range f.Decls {
//			gd, ok := d.(*ast.GenDecl)
//			if !ok {
//				continue
//			}
//			for _, spec := range gd.Specs {
//				vs, ok := spec.(*ast.ValueSpec)
//				if !ok {
//					continue
//				}
//				for _, val := range vs.Values {
//					ce, ok := val.(*ast.CallExpr)
//					if !ok {
//						continue
//					}
//					//fmt.Printf("%+v %T\n", ce.Fun, ce.Fun)
//					se, ok := ce.Fun.(*ast.SelectorExpr)
//					if !ok {
//						continue
//					}
//					//fmt.Printf("%+v %T\n", se.X, se.X)
//					sex, ok := se.X.(*ast.Ident)
//					if !ok {
//						continue
//					}
//					fmt.Println(sex.Name, se.Sel.Name)
//					if sex.Name == "gradient" && se.Sel.Name == "New" {
//
//					}
//				}
//			}
//		}
//	}
//
//	return nil
//}
