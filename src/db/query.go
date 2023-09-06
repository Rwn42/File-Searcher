package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// FindFileByTags returns a map of filepaths to inputted tags that matched
// example: if file a.txt has tags cat dog horse and the function is called with cat, bison as inputs
// the map would contain "a.txt": 1 bc only the cat tag matched any files
func FindFileByTags(db *sql.DB, tags []string, dateStart string, dateEnd string) map[string]int {
	tag_ids := make([]int64, len(tags))

	//populate tag_ids with all tags that match in our database
	added := 0
	for _, tag := range tags {
		tag_id, exists := findTag(db, tag)
		if !exists {
			continue
		}
		tag_ids[added] = tag_id
		added += 1
	}

	paths := make(map[string]int)

	var stmt *sql.Stmt
	var err error
	if dateEnd == "" || dateStart == "" {
		stmt, err = db.Prepare("SELECT files.file_path FROM files INNER JOIN file_tag ON file_tag.file_id = files.id WHERE file_tag.tag_id=?")
	} else {
		stmt, err = db.Prepare(
			"SELECT files.file_path FROM files INNER JOIN file_tag ON file_tag.file_id = files.id WHERE file_tag.tag_id=?" +
				"AND files.created BETWEEN ? AND ?")
	}
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, tag_id := range tag_ids {
		rows, err := stmt.Query(tag_id, dateStart, dateEnd)
		if err != nil {
			log.Fatal(err)
		}

		for {
			if !rows.Next() {
				break
			}

			var file_path string
			if rows.Scan(&file_path) != nil {
				log.Fatalf("Could not scan row: %s", err)
			}

			paths[file_path] += 1

		}
	}

	return paths

}

func findTag(db *sql.DB, tag string) (int64, bool) {
	stmt, err := db.Prepare("SELECT id FROM tags WHERE name = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var tag_id int64

	err = stmt.QueryRow(tag).Scan(&tag_id)

	if err == sql.ErrNoRows {
		return 0, false
	} else if err != nil {
		log.Fatal(err)
	}

	return tag_id, true
}
