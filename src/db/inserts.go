package db

import (
	"database/sql"
	"log"

	"github.com/mattn/go-sqlite3"
)

func InsertFileEntry(db *sql.DB, path string, tags []string) error {
	file_id, err := insertFile(db, path)
	if err != nil {
		return err
	}

	for _, tag := range tags {
		tag_id, has := findTag(db, tag)
		if !has {
			tag_id, err = insertTag(db, tag)
			if err != nil {
				return err
			}
		}
		insertFileTag(db, file_id, tag_id) //update the junction table
	}
	return nil
}

func insertFile(db *sql.DB, path string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO files (file_path) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(path)
	if err != nil {
		if err != sqlite3.ErrConstraintUnique {
			log.Fatal(err)
		}
		return 0, err
	}

	file_id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return file_id, nil
}

func insertTag(db *sql.DB, tag string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO tags (name) VALUES (?)")
	if err != nil {
		log.Fatalf("Could not prepare statement: %s", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(tag)
	if err != nil {
		if err != sqlite3.ErrConstraintUnique {
			log.Fatal(err)
		}
		return 0, err
	}

	tag_id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return tag_id, nil
}

func insertFileTag(db *sql.DB, file_id int64, tag_id int64) {
	stmt, err := db.Prepare("INSERT INTO file_tag (file_id, tag_id) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(file_id, tag_id); err != nil {
		log.Fatal(err)
	}

}
