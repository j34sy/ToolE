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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
