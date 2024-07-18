/*
Copyright © 2024 j34sy j34sy@proton.me
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Todo struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Done        bool   `yaml:"done"`
}

// todosCmd represents the todos command
var todosCmd = &cobra.Command{
	Use:   "todos",
	Short: "A command to manage ToDos - lightweight task manager",
	Long: `This application is a tool to manage ToDos. It allows you to add, remove, mark as done, and list ToDos.
	One ToDo has an unique title, an optional description, and a status (done or not done).
	ToDos are stored in a YAML file in the user's home directory (or otherwise specified directory) within the .ToolE application directory.
	It comes with multiple subcommands to manage ToDos:
	- add: Add a new ToDo
	- remove: Remove a ToDo by title or ID
	- done: Mark a ToDo as done by title or ID
	
	This command lists the ToDos available in the data file.
	The flags:
	- -v or --verbose: will print the description of the ToDos
	- -a or --all: will show all ToDos including the ones marked as done
	- -h or --help: will show the help message for the command

	The numbers in front of the ToDos are the IDs of the ToDos.
	Attention: The IDs are zero-based, so the first ToDo has the ID 0 and they are dynamic, so they change when ToDos are removed.
	
	Examples:
	toodos
	todos -v
	todos -a -v
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if !viper.IsSet("ToDoData") {
			viper.Set("ToDoData", "todos.yaml")
			fullPath := filepath.Join(viper.GetString("toolePath"), "todos.yaml")
			file, err := os.Create(fullPath)
			if err != nil {
				log.Fatalf("Failed to create data file for ToDos: %s", err)
			}
			defer file.Close()
			fmt.Println("ToDo Data file created at: ", fullPath)
			viper.WriteConfig()

		}
		verbose, _ := cmd.Flags().GetBool("verbose")
		all, _ := cmd.Flags().GetBool("all")

		todos, err := fetchTodos()
		if err != nil {
			log.Fatalf("Failed to fetch ToDos: %s", err)
		}

		if len(args) == 0 {

			if len(todos) == 0 {
				fmt.Println("You have no ToDos, all caught up!")
			} else {
				if verbose {
					for i, todo := range todos {
						if all || !todo.Done {
							fmt.Printf("%d: %s \n", i, todo.Title)
							fmt.Printf("    Description: %s \n", todo.Description)
							fmt.Printf("    Done: %s \n", fmt.Sprintf("%t", todo.Done))
						}
					}
				} else {
					for i, todo := range todos {
						if all || !todo.Done {
							fmt.Printf("%d. %s \n", i, todo.Title)
						}
					}
				}

			}
		} else {
			fmt.Println("Please provide a valid command, this command does not take any arguments")
		}
	},
}

func fetchTodos() ([]Todo, error) {
	var data []Todo

	todosRaw, err := os.ReadFile(filepath.Join(viper.GetString("toolePath"), viper.GetString("ToDoData")))
	if err != nil {
		configMap := viper.AllSettings()
		delete(configMap, "tododata")
		newViper := viper.New()
		newViper.MergeConfigMap(configMap)
		newViper.WriteConfigAs(viper.ConfigFileUsed())
		fmt.Println("ToDoData file not existing where it should be, deleting the path. A new data-file will be created on the next run.")
		return nil, err
	}

	err = yaml.Unmarshal(todosRaw, &data)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func saveTodos(todos []Todo) error {

	todosRaw, err := yaml.Marshal(todos)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(viper.GetString("toolePath"), viper.GetString("ToDoData")), todosRaw, 0644)
	if err != nil {
		return err
	}

	return nil
}

func addTodo(todos *[]Todo, newTodo Todo) error {
	*todos = append(*todos, newTodo)
	return nil
}

func removeTodoByTitle(todos *[]Todo, title string) error {
	for i, todo := range *todos {
		if todo.Title == title {
			*todos = append((*todos)[:i], (*todos)[i+1:]...)
			break
		}
	}
	return nil
}

func markDonebyTitle(todos *[]Todo, title string) error {
	for i, todo := range *todos {
		if todo.Title == title {
			(*todos)[i].Done = true
			break
		}
	}
	return nil
}

func markDonebyID(todos *[]Todo, id int) error {
	(*todos)[id].Done = true
	return nil
}

func removeTodoByID(todos *[]Todo, id int) error {
	*todos = append((*todos)[:id], (*todos)[id+1:]...)
	return nil
}

func checkForDuplicate(todos []Todo, title string) bool {
	for _, todo := range todos {
		if todo.Title == title {
			return true
		}
	}
	return false
}

func checkForExistence(todos []Todo, title string) bool {
	for _, todo := range todos {
		if todo.Title == title {
			return true
		}
	}
	return false
}

func checkForExistenceByID(todos []Todo, id int) bool {
	return id < len(todos)
}

func init() {
	rootCmd.AddCommand(todosCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	todosCmd.Flags().BoolP("verbose", "v", false, "This flag will print descriptions of the ToDos (if available)")
	todosCmd.Flags().BoolP("all", "a", false, "Show all flags (marked as done too)")
}
