package handlers

import (
	"encoding/json"
	"net/http"
)

type Classroom struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListClassrooms(w http.ResponseWriter, _ *http.Request) {
	items := []Classroom{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
