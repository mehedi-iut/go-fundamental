package routes

import (
	"encoding/json"
	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := models.GetAllEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event, err := models.GetEventByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(event); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event.DateTime = time.Now()
	event.UserID = uint(userId)
	err = event.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		return
	}
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId := r.Context().Value("userId").(int64)
	e, err := models.GetEventByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if e.UserID != uint(userId) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event.ID = uint(id)
	//event.DateTime = time.Now()
	err = event.Update()
	if err != nil {
		fmt.Println(err)
		return
	}

	response := map[string]string{"message": "Event updated"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId := r.Context().Value("userId").(int64)

	event, err := models.GetEventByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if event.UserID != uint(userId) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = event.Delete()
	if err != nil {
		fmt.Println(err)
		return
	}

	response := map[string]string{"message": "Event deleted"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
		return
	}

}
