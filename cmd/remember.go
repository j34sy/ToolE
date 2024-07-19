/*
Copyright © 2024 j34sy j34sy@proton.me
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rememberCmd represents the remember command
var rememberCmd = &cobra.Command{
	Use:   "remember",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remember called")
	},
}

func init() {
	rootCmd.AddCommand(rememberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rememberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rememberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
