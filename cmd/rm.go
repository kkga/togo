package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "Remove a task",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Cannot parse task number")
		}

		deletedTask, err := txt.DeleteTask(key)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Removed:", deletedTask)
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
