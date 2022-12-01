package kthlargestelement

import "testing"

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	expected := 4
	response := FindKthLargest(nums, k)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestPartition(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	left := 0
	right := len(nums) - 1
	expected := 8
	response := Partition(nums, left, right)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
