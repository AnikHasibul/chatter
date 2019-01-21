package router

import (
	"github.com/anikhasibul/chatter/app/handlers"
	"github.com/go-humble/router"
)

var R = router.New()

// Start starts navigating via history API.
func Start() {
	start()
}

func start() {
	R.HandleFunc("/c/{session}", handlers.ChatWindow)
	R.Start()
}
