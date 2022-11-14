package typedstrings

import "testing"

func TestBackspaceCompare(t *testing.T) {
	s := "ab#c"
	t1 := "#abb##c"
	expected := true
	response := BackspaceCompare(s, t1)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}

func TestBackspaceCompareNoMatch(t *testing.T) {
	s := ""
	t1 := "###"
	expected := true
	response := BackspaceCompare(s, t1)
	if expected != response {
		t.Errorf("expected %v, got %v", expected, response)
	}
}
