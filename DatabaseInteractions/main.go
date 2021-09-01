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

func InsertAnAlbum(db *sql.DB, album *Album) error {
	_, err := db.Exec("INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)", album.title, album.artist, album.price)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
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

	newAlbum := Album{
		title:  "A Love Supreme",
		artist: "John Coltrane",
		price:  199.82,
	}
	err = InsertAnAlbum(db, &newAlbum)
	if err != nil {
		log.Fatal(err)
	}
	allAlbums = ReadAllAlbums(rows)
	fmt.Println(allAlbums)
}
