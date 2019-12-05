package main

import (
	"fmt"
	"net/http"
)

// 05.串联多个处理器
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("called hello()")
	fmt.Fprintf(w, "Hello")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc (func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("called log()")
		fmt.Fprintf(w, "Log...")
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler  {
	return http.HandlerFunc (func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("called protect()")
		fmt.Fprintf(w, "Protect...")
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}

	http.Handle("/hello", protect(log(&HelloHandler{})))

	server.ListenAndServe()
}
