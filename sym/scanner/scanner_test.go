package scanner

import (
	"fmt"
	"github.com/yageek/gas/sym/token"
	"testing"
)

func tok(t token.Token) TokenItem {
	return TokenItem{t, t.String()}
}

func number(value string) TokenItem {
	return TokenItem{token.NUMBER, value}
}

func expectedTokenListTest(expr string, expected []TokenItem, t *testing.T) {
	scanner := Init(expr)

	var result_list []TokenItem

	for {
		item := <-scanner.Items
		result_list = append(result_list, item)
		if item.Token == token.EOF {
			break
		}
	}

	if len(expected) != len(result_list) {
		t.Errorf("Unexpected numbers of tokens\nExpected (%d) :%v \n Result (%d): %v \n", len(expected), expected, len(result_list), result_list)
		return
	}

	for index, _ := range expected {

		expected_item := expected[index]
		result_item := result_list[index]
		msg := fmt.Sprintf("Unexpected Item at index %d: Expected {%v,%v} - Result{%v, %v}\n", index, expected_item.Token, expected_item.Value, result_item.Token, result_item.Value)

		if result_item.Token != expected_item.Token {
			t.Error(msg)
		} else {
			if result_item.Token == token.NUMBER {
				if result_item.Value != expected_item.Value {
					t.Error(msg)
				}
			}
		}
	}
}

func TestAddTokenization(t *testing.T) {
	expr := "2+2+3+5"
	expected := [...]TokenItem{
		number("2"),
		tok(token.ADD),
		number("2"),
		tok(token.ADD),
		number("3"),
		tok(token.ADD),
		number("5"),
		tok(token.EOF),
	}

	expectedTokenListTest(expr, expected[:], t)
}

func TestSubTokenization(t *testing.T) {
	expr := "2-2-3-5"
	expected := [...]TokenItem{
		number("2"),
		tok(token.SUB),
		number("2"),
		tok(token.SUB),
		number("3"),
		tok(token.SUB),
		number("5"),
		tok(token.EOF),
	}

	expectedTokenListTest(expr, expected[:], t)
}

func TestMixOperator(t *testing.T) {
	expr := "2/2+3*5-6"
	expected := [...]TokenItem{
		number("2"),
		tok(token.QUO),
		number("2"),
		tok(token.ADD),
		number("3"),
		tok(token.MUL),
		number("5"),
		tok(token.SUB),
		number("6"),
		tok(token.EOF),
	}

	expectedTokenListTest(expr, expected[:], t)
}

func TestMixOperator2(t *testing.T) {
	expr := "2 / 2 + 3 * 5 - 6"
	expected := [...]TokenItem{
		number("2"),
		tok(token.QUO),
		number("2"),
		tok(token.ADD),
		number("3"),
		tok(token.MUL),
		number("5"),
		tok(token.SUB),
		number("6"),
		tok(token.EOF),
	}

	expectedTokenListTest(expr, expected[:], t)
}

func TestFloatNumbers(t *testing.T) {
	expr := "2.56 / 2.65e3 + 3.6e-6 * 5 - 6"
	expected := [...]TokenItem{
		number("2.56"),
		tok(token.QUO),
		number("2.65e3"),
		tok(token.ADD),
		number("3.6e-6"),
		tok(token.MUL),
		number("5"),
		tok(token.SUB),
		number("6"),
		tok(token.EOF),
	}

	expectedTokenListTest(expr, expected[:], t)
}
