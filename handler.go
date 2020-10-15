package lab2

import (
	"fmt"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	var expression string
	buffer := make([]byte, 1)
	for {
		n, err := ch.Input.Read(buffer)
		if err == io.EOF {
			break
		}
		expression = fmt.Sprintf("%s%s", expression, string(buffer[:n]))
	}
	result, err := PostfixToPrefix(expression)
	if err != nil {
		return err
	}
	ch.Output.Write([]byte(result))
	ch.Output.Write([]byte("\n"))
	return nil
}
