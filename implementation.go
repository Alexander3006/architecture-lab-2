package lab2

import (
	"errors"
	"fmt"
)

type Node interface {
	toPrefix() string
}

type NumberNode struct {
	sum byte
}

func (numberNode NumberNode) toPrefix() string {
	prefix := string(numberNode.sum)
	return prefix
}

type ExpressionNode struct {
	lhs      Node
	rhs      Node
	operator byte
}

func (expressionNode ExpressionNode) toPrefix() string {
	var prefix string = string(expressionNode.operator) + " " + expressionNode.rhs.toPrefix() + " " + expressionNode.lhs.toPrefix()
	return prefix
}

type Stack []byte

func (stack *Stack) pop() (byte, error) {
	if len(*stack) == 0 {
		return 0, errors.New("Stack is empty")
	} else {
		index := len(*stack) - 1
		elem := (*stack)[index]
		*stack = (*stack)[:index]
		return elem, nil
	}

}

func ToAST(stack *Stack) Node {
	top, _ := stack.pop()
	if '0' <= top && top <= '9' {
		return NumberNode{top}
	}
	lhs := ToAST(stack)
	rhs := ToAST(stack)
	return ExpressionNode{lhs, rhs, top}
}

//TODO: add token split function
//Example of use
// tokenStack := Stack{'4', '2', '-', '3', '*', '5', '+'}
// ast := ToAST(&tokenStack)
// fmt.Println(ast.toPrefix())

// TODO: document this function.
// PrefixToPostfix converts
func PrefixToPostfix(input string) (string, error) {
	return "TODO", fmt.Errorf("TODO")
}
