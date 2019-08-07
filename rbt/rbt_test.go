package rbt

import (
	"testing"
)

func TestTreeSize(t *testing.T) {
	tree := NewTree()

	nodeCount := 1000
	for i := 0; i < nodeCount; i++ {
		tree.Put(i, i)
	}

	if tree.FastSize() != nodeCount {
		t.Errorf("expected %d, got %d", nodeCount, tree.FastSize())
	}
}
