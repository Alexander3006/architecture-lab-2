package lab2

import (
	"errors"
	"fmt"
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
	op := string(expressionNode.operator)
	rhs := expressionNode.rhs.toPrefix()
	lhs := expressionNode.lhs.toPrefix()
	return fmt.Sprintf("%s %s %s", op, rhs, lhs)
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

func toOperator(stack *Stack, top Token) (Node, error) {
	lhs, lerr := toAST(stack)
	if lerr != nil {
		return nil, lerr
	}
	rhs, rerr := toAST(stack)
	if rerr != nil {
		return nil, rerr
	}
	return ExpressionNode{lhs, rhs, top.Literal}, nil
}

func toAST(stack *Stack) (Node, error) {
	top, err := stack.pop()
	if err != nil {
		return nil, errors.New("Unexpected end")
	}
	if top.Type == Literal {
		return LiteralNode{top.Literal}, nil
	} else if top.Type == Operator {
		return toOperator(stack, top)
	} else {
		return nil, fmt.Errorf("Unexpected token: %s", string(top.Literal))
	}
}

func PostfixToPrefix(input string) (string, error) {
	var tokens = Stack(Tokenize(input))
	var ast, err = toAST(&tokens)
	if err == nil {
		if len(tokens) != 0 {
			top, _ := tokens.pop()
			err = fmt.Errorf("Unexpected token: %s", string(top.Literal))
			return "", err
		}
		return ast.toPrefix(), nil		
	} else {
		return "", err
	}
}
