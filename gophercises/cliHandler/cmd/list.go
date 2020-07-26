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
			fmt.Printf("ID\tValue\n\n")
			for _, task := range tasks {
				fmt.Printf("%d\t%s\n", task.Id, task.Name)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
