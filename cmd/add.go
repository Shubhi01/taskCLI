package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskDescr := strings.Join(args, " ")
		task := db.Task{Description: taskDescr}
		err := db.AddTask(&task)
		if err != nil {
			fmt.Println("Failed to add task")
		}
		fmt.Printf("Added task \"%v\"", task)
		// db.AddTask(task)
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
