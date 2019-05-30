package rewrite

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"hw2/expr"
	"hw2/simplify"
	"fmt"
	"strconv"
)

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) {
	ast.Inspect(node, func(fn ast.Node) bool{
		switch x := fn.(type){
		case *ast.CallExpr:
			//if not expr.ParseAndEval then end
			switch fun := x.Fun.(type){
			case *ast.Ident:
				if(fun.Name == "ParseAndEval"){
					return false
				}
			}
			lencheck := len(x.Args)
			if(lencheck != 2){
				return false
			}
			switch fun := x.Args[0].(type){
			case *ast.BasicLit:
				q,_ := strconv.Unquote(fun.Value)
				if(string(q[0])=="%"){
					return false
				}
				temp,_:=expr.Parse(q)
				if(temp == nil){
					return false
				}
				simp := simplify.Simplify(temp,nil)
				final := convert(simp)
				final = strconv.Quote(final)
				fun.Value = final
			}
		}
		return true
	})
	//panic("TODO: implement this!")
}

func SimplifyParseAndEval(src string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	rewriteCalls(f)

	var buf bytes.Buffer
	format.Node(&buf, fset, f)
	return buf.String()
}
func convert(i expr.Expr) string{
		cow:= i.Eval(nil)
		var k int = int(cow)
		final:= fmt.Sprintf("%d",k)

		return final
}