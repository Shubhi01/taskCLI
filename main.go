package main

import (
	"fmt"
	"os"
	"task/cmd"
	"task/db"

	"github.com/boltdb/bolt"
)

func main() {
	opts := bolt.Options{}
	err := db.InitDB("task.db", &opts)
	if err != nil {
		fmt.Println("Failed to open db")
		os.Exit(1)
	}
	defer db.CloseDB()
	cmd.RootCmd.Execute()
}
