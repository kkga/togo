package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// archiveCmd represents the archive command
var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Move all completed tasks from todo.txt to done.txt",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("archive called")
	},
}

func init() {
	RootCmd.AddCommand(archiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// archiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// archiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}