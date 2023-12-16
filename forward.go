package main

import (
	"io"
	"log"
	"net/http"
)

var svc = "shared-note-server"

func forward(resp http.ResponseWriter, req *http.Request) (needForward bool) {
	if !(len(req.Header.Get("X-ENV")) > 0 && len(req.Header.Get("X-Redirect")) == 0) {
		return false
	}
	needForward = true
	env := req.Header.Get("X-ENV")
	url := req.URL
	log.Printf("copied url=%+v , %+v, %+v, %+v, %+v, %+v", url, url.Scheme, url.Host, url.Opaque, url.User, url.OmitHost)
	url.Scheme = "http"
	url.Host = svc + env

	proxyReq, err := http.NewRequest(req.Method, url.String(), req.Body)
	if err != nil {
		log.Printf("wtf err1=%+v", err)
	}

	proxyReq.Header.Set("Host", svc+env)
	proxyReq.Header.Set("X-Redirect", "1")
	proxyReq.Header.Set("X-Forwarded-For", req.RemoteAddr)

	for header, values := range req.Header {
		for _, value := range values {
			proxyReq.Header.Add(header, value)
		}
	}

	client := &http.Client{}
	proxyRes, err := client.Do(proxyReq)
	if err != nil {
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte(err.Error()))
		log.Printf("wtf err2=%+v", err)
		return
	}
	resp.WriteHeader(proxyRes.StatusCode)
	proxyBody, err := io.ReadAll(proxyRes.Body)
	if err != nil {
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte(err.Error()))
		log.Printf("wtf err3=%+v", err)
		return
	}

	resp.Write(proxyBody)
	return
}
