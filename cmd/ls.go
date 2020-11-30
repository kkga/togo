package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
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

		// iteration over map happens in random order, so we store the order
		// in a separate slice
		var keys []int
		for k := range tasks {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, k := range keys {
			if strings.HasPrefix(tasks[k], "x ") {
				crossedOut := color.New(color.CrossedOut).SprintFunc()
				taskStr := strings.Replace(tasks[k], "x ", "", 1)
				fmt.Println(fmt.Sprintf("%2d %s %s", k, "[x]", crossedOut(taskStr)))
			} else {
				fmt.Println(fmt.Sprintf("%2d %s %s", k, "[ ]", tasks[k]))
			}
		}

		totalLen, _ := txt.GetTotalTodoLen("todo.txt")
		fmt.Println("------")
		fmt.Printf("%d/%d todos shown\n", len(tasks), totalLen)
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
	// lsCmd.Flags().BoolP("done", "d", false, "List done tasks from done.txt")
}
