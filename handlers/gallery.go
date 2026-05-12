package handlers

import (
	"encoding/json"
	"net/http"
	"picgeon/store"
)

func GalleryHandler(w http.ResponseWriter, r *http.Request) {
	items, err := store.Default.Get()
	if err != nil {
		http.Error(w, "error indexing gallery", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(items)
}
