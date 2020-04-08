package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Book is an interface to table "books" in bookstore database
type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	dbConnectionTemplate := "postgres://%s:%s@localhost/%s?sslmode=disable"
	dbUser := "bond"
	dbPassword := "password"
	dbName := "bookstore"

	connectionString := fmt.Sprintf(dbConnectionTemplate, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ping db if pointer to db returned
	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Printf("Successfully connected to %s\n", dbName)

	rows, err := db.Query(`SELECT * FROM books`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}

	fmt.Printf("Books slice: %+v\n", bks)

	for _, bk := range bks {
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}
