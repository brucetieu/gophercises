package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var RootCmd = &cobra.Command{
	Use: "task",
	Short: "task is a CLI for managing your TODOs.",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

