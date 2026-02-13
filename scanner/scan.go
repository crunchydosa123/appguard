package scanner

import (
	"appguard/rules"

	siiter "github.com/smacker/go-tree-sitter"
)

func Walk(node *siiter.Node, source []byte, findings *[]rules.Finding) {
	rules.CheckTriggers(node, source, findings)

	for i := 0; i < int(node.ChildCount()); i++ {
		Walk(node.Child(i), source, findings)
	}
}
