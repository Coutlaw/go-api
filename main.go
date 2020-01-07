package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-api/app"
	"go-api/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	// ROUTES

	// Create a user (gets a jwt)
	router.HandleFunc("/api/users/new", controllers.CreateAccount).Methods("POST")

	// Get a new JWT if users is expired
	router.HandleFunc("/api/users/login", controllers.Authenticate).Methods("POST")

	// Create contacts for a user
	router.HandleFunc("/api/users/contacts", controllers.CreateContact).Methods("POST")

	// Get all contacts for a user
	router.HandleFunc("/api/users/contacts", controllers.GetContacts).Methods("GET")

	// Get a contact by ID that belongs to a User
	router.HandleFunc("/api/users/contacts/{contactId}", controllers.GetContactById).Methods("GET")


	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println("Listening on port: " + port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
