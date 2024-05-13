package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	tmpDB, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/books_database")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	http.HandleFunc("/", handleListBooks)
	http.HandleFunc("/book.html", handleViewBook)
	http.HandleFunc("/save", handleSaveBook)
	http.HandleFunc("/delete", handleDeleteBook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
