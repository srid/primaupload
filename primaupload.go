package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	ConfigureRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Serving http://%s/\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
