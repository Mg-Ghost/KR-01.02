package users

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func UsersDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "DELETE"{
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
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

	result, err := db.Exec("DELETE FROM users WHERE id = $1", idInt)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		json.NewEncoder(w).Encode(map[string]string{"error": "Users not found"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Users deleted successfuly"})
}