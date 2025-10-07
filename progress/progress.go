package progress

import "database/sql"

type Progress struct {
	ID           int    `json:"id"`
	NameLanguage string `json:"namelanguage"`
	NameUsers    string `json:"nameusers"`
	Progress     string `json:"progress"`
}

var db *sql.DB

func InitDB(database *sql.DB) {
	db = database
}