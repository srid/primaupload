package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		UploadHandler(w, r)
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	yuiUploadKey := "Filedata" // scraped from chrome dev tools
	file, handler, err := r.FormFile(yuiUploadKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: use UUID to avoid conflicts, but store title?
	targetPath := filepath.Join("uploads", handler.Filename)

	target, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer target.Close()

	fmt.Printf("copying to %s\n", targetPath)
	n, err := io.Copy(target, file)
	if err == nil {
		fmt.Printf("copied %d bytes to %s\n", n, targetPath)
	} else {
		fmt.Println(err)
	}
}

func ConfigureRoutes() {
	http.HandleFunc("/", HomeHandler)

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
