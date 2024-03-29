package main

import (
	"html/template"
	"net/http"
)

// 01.模板引擎
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmp1.html")
	t.Execute(w, "Hello World!")
}
