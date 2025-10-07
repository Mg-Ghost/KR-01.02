package language

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func LanguageUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != "PATCH" && r.Method != "PUT" {
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

	var updateData struct {
		Name *string `json:"name"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON: " + err.Error()})
		return
	}
	defer r.Body.Close()

	if updateData.Name == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "No fields to update"})
		return
	}

	_, err = db.Exec("UPDATE language SET name = $1 WHERE id = $2", *updateData.Name, idInt)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Language updated successfully"})
}