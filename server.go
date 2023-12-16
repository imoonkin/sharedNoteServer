package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.OpenFile("note.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	defer f.Close()

	initDB()
	http.HandleFunc("/", fileHandler)
	http.HandleFunc("/range", rangeHandler)
	http.HandleFunc("/add", addNoteHandler)

	log.Fatal(http.ListenAndServe(":12808", nil)) // todo tls
}
