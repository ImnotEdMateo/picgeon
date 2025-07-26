package main

import (
	"log"
	"net/http"
	"os"

	"picgeon/handlers"
)

func main() {
	port := os.Getenv("PICGEON_PORT")
	if port == "" {
		log.Fatal("Port is not defined")
	}

	log.Printf("Running server on  http://0.0.0.0:%s", port)

	http.HandleFunc("/", handlers.GalleryHandler)
	http.Handle("/thumbs/", http.StripPrefix("/thumbs/", http.FileServer(http.Dir("thumbs"))))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
