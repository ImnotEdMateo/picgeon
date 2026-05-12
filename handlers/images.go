package handlers

import (
	"io"
	"os"
	"time"
	"net/http"
)

var ImagesHandler http.Handler

func init() {
	if dir := os.Getenv("PICGEON_LOCAL_DIR"); dir != "" {
		ImagesHandler = http.StripPrefix("/images/", http.FileServer(http.Dir(dir)))
	} else {
		ImagesHandler = http.HandlerFunc(imageProxyHandler)
	}
}

func imageProxyHandler(w http.ResponseWriter, r *http.Request) {
	baseURL := os.Getenv("PICGEON_URL")
	target := baseURL + r.URL.Path[len("/images/"):]

	resp, err := http.Get(target)
	if err != nil {
		http.Error(w, "error fetching image", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	http.ServeContent(w, r, "", time.Time{}, resp.Body.(io.ReadSeeker))
}
