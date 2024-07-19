/*
Copyright © 2024 j34sy j34sy@proton.me
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "This command marks a ToDo as done",
	Long: `This command marks a ToDo as done by title or ID.
	You must provide a title or an ID to mark the ToDo as done.
	
	Examples:
	todos done "Title of the ToDo"
	todos done ShortToDo
	todos done 1
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if !viper.IsSet("ToDoData") {
			viper.Set("ToDoData", "todos.yaml")
			fullPath := filepath.Join(viper.GetString("toolePath"), "todos.yaml")
			file, err := os.Create(fullPath)
			if err != nil {
				fmt.Printf("Failed to create data file for ToDos: %s", err)
				os.Exit(1)
			}
			defer file.Close()
			fmt.Println("ToDo Data file created at: ", fullPath)
			viper.WriteConfig()

		}

		todos, err := fetchTodos()
		if err != nil {
			fmt.Printf("Failed to fetch ToDos: %s", err)
			os.Exit(1)
		}

		if len(args) < 1 {
			fmt.Printf("Please provide a title or an ID to mark the ToDo as done")
			os.Exit(1)
		}
		if len(args) > 1 {
			fmt.Printf("Please provide only a title or an ID to mark the ToDo as done")
			os.Exit(1)
		}
		if _, err := strconv.Atoi(args[0]); err == nil {
			id, _ := strconv.Atoi(args[0])
			if !checkForExistenceByID(todos, id) {
				fmt.Println("ToDo with ID", id, "does not exist")
				return
			}
			err := markDonebyID(&todos, id)
			if err != nil {
				fmt.Printf("Failed to mark ToDo as done: %s", err)
				os.Exit(1)
			}
		} else {
			if !checkForExistence(todos, args[0]) {
				fmt.Println("ToDo with title", args[0], "does not exist")
				return
			}
			err := markDonebyTitle(&todos, args[0])
			if err != nil {
				fmt.Printf("Failed to mark ToDo as done: %s", err)
				os.Exit(1)
			}
		}
		err = saveTodos(todos)
		if err != nil {
			fmt.Printf("Failed to save ToDos: %s", err)
			os.Exit(1)
		}
		fmt.Println("ToDo marked as done successfully")
	},
}

func init() {
	todosCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
