package rules

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

type Finding struct {
	File   string
	Line   int
	Column int
	Code   string
	Type   string
}

func CheckTriggers(node *sitter.Node, source []byte, file string, findings *[]Finding) {
	text := string(source[node.StartByte():node.EndByte()])
	line, col := getLineAndColumn(source, node.StartByte())

	if node.Type() == "call_expression" {
		if contains(text, "query(") && contains(text, "+") {
			*findings = append(*findings, Finding{
				Line:   line,
				Column: col,
				Code:   text,
				Type:   "possible_sql_injection",
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

func getLineAndColumn(source []byte, byteOffset uint32) (int, int) {
	line := 1
	col := 1

	for i := 0; i < int(byteOffset) && i < len(source); i++ {
		if source[i] == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}
	return line, col
}
