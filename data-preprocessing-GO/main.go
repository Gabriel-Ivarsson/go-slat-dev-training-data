// Package main provides main
package main

import (
	"fmt"
	"os"

	"github.com/Gabriel-Ivarsson/code2vec-demo/data-preprocessing-GO/src/ASTparser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a directory")
		return
	}
	dir := os.Args[1]
	var modelChoice string
	if len(os.Args) < 3 {
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
