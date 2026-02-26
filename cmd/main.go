package main

import (
	"linter_project/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	// Запускаем анализатор как отдельную программу
	singlechecker.Main(analyzer.Analyzer)
}