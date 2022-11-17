package reverselinkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() (list string) {
	for l != nil {
		list += fmt.Sprintf("->%v", l.Val)
		l = l.Next
	}
	list = list[2:]
	return
}

func ReverseListInterative(head *ListNode) *ListNode {
	var rev *ListNode
	for head != nil {
		next := head.Next
		head.Next = rev
		rev = head
		head = next
	}
	return rev
}

func ReverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rev := ReverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rev
}

func ReverseBetween(head *ListNode, left, right int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	before := dummy
	current := head
	for i := 0; i < left-1; i++ {
		before = current
		current = current.Next
	}
	var rev *ListNode
	for i := left; i <= right; i++ {
		next := current.Next
		current.Next = rev
		rev = current
		current = next
	}
	before.Next.Next = current
	before.Next = rev
	return dummy.Next
}
