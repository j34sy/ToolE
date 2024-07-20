/*
Copyright © 2024 j34sy j34sy@proton.me
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// notesAddCmd represents the notes add command
var notesAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Creates a new note",
	Long: `Creates a new note with the given name and optional content.
	The name needs to be unique as it is the file name of the note.
	With ToolE notes you can see all existing notes.
	
	Examples:
	ToolE notes add testnote "This is a test note"
	ToolE notes add emptynote
	`,
	Run: func(cmd *cobra.Command, args []string) {
		checkNotesDir()
		fmt.Println("notesAdd called")
	},
}

func init() {
	notesCmd.AddCommand(notesAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notesAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notesAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
