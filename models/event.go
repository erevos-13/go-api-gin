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
	UserId      int64     `json:"userId"`
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

func GetEventById(id int64) (Event, error) {
	query := `SELECT * FROM events WHERE id=?`
	result := db.DB.QueryRow(query, id)

	var event Event
	err := result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (ev *Event) UpdateEvent() error {
	query := `UPDATE events SET name=?, description=?, location=?, date_time=? WHERE id=?`
	stem, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stem.Close()
	_, err = stem.Exec(ev.Name, ev.Description, ev.Location, ev.DateTime, ev.ID)

	return err
}

func (ev *Event) DeleteEvent() error {
	query := `DELETE FROM events WHERE id=?`
	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(ev.ID)
	return err

}

func (ev *Event) Register(userId int64) error {
	query := `INSERT INTO registerations (event_id, user_id) VALUES (?, ?)`
	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(ev.ID, userId)
	return err
}
func (ev Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM registerations WHERE event_id=? AND user_id=?`
	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(ev.ID, userId)
	return err
}
