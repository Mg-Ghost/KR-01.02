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
		NameLanguage *string `json:"namelanguage"`
		NameUsers    *string `json:"nameusers"`
		Progress     *string `json:"progress"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON: " + err.Error()})
		return
	}
	defer r.Body.Close()

	if updateData.NameLanguage == nil && updateData.NameUsers == nil && updateData.Progress == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "No fields to update"})
		return
	}

	query := "UPDATE progress SET"
	params := []interface{}{}
	paramCount := 1

	if updateData.NameLanguage != nil {
		query += " namelanguage = $" + strconv.Itoa(paramCount) + ","
		params = append(params, *updateData.NameLanguage)
		paramCount++
	}

	if updateData.NameUsers != nil {
		query += " nameusers = $" + strconv.Itoa(paramCount) + ","
		params = append(params, *updateData.NameUsers)
		paramCount++
	}

	if updateData.Progress != nil {
		query += " progress = $" + strconv.Itoa(paramCount) + ","
		params = append(params, *updateData.Progress)
		paramCount++
	}

	query = query[:len(query)-1] + " WHERE id = $" + strconv.Itoa(paramCount)
	params = append(params, idInt)

	_, err = db.Exec(query, params...)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Progress updated successfully"})
}