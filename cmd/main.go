package main

import (
	"github.com/valeriamoksokhoeva/test_task_linter/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
