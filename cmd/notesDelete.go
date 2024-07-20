/*
Copyright © 2024 j34sy j34sy@proton.me
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// notesDeleteCmd represents the notes delete command
var notesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a note",
	Long: `Deletes a note with the given name or the note's ID in the list.
	With ToolE notes you can see all existing notes.
	
	Examples:
	ToolE notes delete testnote
	ToolE notes delete 0
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("notesDelete called")
	},
}

func init() {
	notesCmd.AddCommand(notesDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notesDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notesDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
