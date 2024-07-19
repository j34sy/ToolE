/*
Copyright © 2024 j34sy j34sy@proton.me
*/
package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/j34sy/ToolE/pkg/maths"
	"github.com/spf13/cobra"
)

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculates a given term",
	Long: `This commands calculates a given term.
	You can specify the base of the numbers using the flags.
	By default, the numbers are considered decimal.
	Supported operations are:
	- Addition (+)
	- Subtraction (-)
	- Multiplication (*)
	- Division (/)
	- Modulus (%)
	- Exponentiation (^)
	- nth Root (|)

	Each operation is performed on two numbers, so you have to use paranthesis to group operations.
	Example: calculate 2+3*4 is undefined, but calculate (2+3)*4 is 20 and calculate 2+(3*4) is 14.
	You can also use the flags to specify the base of the numbers.
	Available flags are:
	- --decimal (-D) [default]
	- --hexadecimal (-H)  [case-insensitive]
	- --octal (-O)
	- --binary (-B)

	Example: calculate --hexadecimal (a^3) + (8|3) is 3EA (1002 in decimal)
	Please Note:
	- The numbers are considered decimal by default.
	- Only one flag is allowed.
	- Decimal values can be floats except for factorial calculations (only integers).
	- The nth root can be expressed with exponentiation and factorial as well.
	- Especially the non-decimal calculations are not as accurate as the decimal calculations due to rounding errors.
	- Do not use spaces between characters

	Examples:
	- calculate 2+3
	- calculate 2+(3*4)
	- calculate --hexadecimal (a^3)+(8|3)
	`,
	Run: func(cmd *cobra.Command, args []string) {
		decimal, _ := cmd.Flags().GetBool("decimal")
		hexadecimal, _ := cmd.Flags().GetBool("hexadecimal")
		octal, _ := cmd.Flags().GetBool("octal")
		binary, _ := cmd.Flags().GetBool("binary")

		flags := []bool{decimal, hexadecimal, octal, binary}
		count := 0
		for _, flag := range flags {
			if flag {
				count++
			}
		}

		input := splitCalculate(args[0])

		if count > 1 {
			fmt.Println("Only one flag is allowed")
			os.Exit(1)
		}

		if len(args) == 0 {
			fmt.Println("No term provided")
			os.Exit(1)
		}

		if hexadecimal {
			checkHex(input)
			println(maths.HexCalc(input))

		} else if octal {
			checkOct(input)
			println(maths.OctalCalc(input))
		} else if binary {
			checkBin(input)
			println(maths.BinaryCalc(input))
		} else {
			checkDec(input)
			println(maths.DecimalCalc(input))
		}
	},
}

func checkHex(input []string) {
	for _, v := range input {
		for _, c := range v {
			if !checkValList(hexDigits, string(c)) && !checkValList(operatorsChars, string(c)) {
				fmt.Println("Invalid character in hexadecimal")
				os.Exit(1)
			}
		}
	}
}

func checkOct(input []string) {
	for _, v := range input {
		for _, c := range v {
			if !checkValList(octDigits, string(c)) && !checkValList(operatorsChars, string(c)) {
				fmt.Println("Invalid character in octal")
				os.Exit(1)
			}
		}
	}
}

func checkBin(input []string) {
	for _, v := range input {
		for _, c := range v {
			if !checkValList(binDigits, string(c)) && !checkValList(operatorsChars, string(c)) {
				fmt.Println("Invalid character in binary")
				os.Exit(1)
			}
		}
	}
}

func checkDec(input []string) {
	for _, v := range input {
		for _, c := range v {
			if !checkValList(decimalDigits, string(c)) && !checkValList(operatorsChars, string(c)) {
				fmt.Println("Invalid character in decimal")
				os.Exit(1)
			}
		}
	}
}

var hexDigits = []string{"a", "b", "c", "d", "e", "f", "A", "B", "C", "D", "E", "F", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var octDigits = []string{"0", "1", "2", "3", "4", "5", "6", "7"}
var binDigits = []string{"0", "1"}
var decimalDigits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var operatorsChars = []string{"+", "-", "*", "/", "%", "^", "|", "(", ")"}

func splitCalculate(input string) []string {
	re := regexp.MustCompile(`([^+\-*/%^|()]+)`)
	parts := re.FindAllString(input, -1)
	return parts
}

func checkValList(list []string, val string) bool {
	// Add your code here
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(calculateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calculateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	calculateCmd.Flags().BoolP("decimal", "D", false, "Understand the numbers as decimal")
	calculateCmd.Flags().BoolP("hexadecimal", "H", false, "Understand the numbers as hexadecimal")
	calculateCmd.Flags().BoolP("octal", "O", false, "Understand the numbers as octal")
	calculateCmd.Flags().BoolP("binary", "B", false, "Understand the numbers as binary")
}
