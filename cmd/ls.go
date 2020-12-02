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
	Short:   "List todos",
	Example: "togo ls\ntogo ls +myproject\ntogo ls myquery",
	Aliases: []string{"l, list"},
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := txt.ListTodos(args, "todo.txt")
		if err != nil {
			fmt.Println("Failed to get tasks", err)
			os.Exit(1)
		}

		// iteration over map happens in random order, so we store the order
		// in a separate slice
		var keys []int
		for k := range todos {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, k := range keys {
			if strings.HasPrefix(todos[k], "x ") {
				crossedOut := color.New(color.CrossedOut).SprintFunc()
				fmt.Println(fmt.Sprintf("%2d| %s", k, crossedOut(todos[k])))
			} else {
				fmt.Println(fmt.Sprintf("%2d| %s", k, todos[k]))
			}
		}

		totalLen, _ := txt.GetTotalTodoLen("todo.txt")
		fmt.Println("------")
		fmt.Printf("%d/%d todos shown\n", len(todos), totalLen)
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
	// lsCmd.Flags().BoolP("done", "d", false, "List done tasks from done.txt")
}
