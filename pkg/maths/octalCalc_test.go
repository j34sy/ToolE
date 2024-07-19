package maths

import (
	"testing"
)

func TestOctalCalc(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected string
	}{
		{"Addition", []string{"10", "+", "11"}, "21"},
		{"Subtraction", []string{"21", "-", "11"}, "10"},
		{"Multiplication", []string{"10", "*", "11"}, "110"},
		{"Division", []string{"110", "/", "11"}, "10"},
		{"Modulus", []string{"111", "%", "11"}, "1"},
		{"Exponentiation", []string{"10", "^", "11"}, "1000000000"},
		{"Nth Root", []string{"33", "|", "3"}, "3"},
		{"Complex expression", []string{"(", "10", "+", "11", ")", "*", "100"}, "2100"},
		{"Extreme complex expression", []string{"(", "(", "13", "+", "11", ")", "*", "2", ")", "+", "(", "2", "*", "4", ")"}, "60"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := OctalCalc(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
				t.Logf("Input: %v", tc.input)
			}
		})
	}
}
