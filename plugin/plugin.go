package plugin 

import (
    "github.com/golangci/golangci-lint/pkg/goanalysis"
    "golang.org/x/tools/go/analysis"
    
    "github.com/valeriamoksokhoeva/test_task_linter/analyzer" 
)

func New() *goanalysis.Linter {
    return goanalysis.NewLinter(
        analyzer.Analyzer.Name,
        analyzer.Analyzer.Doc,
        []*analysis.Analyzer{analyzer.Analyzer},
        nil,
    )
}