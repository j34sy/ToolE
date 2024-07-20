/*
Copyright © 2024 j34sy j34sy@proton.me
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// notesEditCmd represents the notes edit command
var notesEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a note (not fully implemented yet)",
	Long: `As of now, the edit command takes a note name or ID and the new content. With "" your note will be empty.
	Be aware, the content will be replaced with the new content.
	With ToolE notes you can see all existing notes.
	
	Examples:
	ToolE notes edit testnote "This is an edited test note"
	ToolE notes edit 0 "This is an edited test note"
	ToolE notes edit emptyNote ""
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("notesEdit called")
	},
}

func init() {
	notesCmd.AddCommand(notesEditCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notesEditCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notesEditCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
