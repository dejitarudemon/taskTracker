/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/dejitarudemon/taskTracker/crud"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := crud.List(nil)
		if err != nil {
			fmt.Println(err)
		} else {
			for i, task := range tasks {
				fmt.Printf("%d | ID: %d | DESC: %v | CREATED: %v | UPDATED: %v\n", i, task.Id, task.Description, task.CreatedAt, task.UpdatedAt)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
