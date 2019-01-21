package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/anikhasibul/chatter/api/router"
	logger "github.com/anikhasibul/log"
)

var log = logger.New(os.Stdout)

func main() {
	rand.Seed(time.Now().UnixNano())
	var addr = fmt.Sprintf(":%s", os.Getenv("PORT"))
	if os.Getenv("PORT") == "" {
		addr = "127.0.0.1:8080"
	}
	log.Info("Server starting on", addr)
	router.Start()
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
