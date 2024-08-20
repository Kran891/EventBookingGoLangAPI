package models

import (
	"event-booking/db"
	"fmt"
)

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
}

func (u *User) Save() {
	Query := `
	INSERT INTO USERS (EMAIL,PASSWORD,NAME) VALUES (?,?,?)
	`
	res, err := db.DMLCommand(Query, u.Email, u.Password, u.Name)
	if err != nil {
		fmt.Println(err)
	}
	u.Id, _ = res.LastInsertId()
}
func (u *User) Update() {
	Query := `UPDATE USERS SET NAME=?,EMAIL=? WHERE ID=?`
	res, err := db.DMLCommand(Query, u.Name, u.Email, u.Id)
	if err != nil {
		fmt.Println("An Error Occured while UPdating User...")
	}
	u.Id, _ = res.LastInsertId()
}
