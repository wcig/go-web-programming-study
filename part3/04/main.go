package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// 04.串联2个处理器函数 (调用顺序: log -> main)
func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("called hello()")
	fmt.Fprintf(w, "Hello")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler func called:", name)

		fmt.Println("called log()")
		fmt.Fprintf(w, "Log...")
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}

	http.HandleFunc("/hello", log(hello))

	server.ListenAndServe()
}
