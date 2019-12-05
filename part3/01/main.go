package main

import (
	"fmt"
	"net/http"
)

// 01.开始构建web服务器
func main() {
	// 方式1
	// http.ListenAndServe("127.0.0.1:8080", nil)

	// 方式2
	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// 	Handler: nil,
	// }
	// server.ListenAndServe()

	// 处理请求
	handler := &MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8081",
		Handler: handler,
	}
	server.ListenAndServe()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
}

type MyHandler struct {}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

