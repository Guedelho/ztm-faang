package cycledetection

import "testing"

func TestHasCycleTrue(t *testing.T) {
	l := &ListNode{Val: 0}
	l.Next = &ListNode{Val: 1}
	l.Next.Next = &ListNode{Val: 2}
	l.Next.Next.Next = l.Next
	expected := true
	response := HasCycle(l)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestHasCycleFalse(t *testing.T) {
	l := &ListNode{Val: 0}
	l.Next = &ListNode{Val: 1}
	l.Next.Next = &ListNode{Val: 2}
	expected := false
	response := HasCycle(l)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestDetectCycle(t *testing.T) {
	l := &ListNode{Val: 0}
	l.Next = &ListNode{Val: 1}
	l.Next.Next = &ListNode{Val: 2}
	l.Next.Next.Next = l.Next
	expected := 1
	response := DetectCycle(l).Val
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestDetectCycleNoMatch(t *testing.T) {
	l := &ListNode{Val: 0}
	l.Next = &ListNode{Val: 1}
	l.Next.Next = &ListNode{Val: 2}
	response := DetectCycle(l)
	if nil != response {
		t.Errorf("expected %v, got %v", nil, response)
	}
}
