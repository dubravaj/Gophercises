package cmd

import (
	"fmt"
	"os"

	"github.com/dubravaj/task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Failed to list the tasks.", err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have currently no tasks.")
		} else {
			fmt.Println("Listing current tasks: ")
			for i, task := range tasks {
				fmt.Printf("%d : '%s' \n", i+1, task.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
