package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func nigga(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Vos sos un prbo hpta")
}

func main() {
	port := os.Getenv("PICGEON_PORT")
	if port == "" {
		log.Fatal("Port is not defined")
	}

	log.Printf("Running server on  http://0.0.0.0:%s", port)

	http.HandleFunc("/", nigga)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
