package main

import (
	"encoding/json"
	"net/http"
)

// 04.响应内容
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeHeader", writeHeaderExample)
	http.HandleFunc("/redirect", redirectExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &struct {
		User    string
		Threads []string
	}{
		User:    "San Sheong",
		Threads: []string{"first", "second", "third"},
	}
	bytes, _ := json.Marshal(post)
	w.Write(bytes)
}

func redirectExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://baidu.com")
	w.WriteHeader(302)
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	w.Write([]byte("No such service, try next door"))
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
		<head><title>Go Web Programming</title></head>
		<body>Hello World.</body>
		</html>`
	w.Write([]byte(str))
}
