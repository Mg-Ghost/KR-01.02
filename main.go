package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var db *sql.DB

func readPasswordFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func language(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT id, name FROM language")
	if err != nil {
		http.Error(w, "Ошибка базы данных: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Языки в базе данных:\n\n")

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Fprintf(w, "Ошибка чтения строки: %v\n", err)
			continue
		}
		fmt.Fprintf(w, "ID: %d, Name: %s\n", id, name)
	}

	if err = rows.Err(); err != nil {
		fmt.Fprintf(w, "Ошибка при обработке результатов: %v\n", err)
	}
}

func main() {
	password, err := readPasswordFromFile("pass.txt")
	if err != nil {
		log.Fatal("Ошибка чтения пароля:", err)
	}

	connStr := fmt.Sprintf("user=postgres password=%s dbname=LLP sslmode=disable", password)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка ping:", err)
	}

	fmt.Println("Успешно подключено к PostgreSQL (база LLP)!")

	http.HandleFunc("/language", language)
	fmt.Println("Сервер запущен на http://localhost:8182")
	log.Fatal(http.ListenAndServe(":8182", nil))
}