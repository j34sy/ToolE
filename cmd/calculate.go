/*
Copyright © 2024 j34sy j34sy@proton.me
*/
package cmd

import (
	"fmt"
	"os"

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
	- Factorial (!)

	Each operation is performed on two numbers, so you have to use paranthesis to group operations.
	Example: calculate 2+3*4 is undefined, but calculate (2+3)*4 is 20 and calculate 2+(3*4) is 14.
	You can also use the flags to specify the base of the numbers.
	Available flags are:
	- --decimal (-d)
	- --hexadecimal (-h)
	- --octal (-o)
	- --binary (-b)

	Example: calculate --hexadecimal (a^3) + (8|3) is 3EA
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("calculate called")

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

		if count > 1 {
			fmt.Println("Only one flag is allowed")
			os.Exit(1)
		}

		if hexadecimal {
			fmt.Println("Hexadecimal")
		} else if octal {
			fmt.Println("Octal")
		} else if binary {
			fmt.Println("Binary")
		} else {
			fmt.Println("No flag or decimal flag")
			fmt.Println("Default: Decimal")
		}
	},
}

func init() {
	rootCmd.AddCommand(calculateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calculateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	calculateCmd.Flags().BoolP("decimal", "d", false, "Understand the numbers as decimal")
	calculateCmd.Flags().BoolP("hexadecimal", "h", false, "Understand the numbers as hexadecimal")
	calculateCmd.Flags().BoolP("octal", "o", false, "Understand the numbers as octal")
	calculateCmd.Flags().BoolP("binary", "b", false, "Understand the numbers as binary")
}
