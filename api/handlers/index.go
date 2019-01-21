package handlers

import (
	"fmt"
	"net/http"

	"github.com/anikhasibul/chatter/api/util"
)

// Index creates a session and redirects the user to the session.
func Index(w http.ResponseWriter, r *http.Request) {
	// generate a random session uri
	sessionURI := fmt.Sprintf("/c/%s.chat", util.GenerateID())
	// redirect to the session
	http.Redirect(w, r, sessionURI, 301)
}
