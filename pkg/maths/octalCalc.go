package maths

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func OctalCalc(input []string) string {
	// actual calculation
	if !checkInput(input) {
		os.Exit(1)
	}
	result := recursiveCalcOct(input)
	return result
}

func recursiveCalcOct(input []string) string {
	// catch ends
	if len(input) == 1 {
		_, err := strconv.ParseInt(input[0], 8, 64)
		if err != nil {
			fmt.Println("Invalid input: ", err)
			os.Exit(1)
		}
		return input[0]
	}

	// Remove outer parentheses if they exist
	if input[0] == "(" && input[len(input)-1] == ")" {
		paras := 0
		octOps := 0
		for _, v := range input {
			if v == "(" {
				paras++
			}
			for _, op := range BinaryOperators {
				if v == op {
					octOps++
				}
			}
		}
		if paras == octOps {
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
		if _, err := strconv.ParseInt(input[0], 8, 64); err != nil {
			fmt.Println("Invalid input: no operator found")
			os.Exit(1)
		}
		return recursiveCalcOct([]string{input[0]})
	}

	// Calculate the result of the operation
	switch input[lowestIndex] {
	case "+":
		val := addOct(recursiveCalcOct(input[:lowestIndex]), recursiveCalcOct(input[lowestIndex+1:]))
		return val
	case "-":
		val := subtractOct(recursiveCalcOct(input[:lowestIndex]), recursiveCalcOct(input[lowestIndex+1:]))
		return val
	case "*":
		val := multiplyOct(recursiveCalcOct(input[:lowestIndex]), recursiveCalcOct(input[lowestIndex+1:]))
		return val
	case "/":
		val := divideOct(recursiveCalcOct(input[:lowestIndex]), recursiveCalcOct(input[lowestIndex+1:]))
		return val
	case "%":
		val := modulusOct(recursiveCalcOct(input[:lowestIndex]), recursiveCalcOct(input[lowestIndex+1:]))
		return val
	case "^":
		val := exponentiationOct(recursiveCalcOct(input[:lowestIndex]), recursiveCalcOct(input[lowestIndex+1:]))
		return val
	case "|":
		val := nthRootOct(recursiveCalcOct(input[:lowestIndex]), recursiveCalcOct(input[lowestIndex+1:]))
		return val
	}

	// This line should never be reached
	fmt.Println("Invalid input: unknown error")
	os.Exit(1)
	return ""
}

func addOct(a, b string) string {
	ai, _ := strconv.ParseInt(a, 8, 64)
	bi, _ := strconv.ParseInt(b, 8, 64)
	return strconv.FormatInt(ai+bi, 8)
}

func subtractOct(a, b string) string {
	ai, _ := strconv.ParseInt(a, 8, 64)
	bi, _ := strconv.ParseInt(b, 8, 64)
	return strconv.FormatInt(ai-bi, 8)
}

func multiplyOct(a, b string) string {
	ai, _ := strconv.ParseInt(a, 8, 64)
	bi, _ := strconv.ParseInt(b, 8, 64)
	return strconv.FormatInt(ai*bi, 8)
}

func divideOct(a, b string) string {
	ai, _ := strconv.ParseInt(a, 8, 64)
	bi, _ := strconv.ParseInt(b, 8, 64)
	return strconv.FormatInt(ai/bi, 8)
}

func modulusOct(a, b string) string {
	ai, _ := strconv.ParseInt(a, 8, 64)
	bi, _ := strconv.ParseInt(b, 8, 64)
	return strconv.FormatInt(ai%bi, 8)
}

func exponentiationOct(a, b string) string {
	ai, _ := strconv.ParseInt(a, 8, 64)
	bi, _ := strconv.ParseInt(b, 8, 64)
	return strconv.FormatInt(int64(math.Pow(float64(ai), float64(bi))), 8)
}

func nthRootOct(a, b string) string {
	ai, _ := strconv.ParseInt(a, 8, 64)
	bi, _ := strconv.ParseInt(b, 8, 64)
	result := math.Pow(float64(ai), 1/float64(bi))
	roundedResult := int64(math.Round(result)) // round to nearest whole number
	return strconv.FormatInt(roundedResult, 8)
}
