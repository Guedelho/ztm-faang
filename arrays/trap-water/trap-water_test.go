package trapwater

import "testing"

func TestTrapWater(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	expected := 6
	response := TrapWater(height)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestTrapWaterNoMatch(t *testing.T) {
	height := []int{1}
	expected := 0
	response := TrapWater(height)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
