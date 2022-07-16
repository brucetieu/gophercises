package cmd

import (
	"fmt"

	"github.com/brucetieu/gophercises/ex07/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listTaskCmd)
}

var listTaskCmd = &cobra.Command{
	Use: "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing out incomplete tasks:")
		tasks, err := db.GetTasks()

		if err != nil {
			fmt.Println(err)
		} else {
			for _, task := range tasks {
				if !task.Completed {
					fmt.Printf("%d. %s\n", task.ID, task.Task)
				}
			}
		}
	},
}

