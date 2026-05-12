package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"picgeon/handlers"
	"picgeon/utils/scanning"
	"picgeon/store"
)

func main() {
	port := os.Getenv("PICGEON_PORT")
	if port == "" {
		log.Fatal("PICGEON_PORT not set")
	}

	// inicializar scanner según entorno
	var scanner scanning.Scanner
	if dir := os.Getenv("PICGEON_LOCAL_DIR"); dir != "" {
		scanner = &scanning.LocalScanner{
			Dir:     dir,
			BaseURL: os.Getenv("PICGEON_URL"),
		}
	} else {
		scanner = &scanning.HTTPScanner{
			BaseURL: os.Getenv("PICGEON_URL"),
		}
	}

	store.Default = store.NewStore(scanner, 5*time.Minute)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/gallery", handlers.GalleryHandler)
	mux.Handle("GET /thumbs/", handlers.ThumbsHandler)
	mux.Handle("GET /images/", handlers.ImagesHandler)

	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
