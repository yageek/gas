package parser

import (
	"errors"
	"fmt"
	"github.com/yageek/gas/sym/scanner"
	"github.com/yageek/gas/sym/token"
	"testing"
)

func expectNumberValue(node Node, value string) error {
	number, ok := node.(*NumberNode)
	if !ok {
		return errors.New("Node is node a NodeNumber")
	}

	if number.NumberValue != value {
		return errors.New(fmt.Sprintf("Node has unexpected value %s (Expected %s)", number.NumberValue, value))
	}
	return nil
}

func expectOperator(node Node, t token.Token) error {

	v, ok := node.(*OperatorNode)

	if !ok {
		return errors.New("Node is node a OperatorNode")
	}

	if v.Operator != t {
		return errors.New(fmt.Sprintf("Node has unexpected operator %v (Expected %v)", v.Operator, t))
	}

	return nil
}

func TestSimpleAddParser(t *testing.T) {
	expr := "2+2"

	scanner := scanner.Init(expr)

	parser := Init(scanner)

	node, err := parser.Parse()

	if err != nil {
		t.Error(err)
		return
	}

	err = expectOperator(node, token.ADD)
	if err != nil {
		t.Error(err)
	}

	err = expectNumberValue(node.Lhs(), "2")
	if err != nil {
		t.Error(err)
	}

	err = expectNumberValue(node.Rhs(), "2")
	if err != nil {
		t.Error(err)
	}
}

func TestSimpleSubParser(t *testing.T) {
	expr := "2-2"

	scanner := scanner.Init(expr)

	parser := Init(scanner)

	node, err := parser.Parse()

	if err != nil {
		t.Error(err)
		return
	}
	err = expectOperator(node, token.SUB)
	if err != nil {
		t.Error(err)
	}

	err = expectNumberValue(node.Lhs(), "2")
	if err != nil {
		t.Error(err)
	}

	err = expectNumberValue(node.Rhs(), "2")
	if err != nil {
		t.Error(err)
	}
}

func TestSimpleMixParser(t *testing.T) {
	expr := "2-2+4"

	scanner := scanner.Init(expr)

	parser := Init(scanner)

	node, err := parser.Parse()

	if err != nil {
		t.Error(err)
		return
	}
	err = expectOperator(node, token.SUB)
	if err != nil {
		t.Error(err)
	}

	err = expectNumberValue(node.Lhs(), "2")
	if err != nil {
		t.Error(err)
	}

	err = expectOperator(node.Rhs(), token.ADD)
	if err != nil {
		t.Error(err)
	}

	err = expectNumberValue(node.Rhs().Lhs(), "2")
	if err != nil {
		t.Error(err)
	}

	err = expectNumberValue(node.Rhs().Rhs(), "4")
	if err != nil {
		t.Error(err)
	}

}
