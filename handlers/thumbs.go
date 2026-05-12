package handlers

import "net/http"

var ThumbsHandler = http.StripPrefix("/thumbs/", http.FileServer(http.Dir("thumbs")))
