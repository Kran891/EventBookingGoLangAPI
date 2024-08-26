package models

import (
	"errors"
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
func GetUserById(id int64) (user User) {
	Query := `SELECT ID,EMAIL,PASSWORD,NAME FROM USERS WHERE ID=?`
	res := db.SelectRow(Query, id)
	res.Scan(&user.Id, &user.Email, &user.Password, &user.Name)
	return
}
func DeleteUser(id int64) (int64, error) {
	Query := `DELETE FROM USER WHERE ID=?`
	res, err := db.DMLCommand(Query, id)
	if err != nil {
		fmt.Println("An Error Occured ")
		return 0, errors.New("an error occured")
	}
	return res.RowsAffected()
}
