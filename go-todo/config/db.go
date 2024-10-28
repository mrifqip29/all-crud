package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		fmt.Println("error bang")
		log.Fatal(err)
	}

	// createTable := `
	// CREATE TABLE IF NOT EXISTS todos (
	//     id INTEGER PRIMARY KEY AUTOINCREMENT,
	//     title TEXT,
	//     completed BOOLEAN
	// );`

	// if _, err := db.Exec(createTable); err != nil {
	// 	log.Fatal(err)
	// }

	return db
}
