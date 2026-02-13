package rules

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

type Finding struct {
	Code string
	Type string
}

func CheckTriggers(node *sitter.Node, source []byte, findings *[]Finding) {
	text := string(source[node.StartByte():node.EndByte()])

	if node.Type() == "call_expression" {
		if contains(text, "query(") && contains(text, "+") {
			*findings = append(*findings, Finding{
				Code: text,
				Type: "possible_sql_injection",
			})
		}
	}

	if contains(text, "md5(") || contains(text, "sha1(") {
		*findings = append(*findings, Finding{
			Code: text,
			Type: "weak_crypto",
		})
	}
}

func contains(s, sub string) bool {
	return strings.Contains(s, sub)
}
