/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"learnProject/First-Project-With-Go/auth"
	"learnProject/First-Project-With-Go/server"
	"learnProject/First-Project-With-Go/utils"
)

var SetPort string
var UserName string
var Password string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The Server is ready at your Service...\nport no:", SetPort, "\n")
		auth.SetUserInfo(UserName, Password)
		server.Start(SetPort)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVarP(&SetPort, "port", "d", utils.DefaultPortNo, "set port no")
	startCmd.Flags().StringVarP(&UserName, "user name", "u", "h", "set the user name")
	startCmd.Flags().StringVarP(&Password, "password", "p", "hi", "set password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
