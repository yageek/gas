package parser

import (
	"github.com/yageek/gas/sym/scanner"
	"testing"
)

func TestParser(t *testing.T) {
	expr := "2+2"

	scanner := scanner.Init(expr)

	parser := Init(scanner)

	parser.Parse()
}
