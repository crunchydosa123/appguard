package scanner

import (
	"appguard/internal/parser"
	"appguard/internal/rules"
	"os"
	"path/filepath"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func WalkWithFile(node *sitter.Node, source []byte, file string, findings *[]rules.Finding) {
	rules.CheckTriggers(node, source, file, findings)

	for i := 0; i < int(node.ChildCount()); i++ {
		WalkWithFile(node.Child(i), source, file, findings)
	}
}

func ScanRepo(root string) ([]rules.Finding, error) {
	var allFindings []rules.Finding

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			if shouldSkipDir(info.Name()) {
				return filepath.SkipDir
			}
			return nil
		}

		if isJSFile(path) {
			findings, err := scanFile(path)
			if err == nil {
				allFindings = append(allFindings, findings...)
			}
		}

		return nil
	})

	return allFindings, err
}

func isJSFile(path string) bool {
	return strings.HasSuffix(path, ".js") ||
		strings.HasSuffix(path, ".ts")
}

func shouldSkipDir(name string) bool {
	skip := map[string]bool{
		"node_modules": true,
		".git":         true,
		"dist":         true,
		"build":        true,
		".next":        true,
	}

	return skip[name]
}

func scanFile(path string) ([]rules.Finding, error) {

	source, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	tree, err := parser.Parse(source)
	if err != nil {
		return nil, err
	}

	var findings []rules.Finding

	WalkWithFile(tree.RootNode(), source, path, &findings)

	return findings, nil
}
