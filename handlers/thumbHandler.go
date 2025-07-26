package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"picgeon/utils"
)

func ThumbHandler(w http.ResponseWriter, r *http.Request) {
	name := filepath.Base(r.URL.Path) // seguridad básica
	originalName := strings.TrimSuffix(name, ".jpg")
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		http.Error(w, "BASE_URL not set", 500)
		return
	}

	fullURL := baseURL + originalName
	localThumbPath := filepath.Join("thumbs", name)

	if _, err := os.Stat(localThumbPath); os.IsNotExist(err) {
		log.Println("Generando miniatura para", fullURL)

		isVideo := strings.HasSuffix(originalName, ".mp4") || strings.HasSuffix(originalName, ".webm")

		_, genErr := utils.GetOrCreateThumbnail(fullURL, strings.TrimSuffix(name, ".jpg"), isVideo)
		if genErr != nil {
			http.Error(w, "Error generando miniatura", 500)
			return
		}
	}

	http.ServeFile(w, r, localThumbPath)
}
