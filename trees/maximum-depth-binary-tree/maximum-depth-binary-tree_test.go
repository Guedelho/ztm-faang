package maximumdepthbinarytree

import "testing"

func TestMaxDepth(t *testing.T) {
	tree := &TreeNode{Val: 0}
	tree.Left = &TreeNode{Val: 9}
	tree.Right = &TreeNode{Val: 20}
	expected := 2
	response := MaxDepth(tree)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
