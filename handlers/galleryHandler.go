package handlers

import (
  "html/template"
	"net/http"
	"os"

	"picgeon/utils"
)

func GalleryHandler(w http.ResponseWriter, r *http.Request) {
	var baseURL = os.Getenv("PICGEON_URL")
	if baseURL == "" {
		http.Error(w, "PICGEON_URL not set", http.StatusInternalServerError)
		return
	}

	resp, err := http.Get(baseURL)
	if err != nil {
		http.Error(w, "Error fetching gaallery", 500)
		return
	}
	defer resp.Body.Close()

	media, err := utils.ParseLinks(resp.Body, baseURL)
	if err != nil {
  	http.Error(w, "Can't parse index", 500)
    return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
  tmpl.Execute(w, media)
}
