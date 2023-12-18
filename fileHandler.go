package main

import (
	"net/http"
)

var fileHandler = func(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(200)
	resp.Write([]byte("hahahaha"))
	return
	if len(req.Header.Get("X-ENV-WEB")) > 0 && len(req.Header.Get("X-Redirect")) == 0 {
		env := req.Header.Get("X-ENV-WEB")
		http.FileServer(http.Dir("/root/sharedNoteWeb"+env)).ServeHTTP(resp, req)
		return
	}
	http.FileServer(http.Dir(`/root/sharedNoteWeb`)).ServeHTTP(resp, req)
}
