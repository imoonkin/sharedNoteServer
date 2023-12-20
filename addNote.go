package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var addNoteHandler = func(resp http.ResponseWriter, req *http.Request) {
	if forward(resp, req) {
		return
	}
	var n note
	if err := json.NewDecoder(req.Body).Decode(&n); err != nil || len(n.Title) == 0 || len(n.Address) == 0 {
		resp.WriteHeader(http.StatusForbidden)
		resp.Write([]byte("不太对"))
		log.Fatalf("err=%+v , title=%+v, addr=%+v", err, n.Title, n.Address)
		return
	}

	// todo valid range check
	// todo update DB ^^^
	n.Latitude = float64(int64(n.Latitude*1000000)) / 1000000
	n.Longitude = float64(int64(n.Longitude*1000000)) / 1000000
	if err := addNote(n); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("{}"))
}
