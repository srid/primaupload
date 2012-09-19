package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func appendUuidToFilepath(path string) string {
	id, _ := uuid.NewV4()
	return fmt.Sprintf("%s-%s", id, path)
}

func removeUuidFromFilepath(path string) string {
	return strings.SplitN(filepath.Base(path), "-", 6)[5]
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		UploadHandler(w, r)
		return
	}
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	yuiUploadKey := "Filedata" // scraped from chrome dev tools
	file, handler, err := r.FormFile(yuiUploadKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	// use a UUID in the filename to avoid conflicts
	targetPath := filepath.Join("static", "uploads", appendUuidToFilepath(handler.Filename))

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

	fmt.Fprintf(w, "/"+targetPath)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("savedfile")
	description := r.FormValue("description")
	t := template.Must(template.ParseFiles("view.html"))
	t.Execute(w, map[string]string{
		"Title":       removeUuidFromFilepath(path),
		"Path":        path,
		"Description": description})
}

func ConfigureRoutes() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/save", SaveHandler)

	// static directory handler
	staticDir, err := filepath.Abs("./static")
	if err != nil {
		panic(err)
	}
	staticHandler := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static", staticHandler))
}

func main() {
	ConfigureRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	} 
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Printf("Serving http://%s/\n", addr)
	http.ListenAndServe(addr, nil)
}
