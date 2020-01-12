package main

import (
	"ProjectsAPI/app"
	"ProjectsAPI/controller"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Server started")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	router := controller.NewRouter()
	router.Use(app.JwtAuthentication)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
