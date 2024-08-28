package models

import (
	"time"

	"example.com/gin-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserId      int       `json:"userId"`
}

func (ev *Event) Save() error {
	query := `INSERT INTO events (name, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)` //INFO weuse this to secure the sql injection
	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	result, err := stm.Exec(ev.Name, ev.Description, ev.Location, ev.DateTime, ev.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	(*ev).ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	result, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var events []Event
	for result.Next() {
		var ev Event
		err = result.Scan(&ev.ID, &ev.Name, &ev.Description, &ev.Location, &ev.DateTime, &ev.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, ev)
	}
	return events, nil
}
