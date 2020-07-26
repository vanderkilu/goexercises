package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"../db"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task by id",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(strings.Join(args, " "))
		if err != nil {
			fmt.Println("kindly pass in an integer(id) for task to be deleted")
		} else {
			err := db.DeleteTask(id)
			if err != nil {
				fmt.Println("There was an error deleting the task")
			} else {
				fmt.Println("task of id %d deleted successfully", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
