package analyzer

import (
	"go/ast"
	"go/token"
	"os"
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
	Doc:  Doc,
	Run:  run,
}

var allowedLoggers = map[string]bool{
	"log":  true,
	"slog": true,
	"zap":  true,
}

func run(pass *analysis.Pass) (any, error) {
	configPath := os.Getenv("LINTERLOG_CONFIG")

	if configPath == "" {
		configPath = "../.linterlog.yaml" 
	}

	cfg, err := load_config(configPath)
	if err != nil {
		cfg = default_config()
	}
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			selector, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			obj, ok := selector.X.(*ast.Ident)
			if !ok {
				return true
			}

			if allowedLoggers[obj.Name] && len(call.Args) > 0 {
				checkLogArgs(pass, call, cfg)
			}

			return true
		})
	}

	return nil, nil
}

func checkLogArgs(pass *analysis.Pass, call *ast.CallExpr, cfg *Config) {
	firstArg := call.Args[0]

	switch arg := firstArg.(type) {
	case *ast.BasicLit:
		if arg.Kind == token.STRING {
			text := strings.Trim(arg.Value, "\"`")
			checkRules(pass, arg, text, cfg)
		}

	case *ast.BinaryExpr:
		if arg.Op == token.ADD {
			extractAndCheckStrings(pass, arg, cfg)
		}

	case *ast.CallExpr:
		if sel, ok := arg.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "Sprintf" {
			if len(arg.Args) > 0 {
				if format, ok := arg.Args[0].(*ast.BasicLit); ok && format.Kind == token.STRING {
					text := strings.Trim(format.Value, "\"`")
					checkRules(pass, format, text, cfg)
				}
			}
		}
	}
}

func extractAndCheckStrings(pass *analysis.Pass, expr *ast.BinaryExpr, cfg *Config) {
	var collect func(e ast.Expr)
	collect = func(e ast.Expr) {
		switch x := e.(type) {
		case *ast.BasicLit:
			if x.Kind == token.STRING {
				text := strings.Trim(x.Value, "\"`")
				checkRules(pass, x, text, cfg)
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
