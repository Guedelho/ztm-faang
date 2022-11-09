package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	target := 9
	expected := []int{3, 4}
	response := TwoSum(nums, target)
	if !reflect.DeepEqual(expected, response) {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
func TestTwoSumNoMatch(t *testing.T) {
	nums := []int{1, 2, 3}
	target := 10
	response := TwoSum(nums, target)
	if response != nil {
		t.Errorf("expected %v, got %v", nil, response)
	}
}
