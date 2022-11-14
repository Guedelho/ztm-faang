package longestsubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abccabb"
	expected := 3
	response := LengthOfLongestSubstring(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestLengthOfLongestSubstringNoMatch(t *testing.T) {
	s := ""
	expected := 0
	response := LengthOfLongestSubstring(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestLengthOfLongestSubstringCatch1(t *testing.T) {
	s := "bbbbb"
	expected := 1
	response := LengthOfLongestSubstring(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestLengthOfLongestSubstringCatch2(t *testing.T) {
	s := "dvdf"
	expected := 3
	response := LengthOfLongestSubstring(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
