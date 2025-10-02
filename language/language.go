package language

import (
	"database/sql"
)

type Language struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func InitDB(database *sql.DB) {
	db = database
}