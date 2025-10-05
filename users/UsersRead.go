package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func UsersRead(w http.ResponseWriter, r *http.Request) {
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
	
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		return
	}
	defer rows.Close()

	usersList := []Users{}
	for rows.Next() {
		user := Users{}
		err := rows.Scan(&user.ID, &user.Name, &user.Age) // Добавлено &user.Age
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Error reading row: " + err.Error()})
			return
		}
		usersList = append(usersList, user)
	}
	
	json.NewEncoder(w).Encode(usersList)
}

func GetUserWrapper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	
	if r.Method != "GET" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}
	
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	if path == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
		return
	}
	
	idInt, err := strconv.Atoi(path)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}
	
	GetUserByID(w, idInt)
}

func GetUserByID(w http.ResponseWriter, id int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	
	if db == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database not initialized"})
		return
	}
	
	row := db.QueryRow("SELECT id, name, age FROM users WHERE id = $1", id)

	user := Users{}
	err := row.Scan(&user.ID, &user.Name, &user.Age) 
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(map[string]string{"error": "User not found"}) 
		} else {
			json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		}
		return
	}
	
	json.NewEncoder(w).Encode(user)
}