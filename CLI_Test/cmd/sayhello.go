/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sayhelloCmd represents the sayhello command
var sayhelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		greeting := "Hello"
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "World"
		}
		if viper.GetString("name") != "" {
			name = viper.GetString("name")
		}
		if viper.GetString("greeting") != "" {
			greeting = viper.GetString("greeting")
		}
		fmt.Println(greeting + " " + name)
	},
}

func init() {
	sayCmd.AddCommand(sayhelloCmd)
	viper.AutomaticEnv()
	sayhelloCmd.Flags().StringP("name", "n", viper.GetString("ENVNAME"), "Set your Name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sayhelloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sayhelloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
