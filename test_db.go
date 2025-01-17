package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func testDB() {

	// create a database object which can be used
	// to connect with database.
	// username:password@tcp(0.0.0.0:3306)/dbname
	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/test")

	// handle error, if any.
	if err != nil {
		panic(err)
	}

	// Now its time to connect with oru database,
	// database object has a method Ping.
	// Ping returns error, if unable connect to database.
	err = db.Ping()

	// handle error
	if err != nil {
		panic(err)
	}

	fmt.Print("Pong\n")

	// database object has a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
}

// testDB()
