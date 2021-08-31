package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Album struct {
	id     int
	title  string
	artist string
	price  float64
}

func GetRows(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func GetColumns(rows *sql.Rows) []string {
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	return cols
}

func ReadAllAlbums(rows *sql.Rows) *[]Album {
	allAlbums := make([]Album, 0)
	var iterator int = 0

	for rows.Next() {
		album := Album{}
		err := rows.Scan(&album.id, &album.title, &album.artist, &album.price)
		if err != nil {
			log.Fatal(err)
		}
		allAlbums = append(allAlbums, album)
		iterator++
	}

	return &allAlbums
}

func main() {
	connStr := "user=<your-db-user> password='<your-db-password>' dbname=<your-db-name>"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows := GetRows(db)
	cols := GetColumns(rows)
	fmt.Println(cols)
	allAlbums := ReadAllAlbums(rows)
	fmt.Println(allAlbums)
}
