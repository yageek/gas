package eval

import (
	"github.com/yageek/gas/sym/parser"
	"github.com/yageek/gas/sym/scanner"
	"testing"
)

func testEval(value int64, expr string, t *testing.T) {
	scanner := scanner.Init(expr)

	parser := parser.Init(scanner)

	node, err := parser.Parse()

	if err != nil {
		t.Error(err)
		return
	}

	evaluator := NewEvaluator()

	result, err := evaluator.Eval(node)

	if err != nil {
		t.Error(err)
		return
	}

	if result != value {
		t.Errorf("Expr: %s | Expected: %d |Result :%d", expr, value, result)
	}
}
func TestBasicEval(t *testing.T) {

	testEval(4, "2+2", t)
	testEval(0, "2-2", t)
	testEval(25, "20+10-10-5+2-2", t)
}
