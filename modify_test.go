package piscine

import (
	"fmt"
	"testing"
)

func TestFunctions(t *testing.T) {
	tests := []struct {
        input    string
        expected string
    }{
		{"1E (hex) files were added", "30 files were added"},
        {"It has been 101 (bin) years", "It has been 5 years"},
        {"Ready, set, go (up,2) !", "Ready, SET, GO!"},
        {"I should stop SHOUTING (low)", "I should stop shouting"},
	}
	for _, test := range tests {
        result := ModText(test.input)
        if result != test.expected {
            t.Errorf("ModText(%q) = %q; expected %q", test.input, result, test.expected)
        } else {
            fmt.Println(result)
        }
    }
}