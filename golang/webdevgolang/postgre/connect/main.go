package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// Install: https://www.postgresql.org/download/
// package database/sql is a generic interface around SQL or SQL-like databases
// it must be used in conjunction with a driver (which gets initialized when the program runs)
// we use the _ to load it, and not get the "unused" error. Though it is being used by the sql interface

// CREATE DATABASE bookstore;
// CREATE USER bond WITH PASSWORD 'password';
// GRANT ALL PRIVILEGES ON DATABASE bookstore to bond;

func main() {
	db, err := sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
}
