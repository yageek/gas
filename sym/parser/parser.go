package parser

import (
	"errors"
	"fmt"
	"github.com/yageek/gas/sym/scanner"
	"github.com/yageek/gas/sym/token"
)

var lineCount int = 0

type Node interface {
	Parent() Node
	Lhs() Node
	Rhs() Node
}

func StringAST(n Node) string {
	fmt.Println("", n)
	fmt.Println("/\\")

	var lhs string = ""
	var rhs string = ""
	if n.Lhs() != nil {
		lhs = StringAST(n.Lhs())
	}
	if n.Rhs() != nil {
		rhs = StringAST(n.Lhs())
	}
	return fmt.Sprintf("%v \t %v \n", lhs, rhs)
}

type NumberNode struct {
	ParentNode  Node
	NumberValue string
}

func (n *NumberNode) Init(value string) *NumberNode {
	n.NumberValue = value
	return n
}
func (n *NumberNode) Parent() Node {
	return n.ParentNode
}
func (n *NumberNode) Rhs() Node {
	return nil
}
func (n *NumberNode) Lhs() Node {
	return nil
}

func (n *NumberNode) String() string {
	return n.NumberValue
}

type ListNode struct {
	ParentNode Node
	Operator   token.Token
	rhs, lhs   Node
}

func (n *ListNode) Init(t token.Token) *ListNode {
	n.Operator = t
	return n
}
func (n *ListNode) Parent() Node {
	return n.ParentNode
}
func (n *ListNode) Rhs() Node {
	return n.rhs
}
func (n *ListNode) Lhs() Node {
	return n.lhs
}
func (n *ListNode) String() string {
	return n.Operator.String()
}

type Parser struct {
	scanner     *scanner.Scanner
	itemStack   []Node
	root        Node
	currentNode Node
}

func (p *Parser) next() scanner.TokenItem {
	item := p.scanner.NextItem()

	return item
}
func (p *Parser) push(node Node) {
	p.itemStack = append(p.itemStack, node)
}

func (p *Parser) pop() *Node {
	size := len(p.itemStack)

	if size == 0 {
		return nil
	}

	last := p.itemStack[size-1]
	p.itemStack = append(p.itemStack[:size-1], p.itemStack[0:]...)

	return &last
}

func Init(scanner *scanner.Scanner) *Parser {

	parser := &Parser{
		scanner:     scanner,
		root:        nil,
		currentNode: nil,
	}

	return parser
}

func (p *Parser) Parse() (Node, error) {

	return p.parseList(nil)

}

func (p *Parser) parseList(rhs *ListNode) (Node, error) {
	item := p.next()
	if item.Token != token.NUMBER {
		return nil, errors.New("Expected a Number")
	}

	numberNode := new(NumberNode).Init(item.Value)

	item = p.next()
	switch item.Token {
	case token.ADD, token.SUB, token.MUL, token.QUO:
		node := new(ListNode)
		node.Operator = item.Token
		node.ParentNode = rhs
		node.lhs = numberNode

		var err error
		node.rhs, err = p.parseList(node)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		return node, nil
	default:
		return numberNode, nil
	}

}
