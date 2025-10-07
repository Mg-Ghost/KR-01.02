package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"main0102/language"
	"main0102/progress"
	"main0102/users"

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

	language.InitDB(db)
	users.InitDB(db)
	progress.InitDB(db)

	http.HandleFunc("/language", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			language.LanguageRead(w, r)
		case "POST":
			language.LanguageCreate(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/language/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			language.GetLanguageWrapper(w, r)
		case "DELETE":
			language.LanguageDelete(w, r)
		case "PATCH", "PUT":
        	language.LanguageUpdate(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			users.UsersRead(w, r)
		case "POST":
			users.UsersCreate(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			users.GetUserWrapper(w, r)
		case "DELETE":
			users.UsersDelete(w, r)
		case "PATCH", "PUT":
        	users.UsersUpdate(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/progress", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			progress.ProgressRead(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/progress/", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        progress.GetProgressWrapper(w, r)
    case "PATCH", "PUT":
        progress.ProgressUpdate(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
	})

	fmt.Println("Сервер запущен на http://localhost:8182")
	log.Fatal(http.ListenAndServe(":8182", nil))
}