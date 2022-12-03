package queueusingstacks

import "testing"

func TestMyQueue(t *testing.T) {
	myQueue := Constructor()
	if myQueue.Empty() != true {
		t.Errorf("expected %v, got %v", true, false)
	}
	myQueue.Push(1)
	if myQueue.Empty() != false {
		t.Errorf("expected %v, got %v", false, true)
	}
	peek := myQueue.Peek()
	if peek != 1 {
		t.Errorf("expected %v, got %v", 1, peek)
	}
	pop := myQueue.Pop()
	if pop != 1 {
		t.Errorf("expected %v, got %v", 1, pop)
	}
	if myQueue.Empty() != true {
		t.Errorf("expected %v, got %v", true, false)
	}
}
