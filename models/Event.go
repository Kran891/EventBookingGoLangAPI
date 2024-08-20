package models

import (
	"event-booking/db"
	"fmt"
	"time"
)

var Events []Event

type Event struct {
	Id          int64
	Name        string
	Description string
	Location    string
	CreatedDate time.Time
	UserId      int
}

func (e *Event) Save() {
	const Query = `
     INSERT INTO EVENTS (NAME,Description,Location,CreatedDate,UserId)
	 VALUES (?,?,?,?,?)
	`
	res, err := db.DMLCommand(Query, e.Name, e.Description, e.Location, time.Now(), e.UserId)
	if err != nil {
		panic("An Error Occured While Creating an Event...")
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("An Error Occured")
	}
	e.Id = id
}
func GetAllEvents() []Event {
	const Query = `select * from Events`
	var events []Event

	data, err := db.SelectRows(Query)
	if err != nil {
		panic(err)
	}
	defer data.Close()
	for data.Next() {
		var event Event
		data.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.CreatedDate, &event.UserId)
		events = append(events, event)
	}
	fmt.Println(events)
	return events
}
func (e Event) Update() (int64, error) {
	const Query = `update events set name=?,Description=?,Location=?,UserId=? where Id=?`
	res, err := db.DMLCommand(Query, e.Name, e.Description, e.Location, e.UserId, e.Id)
	if err != nil {
		panic("An Error Occured...")
	}
	return res.LastInsertId()
}
func GetById(id int64) (event Event) {
	const Query = `SELECT * FROM Events WHERE ID=?`
	res := db.SelectRow(Query, id)

	res.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.CreatedDate, &event.UserId)

	return
}
func DeleteById(id int64) (int64, error) {
	const Query = `DELETE FROM EVENTS WHERE ID=?`
	res, err := db.DMLCommand(Query, id)
	if err != nil {
		panic(err)
	}
	return res.LastInsertId()
}
