package language

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func LanguageRead(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		http.Error(w, "Database not initialized", http.StatusInternalServerError)
		return
	}
	
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	rows, err := db.Query("SELECT id, name FROM language")
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	
	Languages := []Language{}
	for rows.Next() {
		lang := Language{}
		err := rows.Scan(&lang.ID, &lang.Name)
		if err != nil {
			http.Error(w, "Error reading row: "+err.Error(), http.StatusInternalServerError)
			return
		}
		Languages = append(Languages, lang)
	}
	
	json.NewEncoder(w).Encode(Languages)
}

func GetLanguageWrapper(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	path := strings.TrimPrefix(r.URL.Path, "/language/")
	if path == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	
	idInt, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	
	GetLanguageByID(w, idInt)
}

func GetLanguageByID(w http.ResponseWriter, id int) {
	if db == nil {
		http.Error(w, "Database not initialized", http.StatusInternalServerError)
		return
	}
	
	row := db.QueryRow("SELECT id, name FROM language WHERE id = $1", id)

	lang := Language{}
	err := row.Scan(&lang.ID, &lang.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Language not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(lang)
}