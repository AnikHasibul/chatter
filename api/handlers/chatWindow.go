package handlers

import (
	"io"
	"net/http"
	"strings"
)

// ChatFrontEnd serves the frontend/html/view page for the chat services.
func ChatFrontEnd(w http.ResponseWriter, r *http.Request) {
	// open the index page
	index := strings.NewReader(`
	<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="initial-scale=1">
	<title>Chatter • v0.1</title>
	<script src="/app/script.js" async defer></script>
</head>
<body>

</body>
</html>
	`)
	// serve the page
	_, err := io.Copy(w, index)
	if err != nil {
		log.Error(err)
	}
}
