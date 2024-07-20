/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// notesShowCmd represents the notes show command
var notesShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show one note",
	Long: `Shows the content of a note with the given name or the note's ID in the list.
	With ToolE notes you can see all existing notes.
	
	Examples:
	ToolE notes show testnote
	ToolE notes show 0
	`,
	Run: func(cmd *cobra.Command, args []string) {
		checkNotesDir()
		fmt.Println("notesShow called")
	},
}

func init() {
	notesCmd.AddCommand(notesShowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notesShowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notesShowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
