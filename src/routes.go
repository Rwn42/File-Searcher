package main

import (
	"File-Search/db"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Results struct {
	Amount int
	Data   map[string]int
}

func uploadRoute(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1024 * 16); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Form Structure"))
		return
	}

	fp, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could Not Read File"))
		return
	}
	defer fp.Close()

	savePath := cfg.SaveDirectory + "/" + header.Filename
	createPath := "public/" + savePath

	newFile, err := os.Create(createPath)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	if _, err = io.Copy(newFile, fp); err != nil {
		log.Fatal(err)
	}

	if hasTags := r.Form.Has("topics"); !hasTags {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No Tags Provided"))
		return
	}

	topics := strings.Split(r.Form.Get("topics"), ",")

	if err := db.InsertFileEntry(db.DB, savePath, topics); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	w.WriteHeader(http.StatusOK)

}

func searchRoute(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()
	if !urlParams.Has("query") {
		w.WriteHeader(http.StatusOK)
		bytes, _ := json.Marshal(Results{0, nil})
		w.Write(bytes)
	}
	query := urlParams.Get("query")
	tags := strings.Split(query, ",")

	data := db.FindFileByTags(db.DB, tags, urlParams.Get("dateStart"), urlParams.Get("dateEnd"))
	bytes, err := json.Marshal(Results{len(data), data})
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}
