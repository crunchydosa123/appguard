package parser

import (
	sitter "github.com/smacker/go-tree-sitter"
	js "github.com/smacker/go-tree-sitter/javascript"
)

func Parse(code []byte) (*sitter.Tree, error) {
	parser := sitter.NewParser()
	parser.SetLanguage(js.GetLanguage())

	tree := parser.Parse(nil, code)
	return tree, nil
}
