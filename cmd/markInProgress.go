/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/dejitarudemon/taskTracker/crud"
	"github.com/spf13/cobra"
)

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "markInProgress",
	Short: "Mark the task as in-progress",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
		}

		if err := crud.Mark(taskId, crud.InProgress); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("OK")
		}
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markInProgressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markInProgressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
