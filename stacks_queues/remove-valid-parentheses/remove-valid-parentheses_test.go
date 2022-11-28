package removevalidparentheses

import "testing"

func TestMinRemoveToMakeValid(t *testing.T) {
	s := "ma(te)(u)s)"
	expected := "ma(te)(u)s"
	response := MinRemoveToMakeValid(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestMinRemoveToMakeValidNoMatch(t *testing.T) {
	s := "))(("
	expected := ""
	response := MinRemoveToMakeValid(s)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
