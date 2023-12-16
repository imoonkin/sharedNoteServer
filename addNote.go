package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var addNoteHandler = func(resp http.ResponseWriter, req *http.Request) {
	if len(req.Header.Get("X-ENV")) > 0 && len(req.Header.Get("X-Redirect")) == 0 {
		env := req.Header.Get("X-ENV")
		req.Header.Set("X-Redirect", "1")
		req.URL.Host = "shared-note-server" + env
		client := http.Client{}
		respproxy, _ := client.Do(req)
		resp.WriteHeader(respproxy.StatusCode)
		body, _ := io.ReadAll(respproxy.Body)
		resp.Write(body)
		return
	}
	var n note
	if err := json.NewDecoder(req.Body).Decode(&n); err != nil || len(n.title) == 0 || len(n.address) == 0 {
		resp.WriteHeader(http.StatusForbidden)
		return
	}

	// todo valid range check
	// todo update DB ^^^
	if err := addNote(n); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
