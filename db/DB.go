package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() *sql.DB {
	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {

		panic(err)
	} else {
		fmt.Println("☑️ ☑️ Connected to DB")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	return DB
}
func CreateTables() {
	const EventsTABLE = `
	CREATE TABLE IF NOT EXISTS EVENTS(
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	Name        TEXT,
	Description TEXT,
	Location    TEXT,
	CreatedDate DATETIME,
	UserId    INTEGER
	)
	`
	const UsersTABLE = `
	CREATE TABLE IF NOT EXISTS USERS(
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	NAME TEXT,
	EMAIL TEXT,
	PASSWORD TEXT
	)
	`
	_, err := DB.Exec(EventsTABLE)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(UsersTABLE)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("☑️☑️Table Created Successfully...")
	}

}
func DMLCommand(query string, data ...any) (sql.Result, error) {
	stmt, err := DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(data...)
	return res, err

}
func SelectRows[T *sql.Rows](query string, data ...any) (T, error) {
	rows, err := DB.Query(query, data...)
	return rows, err
}
func SelectRow(query string, data ...any) *sql.Row {
	row := DB.QueryRow(query, data...)
	return row
}
