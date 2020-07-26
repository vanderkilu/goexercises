package main

import (
	"fmt"
	"path/filepath"

	"./cmd"
	"./db"
	baseDir "github.com/mitchellh/go-homedir"
)

func main() {
	homeDir, _ := baseDir.Dir()
	dbPath := filepath.Join(homeDir, "tasks.db")
	err := db.InitDB(dbPath)
	if err != nil {
		fmt.Println("There was error initializing the db")
	}
	cmd.RootCmd.Execute()
}
