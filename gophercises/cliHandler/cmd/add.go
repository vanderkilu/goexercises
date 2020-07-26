package cmd

import (
	"fmt"
	"strings"

	"../db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add particular task to list of tasks",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := db.AddTask(strings.Join(args, " "))
		if err != nil {
			fmt.Println("Coudn't create the task")
		} else {
			fmt.Println("Task created successfully")
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
