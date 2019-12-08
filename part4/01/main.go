package main

import (
	"fmt"
	"net/http"
)

// 01.获取请求头参数
func headers(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	acceptHeader := r.Header["Accept-Encoding"]
	fmt.Println("acceptHeader:", acceptHeader)

	acceptHeader2 := r.Header.Get("Accept-Encoding")
	fmt.Println("acceptHeader2:", acceptHeader2)
	fmt.Fprintln(w, header)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
