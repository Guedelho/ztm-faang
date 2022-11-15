package palindrome

import "testing"

func TestIsPalindromeTwoPointers(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	expected := true
	response := IsPalindromeTwoPointers(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestIsPalindromeReverse(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	expected := true
	response := IsPalindromeReverse(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestFormatString(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	expected := "amanaplanacanalpanama"
	response := formatString(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestReverseString(t *testing.T) {
	s := "amanaplanacanalpanama"
	expected := "amanaplanacanalpanama"
	response := reverseString(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestAlmostPalindrome(t *testing.T) {
	s := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"
	expected := true
	response := AlmostPalindrome(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestIsPalindrome(t *testing.T) {
	s := "aaa"
	expected := true
	response := AlmostPalindrome(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
