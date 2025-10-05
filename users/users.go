package users

import "database/sql"

type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB

func InitDB(database *sql.DB) {
	db = database
}