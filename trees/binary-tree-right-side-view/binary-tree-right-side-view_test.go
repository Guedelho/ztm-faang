package binarytreerightsideview

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tree := &TreeNode{Val: 0}
	tree.Left = &TreeNode{Val: 9}
	tree.Right = &TreeNode{Val: 20}
	expected := []int{0, 20}
	response := RightSideView(tree)
	if !reflect.DeepEqual(expected, response) {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
