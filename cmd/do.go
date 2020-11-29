package cmd

import (
	"fmt"
	"strconv"

	"github.com/kkga/task/txt"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks tasks as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse: ", arg)
			} else {
				ids = append(ids, id)
			}
		}

		for _, id := range ids {
			err := txt.CompleteTask(id)
			// err := txt.DeleteTask(id)
			if err != nil {
				fmt.Println("Failed to delete", err)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
