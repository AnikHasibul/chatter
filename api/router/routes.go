package router

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/anikhasibul/chatter/api/handlers"
)

func Start() {
	routes()
}

func routes() {
	// index page
	http.HandleFunc("/", handlers.Index)
	// chat front end
	http.HandleFunc("/c/", handlers.ChatFrontEnd)
	// chat backend / websocket
	http.HandleFunc("/socket/", handlers.ChatSocket)
	// static files
	http.Handle("/app/",
		gziphandler.GzipHandler(http.StripPrefix("/app/", http.FileServer(http.Dir("static")))),
	)
}
