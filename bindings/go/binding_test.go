package tree_sitter_mcfunction_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_mcfunction "github.com/futarimiti/tree-sitter-mcfunction-lang/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_mcfunction.Language())
	if language == nil {
		t.Errorf("Error loading MC Function grammar")
	}
}
