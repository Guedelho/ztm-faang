package doublelinkedlist

import "testing"

func TestFlatten(t *testing.T) {
	root := &Node{Val: 1}
	root.Insert(2)
	root.Insert(3)
	root.Insert(4)
	root.Insert(5)
	root.Insert(6)
	child := &Node{Val: 7}
	child.Insert(8)
	child.Insert(9)
	child.Insert(10)
	child2 := &Node{Val: 11}
	child2.Insert(12)
	child.Next.Child = child2
	root.Next.Next.Child = child
	expected := "1->2->3->7->8->11->12->9->10->4->5->6"
	response := Flatten(root).Print()
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
