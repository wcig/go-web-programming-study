package main

import (
	"fmt"
	"net/http"
)

// 03.使用处理器函数
func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello")
}

func world(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "World")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
