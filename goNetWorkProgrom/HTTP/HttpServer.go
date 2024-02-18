package http_server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type HttpHandler struct {}

func (h *HttpHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	time.Sleep(4 * time.Second)
	fmt.Fprint(writer, "ping1")
}

func HttpServer() {
	addr := "localhost:8888"
	http.HandleFunc("/ping", func(writer http.ResponseWriter, req *http.Request){
		fmt.Fprint(writer, "ping")
	})

	http.Handle("/ping1", &HttpHandler{})
	
	log.Printf("server is listening on %s", addr)
	http.ListenAndServe(addr, nil)
}