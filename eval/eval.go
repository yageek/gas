package eval

import (
	"errors"
	"strconv"

	"github.com/yageek/gas/sym/parser"
	"github.com/yageek/gas/sym/token"
)

type Evaluator struct {
}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func (e *Evaluator) Eval(node parser.Node) (int64, error) {

	if node == nil {
		return 0, errors.New("Incorrect AST !")
	}

	switch node.(type) {

	case *parser.NumberNode:
		v, _ := node.(*parser.NumberNode)
		return strconv.ParseInt(v.NumberValue, 10, 32)

	case *parser.OperatorNode:
		v, _ := node.(*parser.OperatorNode)
		v1, _ := e.Eval(node.Lhs())
		v2, _ := e.Eval(node.Rhs())

		switch v.Operator {

		case token.ADD:
			return v1 + v2, nil
		case token.SUB:
			return v1 - v2, nil
		case token.MUL:
			return v1 * v2, nil
		case token.QUO:
			return v1 / v2, nil
		}
	}

	return 0, errors.New("Incorrect AST")
}
