package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added task \"%s\"", task)
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
