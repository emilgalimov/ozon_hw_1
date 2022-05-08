package functionfrequency

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
)

type Sorted struct {
	name  string
	count int
}

func FunctionFrequency(gocode []byte) []string {

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "", gocode, 0)
	if err != nil {
		panic(err)
	}
	functions := make(map[string]int)

	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			switch xx := x.Fun.(type) {
			case *ast.Ident:
				functions[xx.Name]++
			case *ast.SelectorExpr:
				switch xxx := xx.X.(type) {
				case *ast.Ident:
					functions[xxx.Name+"."+xx.Sel.Name]++
				}
			}
		}
		return true
	})

	var sortedSlice []Sorted

	for key, val := range functions {
		sortedSlice = append(sortedSlice, Sorted{key, val})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].count > sortedSlice[j].count
	})

	var result []string

	for i := 0; i < 3; i++ {
		result = append(result, sortedSlice[i].name)
	}

	return result
}
