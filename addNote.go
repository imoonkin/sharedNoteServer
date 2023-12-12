package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var addNoteHandler = func(resp http.ResponseWriter, req *http.Request) {
	var n note
	if err := json.NewDecoder(req.Body).Decode(&n); err != nil || len(n.title) == 0 || len(n.address) == 0 {
		resp.WriteHeader(http.StatusForbidden)
		return
	}

	// todo valid range check
	// todo update DB
	if err := addNote(n); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
