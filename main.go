package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	
	"main0102/language"
	
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

	// Инициализируем пакет language
	language.InitDB(db)
	
	http.HandleFunc("/language", language.LanguageRead)
	http.HandleFunc("/language/", language.GetLanguageWrapper)
	
	fmt.Println("Сервер запущен на http://localhost:8182")
	log.Fatal(http.ListenAndServe(":8182", nil))
}