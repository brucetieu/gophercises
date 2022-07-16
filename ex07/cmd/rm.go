package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brucetieu/gophercises/ex07/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(rmTaskCmd)
}

var rmTaskCmd = &cobra.Command{
	Use: "rm",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskIdString := strings.Join(args, "")
		taskId, _ := strconv.Atoi(taskIdString)
		deletedTask, err := db.DeleteTask(taskId)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("You have deleted the \"%s\" task\n", deletedTask.Task)
		}
	},
}

