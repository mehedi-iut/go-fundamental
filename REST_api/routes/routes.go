package routes

import "net/http"

func RegisterRoutes(sm *http.ServeMux) {
	sm.HandleFunc("GET /events", getEvents)
	sm.HandleFunc("GET /events/{id}", getEvent)
	sm.HandleFunc("POST /events", createEvent)
	sm.HandleFunc("PUT /events/{id}", updateEvent)
	sm.HandleFunc("DELETE /events/{id}", deleteEvent)
	sm.HandleFunc("POST /signup", signup)
	sm.HandleFunc("POST /login", login)
}
