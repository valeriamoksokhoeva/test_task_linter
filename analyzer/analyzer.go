package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const Doc = "check log messages for rules: \n" +
	"1. first letter must be lowercase\n" +
	"2. must be in English\n" +
	"3. no special characters or emojis\n" +
	"4. no sensitive data"

var Analyzer = &analysis.Analyzer{
	Name: "linterlog",
	Doc: Doc,
	Run: run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// Проверяем, является ли узел вызовом функции
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			// Проверяем, что это вызов метода (например, log.Info)
			selector, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			// Получаем объект (левую часть) - например "log" из "log.Info"
			obj, ok := selector.X.(*ast.Ident)
			if !ok {
				return true
			}

			if obj.Name == "log" || obj.Name == "slog" || obj.Name == "zap"{
				if len(call.Args) > 0 {
					checkLogArgs(pass, call)
				}
			}

			return true
		})
	}

	return nil, nil
}


func checkLogArgs(pass *analysis.Pass, call *ast.CallExpr) {
    switch expr := call.Args[0].(type) {
    case *ast.BasicLit:
        if expr.Kind == token.STRING {
            text := strings.Trim(expr.Value, "\"`")
            CheckRules(pass, call.Pos(), text)
        }
        
    case *ast.BinaryExpr:
        if expr.Op == token.ADD {
            extractAndCheckStrings(pass, expr, call.Pos())
        }
        
    case *ast.CallExpr:
        if sel, ok := expr.Fun.(*ast.SelectorExpr); ok {
            if sel.Sel.Name == "Sprintf" {
                if len(expr.Args) > 0 {
                    if format, ok := expr.Args[0].(*ast.BasicLit); ok {
                        if format.Kind == token.STRING {
                            text := strings.Trim(format.Value, "\"`")
                            CheckRules(pass, call.Pos(), text)
                        }
                    }
                }
            }
        }
    }
}

func extractAndCheckStrings(pass *analysis.Pass, expr *ast.BinaryExpr, pos token.Pos) {
    var collect func(e ast.Expr)
    collect = func(e ast.Expr) {
        switch x := e.(type) {
        case *ast.BasicLit:
            if x.Kind == token.STRING {
                text := strings.Trim(x.Value, "\"`")
                CheckRules(pass, pos, text)
            }
        case *ast.BinaryExpr:
            if x.Op == token.ADD {
                collect(x.X)
                collect(x.Y)
            }
        }
    }
    
    collect(expr)
}