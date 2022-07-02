package main

import (
	"html/template"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/brucetieu/gophercises/ex03/handler"
	"github.com/brucetieu/gophercises/ex03/utils"
)


func main() {

	tmpl := template.Must(template.ParseFiles("story.html"))
	jsonFile := "gopher.json"

	handlers := handler.NewHandler(tmpl, utils.ParseJSONStory(jsonFile))
	http.HandleFunc("/", handlers.IndexHandler)

	// Serve my css at /css/
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	
	log.Info("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}