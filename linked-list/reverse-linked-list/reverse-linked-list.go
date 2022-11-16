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
	var rev, dummy *ListNode = nil, head
	for dummy != nil {
		next := dummy.Next
		dummy.Next = rev
		rev = dummy
		dummy = next
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
