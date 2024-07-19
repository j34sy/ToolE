package maths

import (
	"testing"
)

func TestDecimalCalc(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected string
	}{
		{"Addition", []string{"2", "+", "3"}, "5.00000000"},
		{"Subtraction", []string{"5", "-", "3"}, "2.00000000"},
		{"Multiplication", []string{"2", "*", "3"}, "6.00000000"},
		{"Division", []string{"6", "/", "3"}, "2.00000000"},
		{"Modulus", []string{"7", "%", "3"}, "1.00000000"},
		{"Exponentiation", []string{"2", "^", "3"}, "8.00000000"},
		{"Nth Root", []string{"8", "|", "3"}, "2.00000000"},
		{"Complex expression", []string{"(", "2", "+", "3", ")", "*", "4"}, "20.00000000"},
		{"Extreme complex expression", []string{"(", "(", "2", "+", "3", ")", "*", "4", ")", "+", "(", "5", "*", "6", ")"}, "50.00000000"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := DecimalCalc(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
