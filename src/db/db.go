package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB = nil

const create_file = "./src/db/create.sql"

func Open() {
	var err error

	DB, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Unable to open database: %s \n", err)
	}

}

func Create() {
	query, err := os.ReadFile(create_file)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := DB.Exec(string(query)); err != nil {
		log.Fatalf("Could not create database from file %s because of %s", create_file, err)
	}

}

func Close() {
	DB.Close()
}
