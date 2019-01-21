// +build ignore

// This file is for gopherjs build
package main

import (
	"github.com/anikhasibul/chatter/app/router"
	"github.com/gopherjs/vecty"
)

func main() {
	router.Start()
	vecty.AddStylesheet("/app/app.css")
}
