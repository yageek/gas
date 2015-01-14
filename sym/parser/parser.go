package parser

import (
	"github.com/yageek/gas/sym/scanner"
	"github.com/yageek/gas/sym/token"
)

type Parser struct {
	scanner   *scanner.Scanner
	itemStack []scanner.TokenItem
	root      interface{}
	current   interface{}
}

func (p *Parser) push(item scanner.TokenItem) {
	p.itemStack = append(p.itemStack, item)
}

func (p *Parser) pop() *scanner.TokenItem {
	size := len(p.itemStack)

	if size == 0 {
		return nil
	}

	last := p.itemStack[size-1]
	p.itemStack = append(p.itemStack[:size-2], p.itemStack[0:]...)

	return &last
}

type ExprNode struct {
	Operator token.Token
	Rhs, Lhs *ExprNode
}

func Init(scanner *scanner.Scanner) *Parser {
	root := &ExprNode{}
	parser := &Parser{
		scanner: scanner,
		root:    root,
		current: root,
	}

	return parser
}

func (p *Parser) Parse() {
Loop:
	for {
		item := p.scanner.NextItem()
		switch item.Token {

		case token.ADD, token.SUB, token.MUL, token.QUO:

		case token.EOF:
			break Loop

		default:
			p.push(item)
		}
	}
}

func (p *Parser) operator(t token.Token) {

}
