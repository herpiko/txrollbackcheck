package analyzer_test

import (
	"testing"

	"github.com/herpiko/txrollbackcheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()

	checker := analyzer.NewAnalyzer()
	analysistest.Run(t, testdata, checker, "tx")
}
