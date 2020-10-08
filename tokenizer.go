package lab2

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

type TokenType int

const (
	Literal  = 0
	Operator = 1
	Unknown  = 2
	Eof      = 3
)

type Token struct {
	Type TokenType
	Literal   []byte
}

type Tokenizer struct {
	Source []byte
}


func isAlphNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func isOperator(c rune) bool {
	return c == '+' || c  == '-' || c == '*' || c == '/' || c == '^'
}

func (t *Tokenizer) next() Token {
	t.Source = bytes.Trim(t.Source, " \t\n")
	if len(t.Source) == 0 {
		return Token{Eof, []byte("<eof>")}
	}

	var first, size = utf8.DecodeRuneInString(string(t.Source))
	if unicode.IsDigit(first) {
		return t.slice(unicode.IsDigit)
	} else if unicode.IsLetter(first) {
		return t.slice(isAlphNum)
	} else if isOperator(first) {
		var operator = t.Source[0:1]
		t.Source = t.Source[1:]
		return Token{ Operator, operator }
	}

	var unknownSlice = t.Source[0:size]
	t.Source = t.Source[size:]
	return Token{ Unknown, unknownSlice }
}

func (t *Tokenizer) slice(fn func(rune) bool) Token {
	var without = bytes.TrimLeftFunc(t.Source, fn)
	var length = len(t.Source) - len(without)
	var res = t.Source[0:length]
	t.Source = without
	return Token{Literal, res}
}

func (t *Token) Debug() string {
	var tokenType string
	switch t.Type {
	case Literal:
		tokenType = "Literal"
	case Operator:
		tokenType = "Operator"
	case Eof:
		tokenType = "Eof"
	default:
		tokenType = "Unknown"
	}

	value := string(t.Literal)
	return "{ " + tokenType + ", \"" + value + "\" }"
}

func Tokenize(source string) []Token {
	var res = make([]Token, 0)
	var lexer =  Tokenizer{ []byte(source) }
	for tk := lexer.next(); tk.Type != Eof; tk = lexer.next() {
		res = append(res, tk)
	}
	return res
}
