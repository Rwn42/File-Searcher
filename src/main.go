package main

import (
	"File-Search/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := loadConfig()

	db.Open()
	//db.Create() //enable this line to create a new db when the server is run
	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir("./public")))

	http.HandleFunc("/upload", uploadRoute)
	http.HandleFunc("/search", searchRoute)

	http.ListenAndServe(":"+fmt.Sprint(cfg.Port), nil)
}
