package reverselinkedlist

import "testing"

func TestReverseInterative(t *testing.T) {
	l := ListNode{Val: 0, Next: nil}
	l.Next = &ListNode{Val: 1, Next: nil}
	l.Next.Next = &ListNode{Val: 2, Next: nil}
	l.Next.Next.Next = &ListNode{Val: 3, Next: nil}
	expected := "3->2->1->0"
	response := ReverseListInterative(&l).Print()
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestReverseRecursive(t *testing.T) {
	l := ListNode{Val: 0, Next: nil}
	l.Next = &ListNode{Val: 1, Next: nil}
	l.Next.Next = &ListNode{Val: 2, Next: nil}
	l.Next.Next.Next = &ListNode{Val: 3, Next: nil}
	expected := "3->2->1->0"
	response := ReverseListRecursive(&l).Print()
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestPrint(t *testing.T) {
	l := ListNode{Val: 0, Next: nil}
	l.Next = &ListNode{Val: 1, Next: nil}
	l.Next.Next = &ListNode{Val: 2, Next: nil}
	l.Next.Next.Next = &ListNode{Val: 3, Next: nil}
	expected := "0->1->2->3"
	response := l.Print()
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
