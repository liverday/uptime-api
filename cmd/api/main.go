package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"uptime-api/m/v2/cmd/api/deps"
	"uptime-api/m/v2/cmd/api/routes"
)

func main() {
	err := godotenv.Load()
	log.Println("Starting application")
	log.Println("Loading dependencies")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("APP_PORT")
	d := deps.NewDependencies()

	log.Printf("Dependencies loaded, starting server at port %s\n", port)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routes.Routes(d),
	}

	err = s.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
