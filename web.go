package main

import (
	"code.google.com/p/gorilla/mux"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// maps filename to description
// ideally, description should be stored in a database
var descriptionMap = map[string]string{}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request from", r)
	if r.Method == "POST" {
		saveHandler(w, r)
		return
	}
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("Filedata")
	if err != nil {
		serverError(w, err)
		return
	}

	// use a UUID in the filename to avoid conflicts
	targetPath := filepath.Join("static", "uploads", appendUuidToFilepath(handler.Filename))

	target, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		serverError(w, err)
		return
	}
	defer target.Close()

	log.Printf("copying to %s\n", targetPath)
	n, err := io.Copy(target, file)
	if err == nil {
		log.Printf("copied %d bytes to %s\n", n, targetPath)
	} else {
		serverError(w, err)
		return
	}

	fmt.Fprintf(w, "/"+targetPath)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Base(r.FormValue("savedfile"))
	description := r.FormValue("description")
	if path == "" {
		log.Println("error: request did not contain path to uploaded file; submitted before upload completion?")
		return
	}
	descriptionMap[path] = description
	http.Redirect(w, r, "/view/"+path, http.StatusFound)
}

func fileInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	if desc, ok := descriptionMap[filename]; ok {
		t := template.Must(template.ParseFiles("view.html"))
		t.Execute(w, map[string]string{
			"Title":       removeUuidFromFilepath(filename),
			"Path":        fmt.Sprintf("/static/uploads/%s", filename),
			"Description": desc})
	} else {
		http.NotFound(w, r)
	}
}

func flashCrossDomainFile(w http.ResponseWriter, r *http.Request) {
	content := `<cross-domain-policy>
<allow-access-from domain="yui.yahooapis.com" secure="false"/>
</cross-domain-policy>
`
	fmt.Fprintf(w, content)
}

func configureRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/upload", uploadHandler)
	router.HandleFunc("/view/{filename}", fileInfoHandler)
	router.HandleFunc("/crossdomain.xml", flashCrossDomainFile)

	// static directory handler
	staticDir, err := filepath.Abs("./static")
	if err != nil {
		panic(err)
	}
	staticHandler := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static", staticHandler))
	http.Handle("/", router)
}

func appendUuidToFilepath(path string) string {
	id, _ := uuid.NewV4()
	return fmt.Sprintf("%s-%s", id, path)
}

func removeUuidFromFilepath(path string) string {
	return strings.SplitN(filepath.Base(path), "-", 6)[5]
}

func serverError(w http.ResponseWriter, err error) {
	log.Printf("error: %s\n", err)
	http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
}

func main() {
	configureRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Serving http://%s/\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
