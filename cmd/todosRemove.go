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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "This command removes a ToDo",
	Long: `This command removes a ToDo from the list of ToDos by title or ID.
	You must provide a title or an ID to remove the ToDo.
	
	Examples:
	todos remove "Title of the ToDo"
	todos remove ShortToDo
	todos remove 1
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
			fmt.Printf("Please provide a title or an ID to remove the ToDo")
			os.Exit(1)
		}
		if len(args) > 1 {
			fmt.Printf("Please provide only a title or an ID to remove the ToDo")
			os.Exit(1)
		}

		if _, err := strconv.Atoi(args[0]); err == nil {
			id, _ := strconv.Atoi(args[0])
			if !checkForExistenceByID(todos, id) {
				fmt.Println("ToDo with ID", id, "does not exist")
				return
			}
			err := removeTodoByID(&todos, id)
			if err != nil {
				fmt.Printf("Failed to remove ToDo: %s", err)
				os.Exit(1)
			}
		} else {
			if !checkForExistence(todos, args[0]) {
				fmt.Println("ToDo with title", args[0], "does not exist")
				return
			}

			err := removeTodoByTitle(&todos, args[0])
			if err != nil {
				fmt.Printf("Failed to remove ToDo: %s", err)
				os.Exit(1)
			}
		}
		err = saveTodos(todos)
		if err != nil {
			fmt.Printf("Failed to save ToDos: %s", err)
			os.Exit(1)
		}
		fmt.Println("ToDo removed successfully")
	},
}

func init() {
	todosCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
