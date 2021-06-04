package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/partha-sen/ostd/authServer/controller"
	"github.com/partha-sen/ostd/authServer/db"
	"github.com/partha-sen/ostd/authServer/middleware"
	"github.com/partha-sen/ostd/authServer/token"
)

func main() {
	godotenv.Load()
	db.SetupDatabase()

	token.GlobalTokenStore = make(token.TokenStore)
	go token.RemoveExpiredToken()

	http.HandleFunc("/login", controller.HandleLogin)
	http.HandleFunc("/token/status/", controller.HandleTokenValidity)
	http.Handle("/logout", middleware.ValidateTokeMiddleware(controller.HandleLogOut))

	log.Fatal(http.ListenAndServe(":1000", nil))
}
