package main

import (
	"net/http"
)

var fileHandler = func(resp http.ResponseWriter, req *http.Request) {
	if len(req.Header.Get("X-ENV-WEB")) > 0 && len(req.Header.Get("X-Redirect")) == 0 {
		env := req.Header.Get("X-ENV-WEB")
		http.FileServer(http.Dir("/root/static/sharedNoteWeb"+env)).ServeHTTP(resp, req)
		return
	}
	http.FileServer(http.Dir(`/root/static/sharedNoteWeb`)).ServeHTTP(resp, req)
}
