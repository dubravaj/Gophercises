package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dubravaj/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks the task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse: ", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Printf("Error occured: %s", err.Error())
			os.Exit(1)
		}

		for _, taskId := range ids {
			if taskId <= 0 || taskId > len(tasks) {
				fmt.Println("Invalid id: ", taskId)
				continue
			}
			task := tasks[taskId-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to complete task %d, error: %s", taskId, err.Error())
			} else {
				fmt.Printf("Task %d completed.\n", taskId)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
