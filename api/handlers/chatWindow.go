package handlers

import (
	"io"
	"net/http"
	"os"
)

// ChatFrontEnd serves the frontend/html/view page for the chat services.
func ChatFrontEnd(w http.ResponseWriter, r *http.Request) {
	// open the index page
	f, err := os.Open("index.html")
	if err != nil {
		log.Error(err)
	}
	// serve the page
	_, err = io.Copy(w, f)
	if err != nil {
		log.Error(err)
	}
}
