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
	db.Create()
	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir("./public")))

	http.HandleFunc("/upload", uploadRoute)
	http.HandleFunc("/search", searchRoute)

	if cfg.HostSaveDirectory {
		http.Handle("/files", http.FileServer(http.Dir(cfg.SaveDirectory)))
	}

	http.ListenAndServe(":"+fmt.Sprint(cfg.Port), nil)
}
