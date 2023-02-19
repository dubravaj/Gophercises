package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dubravaj/task/cmd"
	"github.com/dubravaj/task/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	homePath, _ := homedir.Dir()
	dbPath := filepath.Join(homePath, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
