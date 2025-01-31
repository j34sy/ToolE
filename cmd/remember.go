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

// rememberCmd represents the remember command
var rememberCmd = &cobra.Command{
	Use:   "remember",
	Short: "This command can remember something for you",
	Long: `Remember can save something either temporary, using the OS temporary folder, or persistently into a file within ToolE's directory

	Remember works on a key - value base or adding order (temporary first); if something to remeber has a key, you can "search" for it, otherwise the top X entries will be shown, while temporary has a higher priority.
	KEYS NEED TO BE UNIQUE

	To save something persistently, use the -p flag while adding something
	During any kind of show/ list command, the kind of saving will be listed but is not a criteria to show something

	For all these functionalities multiple Subcommands are added.
	- add: Additional option to add something
	- remove: Remove an entry (by key or ID of list)
	- list: More options for listing entries
	- search: advanced search options for keys
	- top: show the top X entries

	This command has multiple functions combied on remember to reduce overhead and make it easier to use.
	Some of these take flags:
	- -a or --all: will show all entries
	- -v or --verbose: show persistency of the entry
	- -t or --top: will show the top X entries, default is 5 (hopefully)
	- -s or --search: will search for a key and display the entry
	- -p or --persist: will save the entry persistently
	- -k or --key: will save the entry with a key
	- -h or --help: will show the help message for the command
	
	Usage:
	remember 
	remember -a -v
	remember -t 10
	remember -s key
	remember some string
	remember -p some string
	remember -k key some string
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if !viper.IsSet("PersistentRememberData") {
			viper.Set("PersistentRememberData", "remember.yaml")
			fullPath := filepath.Join(viper.GetString("toolePath"), "remember.yaml")
			file, err := os.Create(fullPath)
			if err != nil {
				log.Fatalf("Failed to create data file for Remember: %s", err)
			}
			defer file.Close()
			fmt.Println("Remember Data file created at: ", fullPath)
			viper.WriteConfig()
		}

		if _, err := os.Stat(filepath.Join(os.TempDir(), "ToolE")); os.IsNotExist(err) {
			os.Mkdir(filepath.Join(os.TempDir(), "ToolE"), os.ModePerm)
		}
		if _, err := os.Stat(filepath.Join(os.TempDir(), "ToolE", "remember.yaml")); os.IsNotExist(err) {
			file, err := os.Create(filepath.Join(os.TempDir(), "ToolE", "remember.yaml"))
			if err != nil {
				log.Fatalf("Failed to create data file for Remember: %s", err)
			}
			defer file.Close()
			fmt.Println("Remember Data temporary file created at: ", filepath.Join(os.TempDir(), "ToolE", "remember.yaml"))
		}

		all, _ := cmd.Flags().GetBool("all")
		verbose, _ := cmd.Flags().GetBool("verbose")
		top, _ := cmd.Flags().GetInt("top")
		search, _ := cmd.Flags().GetString("search")
		persist, _ := cmd.Flags().GetBool("persist")
		key, _ := cmd.Flags().GetString("key")

		rem, err := fetchRemember()
		if err != nil {
			log.Fatalf("Failed to fetch Remember: %s", err)
		}

		if len(args) == 0 {
			if len(rem) == 0 {
				fmt.Println("I remember nothing!")
			} else {
				if search != "" {
					for _, r := range rem {
						if r.Key == search {
							fmt.Printf("Key: %s\nValue: %s\nPersist: %t\n", r.Key, r.Value, r.Persist)
						}
					}
				} else {
					if all {
						for i, r := range rem {
							if r.Key != "" {
								if verbose {
									fmt.Printf("ID: %d\nKey: %s\nValue: %s\nPersist: %t\n", i, r.Key, r.Value, r.Persist)
								} else {
									fmt.Printf("ID: %d\nKey: %s\nValue: %s\n", i, r.Key, r.Value)
								}
							} else {
								if verbose {
									fmt.Printf("ID: %d\nValue: %s\nPersist: %t\n", i, r.Value, r.Persist)
								} else {
									fmt.Printf("ID: %d\nValue: %s\n", i, r.Value)
								}
							}
						}
					} else {
						if top > len(rem) {
							top = len(rem)
						}
						for i := 0; i < top; i++ {
							if rem[i].Key != "" {
								if verbose {
									fmt.Printf("ID: %d\nKey: %s\nValue: %s\nPersist: %t\n", i, rem[i].Key, rem[i].Value, rem[i].Persist)
								} else {
									fmt.Printf("ID: %d\nKey: %s\nValue: %s\n", i, rem[i].Key, rem[i].Value)
								}
							} else {
								if verbose {
									fmt.Printf("ID: %d\nValue: %s\nPersist: %t\n", i, rem[i].Value, rem[i].Persist)
								} else {
									fmt.Printf("ID: %d\nValue: %s\n", i, rem[i].Value)
								}
							}
						}
					}
				}
			}
		} else {
			if persist {
				if key != "" {
					for _, r := range rem {
						if r.Key == key {
							fmt.Println("Key already exists, please use a different key")
							return
						}
					}
					newRem := Remember{Key: key, Value: args[0], Persist: true}
					err = saveRemember(newRem)
				} else {
					newRem := Remember{Key: "", Value: args[0], Persist: true}
					err = saveRemember(newRem)
				}
				if err != nil {
					log.Fatalf("Failed to save Remember: %s", err)
				}
			} else {
				if key != "" {
					for _, r := range rem {
						if r.Key == key {
							fmt.Println("Key already exists, please use a different key")
							return
						}
					}
					newRem := Remember{Key: key, Value: args[0], Persist: false}
					err = saveRemember(newRem)
				} else {
					newRem := Remember{Key: "", Value: args[0], Persist: false}
					err = saveRemember(newRem)
				}
				if err != nil {
					log.Fatalf("Failed to save Remember: %s", err)
				}
			}
		}

	},
}

type Remember struct {
	Key     string
	Value   string
	Persist bool
}

func fetchRemember() ([]Remember, error) {
	var data []Remember
	tempPath := filepath.Join(os.TempDir(), "ToolE", "remember.yaml")
	persistPath := filepath.Join(viper.GetString("toolePath"), "remember.yaml")

	tempData, err := os.ReadFile(tempPath)
	if err != nil {
		return nil, err
	}
	persistData, err := os.ReadFile(persistPath)
	if err != nil {
		return nil, err
	}

	var tempRemembers []Remember
	err = yaml.Unmarshal(tempData, &tempRemembers)
	if err != nil {
		return nil, err
	}

	var persistRemembers []Remember
	err = yaml.Unmarshal(persistData, &persistRemembers)
	if err != nil {
		return nil, err
	}
	for i := len(tempRemembers) - 1; i >= 0; i-- {
		data = append(data, tempRemembers[i])
	}
	data = append(data, persistRemembers...)

	return data, nil
}

func saveRemember(data Remember) error {
	tempPath := filepath.Join(os.TempDir(), "ToolE", "remember.yaml")
	persistPath := filepath.Join(viper.GetString("toolePath"), "remember.yaml")

	var remembers []Remember

	if data.Persist {
		fileData, err := os.ReadFile(persistPath)
		if err == nil {
			err = yaml.Unmarshal(fileData, &remembers)
			if err != nil {
				return err
			}
		}
		remembers = append(remembers, data)
		newData, err := yaml.Marshal(&remembers)
		if err != nil {
			return err
		}
		err = os.WriteFile(persistPath, newData, 0644)
		if err != nil {
			return err
		}
	} else {
		fileData, err := os.ReadFile(tempPath)
		if err == nil {
			err = yaml.Unmarshal(fileData, &remembers)
			if err != nil {
				return err
			}
		}
		remembers = append(remembers, data)
		newData, err := yaml.Marshal(&remembers)
		if err != nil {
			return err
		}
		err = os.WriteFile(tempPath, newData, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(rememberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rememberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	rememberCmd.Flags().BoolP("all", "a", false, "Show all entries")
	rememberCmd.Flags().BoolP("verbose", "v", false, "Show persistency of the entry")
	rememberCmd.Flags().IntP("top", "t", 5, "Show the top X entries")
	rememberCmd.Flags().StringP("search", "s", "", "Search for a key")
	rememberCmd.Flags().BoolP("persist", "p", false, "Save the entry persistently")
	rememberCmd.Flags().StringP("key", "k", "", "Save the entry with a key")
}
