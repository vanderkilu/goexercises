package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as done by specifying id or lists of ids",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to convert %d to string", arg)
			} else {
				fmt.Println("Task %d is marked as done", id)
			}
			ids = append(ids, id)
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
