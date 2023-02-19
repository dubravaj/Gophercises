package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dubravaj/task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task to task list",
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		_, err := db.CreateTask(taskName)
		if err != nil {
			fmt.Printf("Failed to add task '%s', error: %s\n", taskName, err.Error())
			os.Exit(1)
		} else {
			fmt.Printf("Task '%s' added to your tasks list. \n", taskName)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
