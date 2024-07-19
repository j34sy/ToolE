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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command adds a new ToDo",
	Long: `This command adds a new ToDo to the list of ToDos.
	You must provide a title and you can provide an optional description for the ToDo.
	
	Examples:
	todos add "Title of the ToDo"
	todos add "Title of the ToDo" "Description of the ToDo"
	todos add ShortToDo
	todos add ShortToDo ShortDescription
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

		if len(args) > 2 {
			fmt.Println("Please provide a title and an optional description for the ToDo")
			os.Exit(1)
		}
		if len(args) < 1 {
			fmt.Println("Please provide a title for the ToDo, the description is optional")
			os.Exit(1)
		}

		if checkForDuplicate(todos, args[0]) {
			fmt.Printf("ToDo with title %s already exists", args[0])
			os.Exit(1)
		}

		if len(args) == 1 {
			args = append(args, "")
		}

		newTodo := Todo{
			Title:       args[0],
			Description: args[1],
			Done:        false,
		}
		err = addTodo(&todos, newTodo)
		if err != nil {
			fmt.Printf("Failed to add ToDo: %s", err)
			os.Exit(1)
		}
		err = saveTodos(todos)
		if err != nil {
			fmt.Printf("Failed to save ToDos: %s", err)
			os.Exit(1)
		}
		fmt.Println("ToDo added successfully")
	},
}

func init() {
	todosCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
