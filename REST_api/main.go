package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/middlewares"
	"example.com/rest-api/routes"
	"log"
	"net/http"
)

func main() {
	db.InitDB()

	//middleware := middlewares.

	router := http.NewServeMux()
	router.HandleFunc("GET /events", routes.GetEvents)
	router.HandleFunc("GET /events/{id}", routes.GetEvent)

	// protected route
	protectedRouter := http.NewServeMux()
	protectedRouter.HandleFunc("POST /events", routes.CreateEvent)
	protectedRouter.HandleFunc("PUT /events/{id}", routes.UpdateEvent)
	protectedRouter.HandleFunc("DELETE /events/{id}", routes.DeleteEvent)
	protectedRouter.HandleFunc("POST /events/{id}/register", routes.RegisterForEvent)
	protectedRouter.HandleFunc("DELETE /events/{id}/register", routes.CancelRegistrationForEvent)
	router.Handle("/", middlewares.Authenticate(protectedRouter))

	router.HandleFunc("POST /signup", routes.Signup)
	router.HandleFunc("POST /login", routes.Login)

	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}
	//sm := http.NewServeMux()
	//routes.RegisterRoutes(sm)

	log.Printf("Listening on port 9090.....")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
	//err := http.ListenAndServe(":9090", sm)
	//if err != nil {
	//	panic(err)
	//}
}
