package maths

import (
	"testing"
)

func TestHexCalc(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "Addition",
			input: []string{"a", "+", "b"},
			want:  "15", // a+b = 10+11 = 21 in decimal = 15 in hex
		},
		{
			name:  "Subtraction",
			input: []string{"f", "-", "a"},
			want:  "5", // f-a = 15-10 = 5 in decimal = 5 in hex
		},
		{
			name:  "Multiplication",
			input: []string{"a", "*", "2"},
			want:  "14", // a*2 = 10*2 = 20 in decimal = 14 in hex
		},
		{
			name:  "Division",
			input: []string{"f", "/", "3"},
			want:  "5", // f/3 = 15/3 = 5 in decimal = 5 in hex
		},
		{
			name:  "Modulus",
			input: []string{"f", "%", "a"},
			want:  "5", // f%a = 15%10 = 5 in decimal = 5 in hex
		},
		{
			name:  "Exponentiation",
			input: []string{"2", "^", "3"},
			want:  "8", // 2^3 = 2^3 = 8 in decimal = 8 in hex
		},
		{
			name:  "Nth Root",
			input: []string{"100", "|", "2"},
			want:  "10", // sqrt(100) = sqrt(256) = 16 in decimal = 10 in hex
		},
		{
			name:  "Nested Parentheses",
			input: []string{"(", "a", "+", "(", "2", "*", "3", ")", ")", "*", "2"},
			want:  "20", // (a+(2*3)) = (10+(2*3)) = (10+6) = 16 in decimal = 10 in hex
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexCalc(tt.input); got != tt.want {
				t.Errorf("HexCalc() = %v, want %v", got, tt.want)
				t.Logf("Input: %v", tt.input)
			}
		})
	}
}
