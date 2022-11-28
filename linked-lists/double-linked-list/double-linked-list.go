package doublelinkedlist

import "fmt"

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func (l *Node) Print() (list string) {
	for l != nil {
		list += fmt.Sprintf("->%v", l.Val)
		l = l.Next
	}
	list = list[2:]
	return
}

func (l *Node) Insert(val int) {
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &Node{Val: val, Prev: l}
}

func Flatten(root *Node) *Node {
	dummy := root
	for dummy != nil {
		if dummy.Child == nil {
			dummy = dummy.Next
			continue
		}
		curr := dummy.Child
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = dummy.Next
		if dummy.Next != nil {
			dummy.Next.Prev = curr
		}
		dummy.Next = dummy.Child
		dummy.Child.Prev = dummy
		dummy.Child = nil
	}
	return root
}
