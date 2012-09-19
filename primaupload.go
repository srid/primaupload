package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: don't laod into memory
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: use UUID to avoid conflicts, but store title?
	targetPath := filepath.Join("uploads", handler.Filename)
	err = ioutil.WriteFile(targetPath, data, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func ConfigureRoutes() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/upload", UploadHandler)

	// static directory handler
	staticDir, err := filepath.Abs("./static")
	if err != nil {
		panic(err)
	}
	staticHandler := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static", staticHandler))
}

func main() {
	fmt.Println("Serving http://localhost:8080/")
	ConfigureRoutes()
	http.ListenAndServe(":8080", nil)
}
