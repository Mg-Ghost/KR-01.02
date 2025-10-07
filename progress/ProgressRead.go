package progress

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ProgressRead(w http.ResponseWriter, r *http.Request) {
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
	rows, err := db.Query("SELECT id, namelanguage, nameusers, progress FROM progress")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		return
	}
	defer rows.Close()

	progressList := []Progress{}
	for rows.Next() {
		progress := Progress{}
		err := rows.Scan(&progress.ID, &progress.NameLanguage, &progress.NameUsers, &progress.Progress) 
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Error reading row: " + err.Error()})
			return
		}
		progressList = append(progressList, progress)
	}
	
	json.NewEncoder(w).Encode(progressList)
}

func GetProgressWrapper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	
	if r.Method != "GET" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}
	
	path := strings.TrimPrefix(r.URL.Path, "/progress/")
	if path == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
		return
	}
	
	idInt, err := strconv.Atoi(path)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}
	
	GetProgressByID(w, idInt)
}

func GetProgressByID(w http.ResponseWriter, id int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	
	if db == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database not initialized"})
		return
	}
	
	row := db.QueryRow("SELECT id, namelanguage, nameusers, progress FROM progress WHERE id = $1", id)

	progress := Progress{}
	err := row.Scan(&progress.ID, &progress.NameLanguage, &progress.NameUsers, &progress.Progress) 
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(map[string]string{"error": "Progress not found"}) 
		} else {
			json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		}
		return
	}
	
	json.NewEncoder(w).Encode(progress)
}