package handler

import (
	"html/template"
	"net/http"

	reps "github.com/brucetieu/gophercises/ex03/representations"
	log "github.com/sirupsen/logrus"
)

var (
	Intro = "intro"
	NewYork = "new-york"
	Debate = "debate"
	SeanKelly = "sean-kelly"
	MarkBates = "mark-bates"
	Denver = "denver"
	Home = "home"
)

type Handler struct {
	tmpl *template.Template
	storyMap map[string]reps.Page
}

func NewHandler(tmpl *template.Template, storyMap map[string]reps.Page) *Handler {
	return &Handler{
		tmpl: tmpl,
		storyMap: storyMap,
	}
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Info("path: " + path)
	if path == "/" || path == "" {
		path = "/" + Intro
	} 

	h.tmpl.Execute(w, h.storyMap[path[1:]])
}
