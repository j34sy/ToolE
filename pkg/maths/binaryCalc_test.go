package maths

import (
	"testing"
)

func TestBinaryCalc(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected string
	}{
		{"Addition", []string{"10", "+", "11"}, "101"},
		{"Subtraction", []string{"101", "-", "11"}, "10"},
		{"Multiplication", []string{"10", "*", "11"}, "110"},
		{"Division", []string{"110", "/", "11"}, "10"},
		{"Modulus", []string{"111", "%", "11"}, "1"},
		{"Exponentiation", []string{"10", "^", "11"}, "1000"},
		{"Nth Root", []string{"11011", "|", "11"}, "11"},
		{"Complex expression", []string{"(", "10", "+", "11", ")", "*", "100"}, "10100"},
		{"Extreme complex expression", []string{"(", "(", "10", "+", "11", ")", "*", "100", ")", "+", "(", "101", "*", "110", ")"}, "110010"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BinaryCalc(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
				t.Logf("Input: %v", tc.input)
			}
		})
	}
}
