package main

import (
	"encoding/json"
	"net/http"
)

type Opening struct {
	Id         int
	Name       string
	Company    string
	Position   string
	Skills     string
	Experience int
}

func (o *Opening) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	strOpening, err := json.Marshal(o)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(strOpening)
}
