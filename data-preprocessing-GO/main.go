// Package main provides main
package main

import (
	"fmt"
	"os"

	"github.com/Gabriel-Ivarsson/code2vec-demo/data-preprocessing-GO/src/ASTparser"
)

func main() {
	dir := os.Args[1]
	if os.Args[1] == "" {
		fmt.Println("Please provide a directory")
		return
	}
	var modelChoice string
	if os.Args[2] == "" {
		modelChoice = "fast"
	} else {
		modelChoice = os.Args[2]
	}
	model, err := astParser.GetModel(modelChoice)
	if err != nil {
		fmt.Print(err)
		return
	}
	astParser.ParseDir(model, dir)
	return
}
