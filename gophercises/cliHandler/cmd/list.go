package cmd

import (
	"fmt"

	"../db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Couldn't retrieve the tasks")
		} else {
			for _, task := range tasks {
				fmt.Println("%d. %s", task.Id, task.Name)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
