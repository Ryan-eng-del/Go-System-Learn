package http_server

import (
	"net/http"
	"time"
)


func HttpCustomServer() {
	server := http.Server{
		Handler: &HttpHandler{},
		Addr: "localhost:8888",
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 3 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		MaxHeaderBytes: 1 << 10,
	}

	server.ListenAndServe()
}
