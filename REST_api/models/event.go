package models

import (
	"example.com/rest-api/db"
	"log"
	"time"
)

type Event struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name" required:"true"`
	Description string    `json:"description" required:"true"`
	Location    string    `json:"location" required:"true"`
	DateTime    time.Time `json:"datetime" required:"true"`
	UserID      uint      `json:"user_id"`
}

var events = []Event{} // later on save to database

func (e Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	var id int
	//result, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	err = stmt.QueryRow(e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&id)
	defer stmt.Close()
	if err != nil {
		return err
	}
	//id, err := result.LastInsertId()
	// lib/pq doesn't support LastInsertId

	e.ID = uint(id)
	return err
}

// Get all event, as we return all event
// it should not be a method of specific event rather a normal function
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		e := Event{}
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func GetEventByID(id uint) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"
	row := db.DB.QueryRow(query, id)
	e := Event{}
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name=$1, description=$2, location=$3, datetime=$4
	WHERE id=$5
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

func (e Event) Delete() error {
	query := "DELETE FROM events WHERE id = $1"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}
