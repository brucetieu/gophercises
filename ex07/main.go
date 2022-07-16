package main

import (
	"github.com/brucetieu/gophercises/ex07/cmd"
	"github.com/brucetieu/gophercises/ex07/db"
)

func main() {
	db.InitDB()
	defer db.DB.Close()
	cmd.RootCmd.Execute()
}