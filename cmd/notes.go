/*
Copyright © 2024 j34sy j34sy@proton.me
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "A small note taking application",
	Long: `Since I currently haven't developed an editor YET, any writing needs to be done in the argument itself.
	The notes will be stored in .md files inside the notes directory within the ToolE directory.
	The notes command will list all the notes available in the notes directory.
	
	Available commands:
	- add: Add a new note
	- show: Show one note
	- edit: Edit a note
	- delete: Delete a note

	Examples:
	ToolE notes add testnote "This is a test note"
	ToolE notes show testnote
	ToolE notes edit testnote "This is an edited test note"
	ToolE notes delete testnote

	`,
	Run: func(cmd *cobra.Command, args []string) {
		checkNotesDir()
		fmt.Println("notes called")
	},
}

func init() {
	rootCmd.AddCommand(notesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkNotesDir() {
	// Check if the notes directory exists
	if !viper.IsSet("notesDir") {
		fmt.Println("notesDir not set in config file")
		viper.Set("notesDir", "notes")
		fullPath := filepath.Join(viper.GetString("toolepath"), viper.GetString("notesDir"))
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			err := os.Mkdir(fullPath, 0755)
			if err != nil {
				fmt.Println("Error creating notes directory")
				os.Exit(1)
			}
		}
		fmt.Println("notesDir set to default: notes and created if not existing")
		return
	}
}
