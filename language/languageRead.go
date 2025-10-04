package language

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func LanguageRead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	
	if db == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database not initialized"})
		return
	}
	
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}
	
	rows, err := db.Query("SELECT id, name FROM language")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		return
	}
	defer rows.Close()

	languages := []Language{}
	for rows.Next() {
		lang := Language{}
		err := rows.Scan(&lang.ID, &lang.Name)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Error reading row: " + err.Error()})
			return
		}
		languages = append(languages, lang)
	}
	
	json.NewEncoder(w).Encode(languages)
}

func GetLanguageWrapper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	
	if r.Method != "GET" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}
	
	path := strings.TrimPrefix(r.URL.Path, "/language/")
	if path == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
		return
	}
	
	idInt, err := strconv.Atoi(path)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}
	
	GetLanguageByID(w, idInt)
}

func GetLanguageByID(w http.ResponseWriter, id int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	
	if db == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database not initialized"})
		return
	}
	
	row := db.QueryRow("SELECT id, name FROM language WHERE id = $1", id)

	lang := Language{}
	err := row.Scan(&lang.ID, &lang.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(map[string]string{"error": "Language not found"})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		}
		return
	}
	
	json.NewEncoder(w).Encode(lang)
}