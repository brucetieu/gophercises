package cmd

import (
	"fmt"

	"github.com/brucetieu/gophercises/ex07/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(completedCmd)
}

var completedCmd = &cobra.Command{
	Use: "completed",
	Short: "List all of your completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing out completed tasks:")
		tasks, err := db.GetTasks()

		if err != nil {
			e := fmt.Errorf("error getting all tasks: %s", err)
			fmt.Println(e)
		} else {
			for _, task := range tasks {
				if task.Completed {
					fmt.Printf("- %s\n", task.Task)
				}
			}
		}
	},
}

