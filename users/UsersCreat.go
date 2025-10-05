package users

import (
	"encoding/json"
	"net/http"
)

func UsersCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != "POST" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	var user Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON: " + err.Error()})
		return
	}
	defer r.Body.Close()

	_, err := db.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		return
	}

	json.NewEncoder(w).Encode(user)
}