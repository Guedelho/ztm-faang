package containerwithmostwater

import "testing"

func TestMaxArea(t *testing.T) {
	height := []int{7, 1, 2, 3, 9}
	expected := 28
	response := MaxArea(height)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestMaxAreaNoMatch(t *testing.T) {
	height := []int{1}
	expected := 0
	response := MaxArea(height)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
