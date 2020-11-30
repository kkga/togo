package main

import (
	"os"

	"github.com/kkga/togo/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
