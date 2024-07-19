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

var toolePath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ToolE",
	Short: "A collection of useful tools",
	Long: `A list of tools that can be used to make your life easier:
	- subnet: A Subnet Calculator for IPv4 
	- TBD`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		PrintToolE()
		version, _ := cmd.Flags().GetBool("version")
		if version {
			fmt.Println("ToolE v0.0.1")
		} else {
			fmt.Println("Use 'ToolE --help' for more information")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&toolePath, "config", "", "ToolE directory path (default is $HOME/.ToolE)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("version", "v", false, "Print the version number of ToolE")
}

func PrintToolE() {
	// Print ToolE logo (figlet)
	fmt.Println("  :::::::::::       ::::::::       ::::::::       :::        :::::::::: ")
	fmt.Println("     :+:          :+:    :+:     :+:    :+:      :+:        :+:         ")
	fmt.Println("    +:+          +:+    +:+     +:+    +:+      +:+        +:+          ")
	fmt.Println("   +#+          +#+    +:+     +#+    +:+      +#+        +#++:++#      ")
	fmt.Println("  +#+          +#+    +#+     +#+    +#+      +#+        +#+            ")
	fmt.Println(" #+#          #+#    #+#     #+#    #+#      #+#        #+#             ")
	fmt.Println("###           ########       ########       ########## ##########       ")
}

func initConfig() {
	viper.AutomaticEnv()

	if toolePath != "" {
		// Use path from flag
		dirPath, err := filepath.Abs(toolePath)
		if err != nil {
			fmt.Println("Error getting absolute path: ", err)
			os.Exit(1)
		}

		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			err := os.Mkdir(dirPath, 0755)
			if err != nil {
				fmt.Println("Error creating ToolE directory: ", err)
				os.Exit(1)
			}
		}

		viper.AddConfigPath(dirPath)
		viper.SetConfigName(".ToolE")
		viper.Set("toolePath", dirPath)
	} else {
		// Use default path
		homeDir, _ := os.UserHomeDir()
		path := filepath.Join(homeDir, ".ToolE")

		dirPath := filepath.Join(homeDir, ".ToolE")
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			err := os.Mkdir(dirPath, 0755)
			if err != nil {
				fmt.Println("Error creating ToolE directory: ", err)
				os.Exit(1)
			}
		}

		viper.AddConfigPath(path)
		viper.SetConfigName(".ToolE")
		viper.Set("toolePath", path)
	}

	if err := viper.ReadInConfig(); err == nil {
		// Config file was found and loaded
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		// Config file was not found, create it
		fmt.Println("Config file not found, creating one at: ", viper.GetString("toolePath"))
		err := viper.SafeWriteConfigAs(filepath.Join(viper.GetString("toolePath"), ".ToolE.yaml"))
		if err != nil {
			fmt.Println("Failed to write config file during init: ", err)
			os.Exit(1)
		}
	}

}
