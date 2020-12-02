package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:     "do [NUM]",
	Aliases: []string{"d, done"},
	Short:   "Mark todo as done",
	Run: func(cmd *cobra.Command, args []string) {
		var keys []int
		for _, arg := range args {
			key, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse: ", arg)
			} else {
				keys = append(keys, key)
			}
		}

		for _, k := range keys {
			todo, err := txt.CompleteTodo(k)
			if err != nil {
				fmt.Println("Can't complete todo #", err)
				os.Exit(1)
			}

			fmt.Println(fmt.Sprintf("%s", todo.Subject))
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
