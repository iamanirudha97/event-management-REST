package models

import (
	"fmt"
	"log"
	"time"

	"example.com/eventbrite/db"
)

type Event struct {
	Id          int
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	UserId      int       `binding: "required"`
}

func (e Event) SaveEvent() {
	query := `INSERT INTO events(name, description, location, datetime, userId) 
				VALUES($1, $2, $3, $4, $5)`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		log.Fatal(err)
	}

	result.LastInsertId()
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	fmt.Println(events)
	return events, nil
}
