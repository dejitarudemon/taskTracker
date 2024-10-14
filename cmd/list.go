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
		var status_func *string
		status, _ := cmd.Flags().GetString("status")

		if status == "" {
			status_func = nil
		} else {
			status_func = &status
		}

		tasks, err := crud.List(status_func)
		if err != nil {
			fmt.Println(err)
		} else {
			for i, task := range tasks {
				fmt.Printf("%d | ID: %d | DESC: %v | STATUS: %v | CREATED: %v | UPDATED: %v\n", i, task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().String("status", "", "The task status")
}
