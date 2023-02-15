// Package astParser provides astParser
package astParser

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	word2vec = 1
	fast     = 2
)

var (
	funcType    = "funcDecl"
	packageType = "package"
)

// GetModel function
// takes string `fast` or `w2v`
func GetModel(modelName string) (int, error) {
	if modelName == "fast" {
		return fast, nil
	} else if modelName == "w2v" {
		return word2vec, nil
	}

	return 0, errors.New("Error; No valid model was given, valid: {\"w2v\", \"fast\"}")
}

func getTypeWord(model int, typeName string) string {
	if model == word2vec {
		return typeName
	}
	return ""
}

func sanitizeName(name string) string {
	name = strings.ToLower(name)
	name = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(name, "")
	return name
}

func generateContext(model int, node *ast.File) []string {
	var context []string
	var parentPackage *ast.File

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.File:
			packageName := sanitizeName(x.Name.Name)
			if strings.HasPrefix(packageName, "test") {
				// early return, we dont want to add test stuff to context.
				return false
			}
			parentPackage = x
		case *ast.FuncDecl:
			funcName := sanitizeName(x.Name.Name)
			packageName := sanitizeName(parentPackage.Name.Name)
			if strings.HasPrefix(funcName, "test") {
				// early return, we dont want to add test stuff to context.
				return false
			}
			context = append(
				context,
				fmt.Sprintf("%s%s %s %s",
					getTypeWord(model, "funcDecl"),
					getTypeWord(model, "package"),
					funcName,
					packageName,
				),
			)
		}
		return true
	})
	return context
}

func context2json(contextList []string) []string {
	var json []string
	for i := 0; i < len(contextList); i++ {
		jsonLine := "{\"context\":" + "\"" + contextList[i] + "\"}"
		json = append(json, jsonLine)
	}
	return json
}

func getGoFileNode(path string) (*ast.File, error) {
	if filepath.Ext(path) == ".go" {
		fset := token.NewFileSet()

		node, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return nil, err
		}
		return node, nil
	}
	return nil, nil
}

func write2DataFile(json []string, fileName string) error {
	f, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	for i := 0; i < len(json); i++ {
		if _, err := f.WriteString(string(json[i]) + "\n"); err != nil {
			return err
		}
	}

	f.Close()
	return nil
}

// ParseDir function
func ParseDir(model int, dir string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		node, err := getGoFileNode(path)
		if node != nil {
			contextList := generateContext(model, node)
			json := context2json(contextList)
			write2DataFile(json, "data.json")
		} else {
			if err != nil {
				log.Println(err)
			}
		}
		return nil
	})
	if err != nil {
	}
	return nil
}
