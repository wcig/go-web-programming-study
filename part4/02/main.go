package main

import (
	"fmt"
	"net/http"
)

// 02.读取请求体数据
func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	bytes := make([]byte, len)
	r.Body.Read(bytes)
	fmt.Fprintln(w, string(bytes))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
