package cmd

import (
	"fmt"
	"strings"

	"github.com/brucetieu/gophercises/ex07/db"
	"github.com/brucetieu/gophercises/ex07/models"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(addTaskCmd)
}

var addTaskCmd = &cobra.Command{
	Use: "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		task := models.Task{
			Task: strings.Join(args, " "),
			Completed: false,
		}
		err := db.CreateTask(&task)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Added task: %s\n", strings.Join(args, " "))
		}
	},
}

