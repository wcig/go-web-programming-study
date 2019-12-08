package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 03.Form字段
func process(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// fmt.Println("r.Form:", r.Form) // 请求体+URL键值对
	// fmt.Println("r.PostForm:", r.PostForm) // 请求体键值对
	// fmt.Fprintln(w, r.Form)

	// r.ParseMultipartForm(1024)
	// fmt.Println("r.MultipartForm:", r.MultipartForm) // multipart/form-data
	// fmt.Fprintln(w, r.MultipartForm)

	// r.ParseMultipartForm(1024)
	// fileHeader := r.MultipartForm.File["uploaded"][0]
	// file, err := fileHeader.Open()
	// if err == nil {
	// 	data, err := ioutil.ReadAll(file)
	// 	if err == nil {
	// 		fmt.Fprintln(w, string(data))
	// 	}
	// }

	if file, _, err := r.FormFile("uploaded"); err == nil {
		if data, err := ioutil.ReadAll(file); err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
