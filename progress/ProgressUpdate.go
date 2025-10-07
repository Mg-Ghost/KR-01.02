package progress

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ProgressUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != "PATCH" && r.Method != "PUT" {
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

	var updateData struct {
		Progress string `json:"progress"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON: " + err.Error()})
		return
	}
	defer r.Body.Close()

	_, err = db.Exec("UPDATE progress SET progress = $1 WHERE id = $2", updateData.Progress, idInt)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Progress updated successfully"})
}