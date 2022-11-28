package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	s := "()[]{}"
	expected := true
	response := IsValid(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestIsValidNoMatch(t *testing.T) {
	s := "("
	expected := false
	response := IsValid(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
