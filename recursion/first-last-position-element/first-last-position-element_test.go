package firstlastpositionelement

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := []int{3, 4}
	response := SearchRange(nums, target)
	if !reflect.DeepEqual(expected, response) {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
func TestSearchRangeNoMatch(t *testing.T) {
	nums := []int{1, 3, 3, 5}
	target := 6
	expected := []int{-1, -1}
	response := SearchRange(nums, target)
	if !reflect.DeepEqual(expected, response) {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
