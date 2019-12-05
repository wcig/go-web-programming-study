package main

import (
	"fmt"
	"net/http"
)

// 02.使用多个处理器
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "World")
}


func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// Addr: "0.0.0.0:8080",
		Handler: nil,
	}

	helloHandler := &HelloHandler{}
	worldHandler := &WorldHandler{}
	http.Handle("/hello", helloHandler)
	http.Handle("/world", worldHandler)

	// http.Handle("/hello", &HelloHandler{})
	// http.Handle("/world", &WorldHandler{})

	server.ListenAndServe()
}
