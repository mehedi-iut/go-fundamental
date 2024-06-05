package routes

import (
	"encoding/json"
	"example.com/rest-api/models"
	"fmt"
	"net/http"
	"strconv"
)

func RegisterForEvent(w http.ResponseWriter, r *http.Request) {
	eventId, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId := r.Context().Value("userId").(int64)

	event, err := models.GetEventByID(uint(eventId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = event.Register(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := map[string]string{"message": "Registration successful"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func CancelRegistrationForEvent(w http.ResponseWriter, r *http.Request) {
	eventId, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId := r.Context().Value("userId").(int64)

	var event models.Event
	event.ID = uint(eventId)
	err = event.CancelRegistration(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Registration cancelled successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
		return
	}
}
