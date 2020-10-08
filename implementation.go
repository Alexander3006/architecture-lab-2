package lab2

import (
	"errors"
)

type Node interface {
	toPrefix() string
}

type LiteralNode struct {
	value []byte
}

func (literalNode LiteralNode) toPrefix() string {
	prefix := string(literalNode.value)
	return prefix
}

type ExpressionNode struct {
	lhs      Node
	rhs      Node
	operator []byte
}

func (expressionNode ExpressionNode) toPrefix() string {
	var prefix = string(expressionNode.operator) + " " + expressionNode.rhs.toPrefix() + " " + expressionNode.lhs.toPrefix()
	return prefix
}

type Stack []Token

func (stack *Stack) pop() (Token, error) {
	if len(*stack) == 0 {
		return Token{}, errors.New("Stack is empty")
	} else {
		index := len(*stack) - 1
		elem := (*stack)[index]
		*stack = (*stack)[:index]
		return elem, nil
	}
}

func toAST(stack *Stack) Node {
	top, _ := stack.pop()
	if top.Type == Literal {
		return LiteralNode{top.Literal}
	}
	lhs := toAST(stack)
	rhs := toAST(stack)
	return ExpressionNode{lhs, rhs, top.Literal}
}

func PostfixToPrefix(input string) (string, error) {
	var tokens = Stack(Tokenize(input))
	for _, tk := range tokens {
		if tk.Type != Literal && tk.Type != Operator {
			return "", errors.New("Unknown token: " + string(tk.Literal))
		}
	}
	var ast = toAST(&tokens)
	return ast.toPrefix(), nil
}
