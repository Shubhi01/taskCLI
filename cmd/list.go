package cmd

import (
	"fmt"
	"os"
	"task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all todo tasks",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list called")
		tasks, err := db.ListTask()
		if err != nil {
			fmt.Println("Something went wrong")
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks left! Go for a vacation!")
		} else {
			for i, task := range tasks {
				fmt.Printf("%d: %s", task.Id+1, task.Description)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
