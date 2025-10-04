package language

import (
	"encoding/json"
	"net/http"
)

func LanguageCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodPost {
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	var lang Language
	if err := json.NewDecoder(r.Body).Decode(&lang); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON: " + err.Error()})
		return
	}
	defer r.Body.Close()

	_, err := db.Exec("INSERT INTO language (name) VALUES ($1)", lang.Name)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		return
	}

	json.NewEncoder(w).Encode(lang)
}