package cmd

import (
	"fmt"
	"os"
	"strings"

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
			statusStr := "[ ]"
			if strings.HasPrefix(task, "x ") {
				statusStr = "[x]"
				task = strings.Replace(task, "x ", "", 1)
			}
			fmt.Println(fmt.Sprintf("%2d %s %s", i+1, statusStr, task))
		}
		fmt.Println("-------------------------")
		fmt.Println("Total tasks: ", len(tasks))
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
