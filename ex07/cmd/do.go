package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brucetieu/gophercises/ex07/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(doTaskCmd)
}

var doTaskCmd = &cobra.Command{
	Use: "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		taskIdStr := strings.Join(args, "")
		taskId, _ := strconv.Atoi(taskIdStr)
		updatedTask, err := db.UpdateTask(taskId)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("You have completed the \"%s\" task.\n", updatedTask.Task)
		}
	},
}

