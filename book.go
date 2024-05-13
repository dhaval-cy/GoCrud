package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

func getBook(bookID int) (Book, error) {
	//Retrieve
	res := Book{}

	var id int
	var name string
	var author string
	var pages int
	var publicationDate mysql.NullTime

	err := db.QueryRow(`SELECT id, name, author, pages, publication_date FROM books where id = ?`, bookID).Scan(&id, &name, &author, &pages, &publicationDate)
	if err == nil {
		res = Book{ID: id, Name: name, Author: author, Pages: pages, PublicationDate: publicationDate.Time}
	}

	return res, err
}

func allBooks() ([]Book, error) {
	//Retrieve
	books := []Book{}

	rows, err := db.Query(`SELECT id, name, author, pages, publication_date FROM books order by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var author string
		var pages int
		var publicationDate mysql.NullTime

		err = rows.Scan(&id, &name, &author, &pages, &publicationDate)
		if err != nil {
			return books, err
		}

		currentBook := Book{ID: id, Name: name, Author: author, Pages: pages}
		if publicationDate.Valid {
			currentBook.PublicationDate = publicationDate.Time
		}

		books = append(books, currentBook)
	}

	return books, err
}

func insertBook(name, author string, pages int, publicationDate time.Time) (int64, error) {
	//Create
	var err error

	var ins *sql.Stmt
	// don't use _, err := db.Query()
	// func (db *DB) Prepare(query string) (*Stmt, error)
	ins, err = db.Prepare("INSERT INTO books(name, author, pages, publication_date) VALUES(?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer ins.Close()
	// func (s *Stmt) Exec(args ...interface{}) (Result, error)
	res, err := ins.Exec(name, author, pages, publicationDate)

	rowsAffec, _ := res.RowsAffected()
	if err != nil || rowsAffec != 1 {
		fmt.Println("Error inserting row:", err)
	}
	lastInserted, _ := res.LastInsertId()
	rowsAffected, _ := res.RowsAffected()
	fmt.Println("ID of last row inserted:", lastInserted)
	fmt.Println("number of rows affected:", rowsAffected)

	return lastInserted, err
}

func updateBook(id int, name, author string, pages int, publicationDate time.Time) (int, error) {
	//Create
	res, err := db.Exec(`UPDATE books set name=?, author=?, pages=?, publication_date=? where id=?`, name, author, pages, publicationDate, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

func removeBook(bookID int) (int, error) {
	//Delete
	res, err := db.Exec(`delete from books where id = ?`, bookID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
