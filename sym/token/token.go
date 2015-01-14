package token

import (
	"strconv"
)

type Token int
type Pos int

const (
	//Specials tokens
	ILLEGAL Token = iota
	EOF
	COMMENT
	IDENT

	// Operators and delimiters
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	POW // ^

	LSS    // <
	GTR    // >
	ASSIGN // =
	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN // )
	RBRACK // ]
	RBRACE // }

	NUMBER
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	COMMENT: "COMMNENT",
	IDENT:   "IDENT",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	POW: "^",

	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",
	RPAREN: ")",
	RBRACK: "]",
	RBRACE: "}",

	NUMBER: "NUMBER",
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}
