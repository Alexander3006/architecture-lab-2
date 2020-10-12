package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression to compute")
	outputFile      = flag.String("o", "", "File to output the result")
)

func main() {
	flag.Parse()
	var reader io.Reader
	var writer io.Writer
	//Reader
	if len(*inputExpression) > 0 {
		reader = strings.NewReader(*inputExpression)
	} else if len(*inputFile) > 0 {
		file, err := os.Open(*inputFile)
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		reader = file
	} else {
		fmt.Println("Reader interface not defined")
		os.Exit(1)
	}
	//Writer
	if len(*outputFile) > 0 {
		file, err := os.Create(*outputFile)
		defer file.Close()
		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		writer = file
	} else {
		writer = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
