package cmd

import (
	"fmt"
	"os"

	"github.com/kkga/task/txt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := txt.AllTasks()
		if err != nil {
			fmt.Println("Failed to get tasks", err)
			os.Exit(1)
		}

		for i, task := range tasks {
			fmt.Println(fmt.Sprintf("%2d | %s", i+1, task))
		}
		fmt.Println("-------------------------")
		fmt.Println("Total tasks: ", len(tasks))
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
