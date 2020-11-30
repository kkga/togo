package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:     "do",
	Aliases: []string{"d, done"},
	Short:   "Mark tasks as complete",
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
			task, err := txt.CompleteTask(id)
			// err := txt.DeleteTask(id)
			if err != nil {
				fmt.Println("Failed to delete", err)
			}
			statusStr := "[ ]"
			if strings.HasPrefix(task, "x ") {
				statusStr = "[x]"
				task = strings.Replace(task, "x ", "", 1)
			}
			fmt.Println(fmt.Sprintf("%s %s", statusStr, task))
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
