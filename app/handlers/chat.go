package handlers

import (
	"github.com/anikhasibul/chatter/app/components"
	"github.com/go-humble/router"
	"github.com/gopherjs/vecty"
)

var chatCOMP = &components.ChatWindow{}

//
func ChatWindow(p *router.Context) {
	// get the session
	session := p.Params["session"]
	// connect to websocket
	Socket = NewSocket(session)
	// render the html
	chatCOMP.Session = session
	chatCOMP.EvFunc = submitForm
	vecty.RenderBody(chatCOMP)
}
