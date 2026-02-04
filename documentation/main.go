package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"
)

func listFunctions(filePath string, shorten bool) {
	fset := token.NewFileSet()

	// Parse the file
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments+parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate through the declarations and print function names
	for _, decl := range file.Decls {
		if f, ok := decl.(*ast.FuncDecl); ok {
			// Get the doc comments for the function
			c := f.Doc
			if c == nil {
				continue
			}
			comment := strings.TrimSpace(c.Text())

			// relevant?
			if !strings.HasPrefix(comment, "Doc: ") {
				continue
			}

			// shorten?
			result := f.Name.Name
			if shorten {
				result = shortenFuncName(result)
			}
			fmt.Println(result)
		}
	}
}

func getFunctionDoc(filePath, functionName string) {
	file, err := getFileSet(filePath)
	if err != nil {
		errOut(err, 1)
	}

	// Iterate through the declarations to find the target function
	for _, decl := range file.Decls {
		if f, ok := decl.(*ast.FuncDecl); ok && f.Name.Name == functionName {
			// Get the doc comments for the function
			fmt.Println(strings.TrimPrefix(strings.TrimSpace(f.Doc.Text()), "Doc: "))
			return
		}
	}
}

type FuncParamData struct {
	Name string
	Type string
	Doc  *ast.CommentGroup
}

type FuncResultData struct {
	Name string
	Type string
	Doc  *ast.CommentGroup
}

type FuncData struct {
	TemplateName string
	FunctionName string
	Doc          string
	Params       []FuncParamData
	Results      []FuncResultData
}

func printFunctionData(filePath string) {
	result := make([]FuncData, 0)
	file, err := getFileSet(filePath)
	if err != nil {
		errOut(err, 1)
	}

	const prefix = "Doc: "

	// iterate through the relevant function declarations
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if !strings.HasPrefix(fn.Doc.Text(), prefix) {
				continue
			}

			params := make([]FuncParamData, 0, 3)
			for _, p := range fn.Type.Params.List {
				// https://stackoverflow.com/questions/50524607/go-lang-func-parameter-type/50525439#50525439

				for i := 0; i < len(p.Names); i++ {
					param := FuncParamData{
						Doc:  p.Doc,
						Name: p.Names[i].Name,
						Type: types.ExprString(p.Type),
					}
					params = append(params, param)
				}
			}

			results := make([]FuncResultData, 0, 2)
			for _, p := range fn.Type.Results.List {
				// https://stackoverflow.com/questions/50524607/go-lang-func-parameter-type/50525439#50525439

				// named?
				names := []*ast.Ident{
					{
						NamePos: 0,
						Name:    "",
						Obj:     nil,
					},
				}
				if len(p.Names) > 0 {
					names = p.Names
				}

				for i := 0; i < len(names); i++ {
					result := FuncResultData{
						Doc:  p.Doc,
						Name: names[i].Name,
						Type: types.ExprString(p.Type),
					}
					results = append(results, result)
				}
			}

			f := FuncData{
				TemplateName: shortenFuncName(fn.Name.Name),
				FunctionName: fn.Name.Name,
				Doc:          strings.TrimPrefix(strings.TrimSpace(fn.Doc.Text()), prefix),
				Params:       params,
				Results:      results,
			}
			result = append(result, f)
		}
	}

	json.NewEncoder(os.Stdout).Encode(result)
}

func errOut(err error, exitCode int) {
	fmt.Fprint(os.Stderr, err.Error())
	os.Exit(exitCode)
}

func getFileSet(filePath string) (*ast.File, error) {
	fileSet := token.NewFileSet()
	return parser.ParseFile(fileSet, filePath, nil, parser.ParseComments+parser.AllErrors)
}

func shortenFuncName(input string) string {
	return strings.TrimSuffix(input, "Func")
}

func main() {
	command := os.Args[1]

	switch command {
	case "list_funcs":
		printFunctionData(os.Args[2])
	}

	os.Exit(1)
}
