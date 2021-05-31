package main

import (
	"log"
	"net/http"

	"github.com/partha-sen/ostd/webservice/controller"
	"github.com/partha-sen/ostd/webservice/db"
)

func main() {
	db.SetupDatabase()

	http.HandleFunc("/openings", controller.HandleOpenings)
	http.HandleFunc("/openings/", controller.HandleOpening)

	if err := http.ListenAndServe(":1000", nil); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %v", err)
	}
}
