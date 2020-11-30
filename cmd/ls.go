package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "ls [query...]",
	Short:   "List tasks",
	Example: "task ls\ntask ls +myproject\ntask ls myquery",
	Aliases: []string{"list"},
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := txt.ListTasks(args)
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
	RootCmd.AddCommand(lsCmd)
}
