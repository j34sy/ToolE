package maths

import (
	"math"
	"testing"
)

func TestDecimalCalc(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected float64
	}{
		{"Addition", []string{"2", "+", "3"}, 5.0},
		{"Subtraction", []string{"5", "-", "3"}, 2.0},
		{"Multiplication", []string{"2", "*", "3"}, 6.0},
		{"Division", []string{"6", "/", "3"}, 2.0},
		{"Modulus", []string{"7", "%", "3"}, 1.0},
		{"Exponentiation", []string{"2", "^", "3"}, 8.0},
		{"Nth Root", []string{"8", "|", "3"}, 2.0},
		{"Complex expression", []string{"(", "2", "+", "3", ")", "*", "4"}, 20.0},
		{"Extreme complex expression", []string{"(", "(", "2", "+", "3", ")", "*", "4", ")", "+", "(", "5", "*", "6", ")"}, 50.0},
	}

	const tolerance = 1e-9

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := DecimalCalc(tc.input)
			diff := math.Abs(result - tc.expected)
			if diff > tolerance {
				t.Errorf("Expected %f, but got %f", tc.expected, result)
			}
		})
	}
}
