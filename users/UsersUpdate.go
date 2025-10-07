package users

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func UsersUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != "PATCH" && r.Method != "PUT" {
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

	var updateData struct {
		Name *string `json:"name"`
		Age  *int    `json:"age"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON: " + err.Error()})
		return
	}
	defer r.Body.Close()

	if updateData.Name == nil && updateData.Age == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "No fields to update"})
		return
	}

	query := "UPDATE users SET"
	params := []interface{}{}
	paramCount := 1

	if updateData.Name != nil {
		query += " name = $" + strconv.Itoa(paramCount) + ","
		params = append(params, *updateData.Name)
		paramCount++
	}

	if updateData.Age != nil {
		query += " age = $" + strconv.Itoa(paramCount) + ","
		params = append(params, *updateData.Age)
		paramCount++
	}

	query = query[:len(query)-1] + " WHERE id = $" + strconv.Itoa(paramCount)
	params = append(params, idInt)

	_, err = db.Exec(query, params...)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}