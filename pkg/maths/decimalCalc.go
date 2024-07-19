package maths

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var BinaryOperators = []string{"+", "-", "*", "/", "%", "^", "|"} // binary operators

func checkInput(input []string) bool {
	// check if the input is valid (number of paranthesis and operators match)
	countParanthesisopen := 0
	countParanthesisclose := 0
	countBinaryOperators := 0

	for _, v := range input {
		if v == "(" {
			countParanthesisopen++
		} else if v == ")" {
			countParanthesisclose++
		}
		for _, op := range BinaryOperators {
			if v == op {
				countBinaryOperators++
			}
		}
	}

	if countParanthesisopen != countParanthesisclose {
		return false
	}
	if countBinaryOperators == countParanthesisopen || countBinaryOperators == countParanthesisclose+1 {
		return true
	}
	return false
}

func DecimalCalc(input []string) float64 {
	// actual calculation
	if !checkInput(input) {
		os.Exit(1)
	}
	result := recursiveCalc2(input)
	return result
}

func recursiveCalc2(input []string) float64 {
	// catch ends
	if len(input) == 1 {
		val, err := strconv.ParseFloat(input[0], 64)
		if err != nil {
			fmt.Println("Invalid input: ", err)
			os.Exit(1)
		}
		return val
	}

	// Remove outer parentheses if they exist

	if input[0] == "(" && input[len(input)-1] == ")" {
		paras := 0
		binOps := 0
		for _, v := range input {
			if v == "(" {
				paras++
			}
			for _, op := range BinaryOperators {
				if v == op {
					binOps++
				}
			}
		}
		if paras == binOps {
			input = input[1 : len(input)-1]
		}
	}

	// Find the operator with the lowest precedence
	lowestPrecedence := len(BinaryOperators)
	lowestIndex := -1
	depth := 0
	for i, v := range input {
		if v == "(" {
			depth++
		} else if v == ")" {
			depth--
		} else {
			for j, op := range BinaryOperators {
				if v == op && depth == 0 && j <= lowestPrecedence {
					lowestPrecedence = j
					lowestIndex = i
				}
			}
		}
	}

	// If no operator was found, the input is invalid
	if lowestIndex == -1 {
		// Check if the input is a single number
		if _, err := strconv.ParseFloat(input[0], 64); err != nil {
			fmt.Println("Invalid input: no operator found")
			os.Exit(1)
		}
		return recursiveCalc2([]string{input[0]})
	}

	// Calculate the result of the operation
	switch input[lowestIndex] {
	case "+":
		val := add(recursiveCalc2(input[:lowestIndex]), recursiveCalc2(input[lowestIndex+1:]))
		return val
	case "-":
		val := subtract(recursiveCalc2(input[:lowestIndex]), recursiveCalc2(input[lowestIndex+1:]))
		return val
	case "*":
		val := multiply(recursiveCalc2(input[:lowestIndex]), recursiveCalc2(input[lowestIndex+1:]))
		return val
	case "/":
		val := divide(recursiveCalc2(input[:lowestIndex]), recursiveCalc2(input[lowestIndex+1:]))
		return val
	case "%":
		val := modulus(recursiveCalc2(input[:lowestIndex]), recursiveCalc2(input[lowestIndex+1:]))
		return val
	case "^":
		val := exponentiation(recursiveCalc2(input[:lowestIndex]), recursiveCalc2(input[lowestIndex+1:]))
		return val
	case "|":
		val := nthRoot(recursiveCalc2(input[:lowestIndex]), recursiveCalc2(input[lowestIndex+1:]))
		return val
	}

	// This line should never be reached
	fmt.Println("Invalid input: unknown error")
	os.Exit(1)
	return 0
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}

func modulus(a, b float64) float64 {
	return a - (math.Floor(a/b) * b)
}

func exponentiation(a, b float64) float64 {
	return math.Pow(a, b)
}

func nthRoot(a, b float64) float64 {
	return math.Pow(a, 1/b)
}
