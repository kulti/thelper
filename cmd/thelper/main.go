// Package main provides the thelper command line tool.
package main

import (
	"github.com/kulti/thelper/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.NewAnalyzer())
}
