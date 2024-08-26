package models

import (
	"errors"
	"event-booking/db"
	"event-booking/utils"

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
	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		fmt.Print("An Error Occured")
	}
	res, err := db.DMLCommand(Query, u.Email, hash, u.Name)
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
func (u *User) Login() {
	Query := `SELECT ID,PASSWORD FROM USERS WHERE EMAIL=?`
	res := db.SelectRow(Query, u.Email)
	var password string
	err := res.Scan(&u.Id, &password)
	if err != nil {
		fmt.Println(err)
	}
	val := utils.CompareHash(password, u.Password)
	if !val {
		fmt.Println(err)
	}
}
func DeleteUser(id int64) (int64, error) {
	Query := `DELETE FROM USERS WHERE ID=?`
	res, err := db.DMLCommand(Query, id)
	if err != nil {
		fmt.Println("An Error Occured ")
		return 0, errors.New("an error occured")
	}
	return res.RowsAffected()
}
