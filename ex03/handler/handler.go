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
	log.Info("path: " + r.URL.Path)
	
	switch r.URL.Path {
	case "/":
		h.tmpl.Execute(w, h.storyMap[Intro])
	case "/" + NewYork:
		h.tmpl.Execute(w, h.storyMap[NewYork])
	case "/" + Debate:
		h.tmpl.Execute(w, h.storyMap[Debate])
	case "/" + SeanKelly:
		h.tmpl.Execute(w, h.storyMap[SeanKelly])
	case "/" + MarkBates:
		h.tmpl.Execute(w, h.storyMap[MarkBates])
	case "/" + Denver:
		h.tmpl.Execute(w, h.storyMap[Denver])
	case "/" + Home:
		h.tmpl.Execute(w, h.storyMap[Home])
	}
	
}
