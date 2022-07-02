package utils

import (
	"encoding/json"
	"io/ioutil"

	reps "github.com/brucetieu/gophercises/ex03/representations"
	log "github.com/sirupsen/logrus"
)

func ParseJSONStory(filename string) map[string]reps.Page {
	jsonBytes, _ := ioutil.ReadFile(filename)

	var storyMap map[string]reps.Page

	err := json.Unmarshal(jsonBytes, &storyMap)
	if err != nil {
		log.Fatal(err.Error())
	}

	return storyMap
}

func PrettyFormat(data interface{}) string {
	bytes, _ := json.MarshalIndent(data, "", "\t")
	return string(bytes)
}