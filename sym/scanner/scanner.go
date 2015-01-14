package scanner

import (
	"fmt"
	"github.com/yageek/gas/sym/token"
	"strings"
	"unicode"
	"unicode/utf8"
)

const eof = -1

type TokenItem struct {
	Token token.Token
	Value string
}

type Scanner struct {
	input string //The expression, code to parse
	start int    // start position of the item.
	pos   int    // current position in the input.
	width int    // width of the rune read from input.

	Items chan TokenItem // channel of scanned item

	state stateFn // The current state of the lexer
}

type stateFn func(*Scanner) stateFn

func Init(input string) *Scanner {
	s := &Scanner{
		input: input,
		Items: make(chan TokenItem, 2),
	}
	go s.run()
	return s
}

func (s *Scanner) run() {
	for state := exprFn; state != nil; {
		state = state(s)
	}
	close(s.Items)

}
func (s *Scanner) emit(t token.Token) {
	if t == token.EOF {
		s.Items <- TokenItem{t, t.String()}
	} else {
		s.Items <- TokenItem{t, s.input[s.start:s.pos]}
	}
	s.start = s.pos
}

func (s *Scanner) errorf(format string, args ...interface{}) stateFn {

	s.Items <- TokenItem{
		token.ILLEGAL, fmt.Sprintf(format, args...),
	}
	return nil
}
func (s *Scanner) next() rune {
	if s.pos >= len(s.input) {
		s.width = 0
		return eof
	}

	r, w := utf8.DecodeRuneInString(s.input[s.pos:])
	s.width = w
	s.pos += s.width
	return r
}

func (s *Scanner) backup() {
	s.pos -= s.width
}

func (s *Scanner) peek() rune {
	rune := s.next()
	s.backup()
	return rune
}

func (s *Scanner) ignore() {
	s.start = s.pos
}

func (s *Scanner) accept(valid string) bool {
	if strings.IndexRune(valid, s.next()) >= 0 {
		return true
	}
	s.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (s *Scanner) acceptRun(valid string) {
	for strings.IndexRune(valid, s.next()) >= 0 {
	}
	s.backup()
}

func exprFn(s *Scanner) stateFn {
Loop:
	for {
		switch r := s.next(); {
		case r == eof:
			break Loop
		case unicode.IsDigit(r):
			s.backup()
			return scanNumber
		case r == '+':
			s.emit(token.ADD)
		case r == '-':
			s.emit(token.SUB)
		case r == '/':
			s.emit(token.QUO)
		case r == '*':
			s.emit(token.MUL)
		case r == '^':
			s.emit(token.POW)
		case r == '(':
			s.emit(token.LPAREN)
		case r == ')':
			s.emit(token.RPAREN)
		default:
			s.ignore()
		}
	}

	s.emit(token.EOF)
	return nil
}

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func scanNumber(s *Scanner) stateFn {
	digits := "0123456789"

	s.acceptRun(digits)

	if s.accept(".") {
		s.acceptRun(digits)
	}

	if s.accept("eE") {
		s.accept("+-")
		s.acceptRun("0123456789")
	}

	if isAlphaNumeric(s.peek()) {
		s.next()
		return s.errorf("bad number syntax: %q",
			s.input[s.start:s.pos])
	}

	s.emit(token.NUMBER)
	return exprFn
}
