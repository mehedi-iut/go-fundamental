package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"log"
	"net/http"
)

func main() {
	db.InitDB()
	sm := http.NewServeMux()
	routes.RegisterRoutes(sm)

	log.Printf("Listening on port 9090.....")
	err := http.ListenAndServe(":9090", sm)
	if err != nil {
		panic(err)
	}
}
