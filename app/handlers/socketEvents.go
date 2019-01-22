package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/anikhasibul/chatter/app/model"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/websocket/websocketjs"
)

type inputModel model.InComing

var Me string

var Socket *websocketjs.WebSocket

func NewSocket(session string) *websocketjs.WebSocket {
	uri := fmt.Sprintf("ws://%s/socket/%s", js.Global.Get("window").Get("location").Get("host").String(), session)
	ws, err := websocketjs.New(uri)
	if err != nil {
		panic(err)
	}
	setUpSock(ws)
	return ws
}

func setUpSock(socket *websocketjs.WebSocket) {
	// EVENT
	onOpen := func(e *js.Object) {
		js.Global.
			Get("window").
			Call("alert", "Connected!")
	}
	// EVENT
	onMessage := func(e *js.Object) {
		// Parse json
		data := receive(e.Get("data").String())
		// send pong to ping messages
		if data.Type == 0 {
			Me = data.Me
			send(inputModel{
				Type: 0,
				Me:   data.Me,
				Msg:  "pong",
			})
			return
		}
		// user 0 is admin
		// so we have to ignore it
		if data.From != "0" {
			// if the user is not self
			if data.From != Me {
				// username
				data.From = data.From[22:29]
			} else {
				data.From = "You"
			}
		} else {
			data.From = "0"
		}
		chatCOMP.Chats = append(chatCOMP.Chats, &model.Message{
			From:    data.From,
			Message: data.Msg,
		})
		vecty.Rerender(chatCOMP)

	}
	// if it closes
	onClose := func(e *js.Object) {
		chatCOMP.Chats = append(chatCOMP.Chats, &model.Message{
			From:    "0",
			Message: "Connection closed! Code: " + e.Get("code").String(),
		})
		vecty.Rerender(chatCOMP)
		session := js.Global.Get("location").Get("pathname").String()[len("/c/"):]
		Socket = NewSocket(session)
	}
	// if error occurs
	onError := func(e *js.Object) {
		js.Global.Get("window").Call("alert", "Connection Failed!")
	}

	// setup events
	socket.AddEventListener("open", false, onOpen)
	socket.AddEventListener("message", false, onMessage)
	socket.AddEventListener("close", false, onClose)
	socket.AddEventListener("error", false, onError)

}

func send(out inputModel) {
	out.Me = Me
	bson, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}
	if Socket != nil {
		err = Socket.Send(
			string(bson),
		)
		if err != nil {
			panic(err)
		}
	}
}

func receive(str string) inputModel {
	var in inputModel
	err := json.Unmarshal([]byte(str), &in)
	if err != nil {
		panic(err)
	}
	return in
}

func submitForm(id string) *vecty.EventListener {
	return event.Submit(
		func(_ *vecty.Event) {
			msg := js.Global.
				Get("document").
				Call("getElementById", id).
				Get("value").
				String()

			send(inputModel{
				Type: 1,
				Msg:  msg,
			})
			js.Global.
				Get("document").
				Call("getElementById", id).
				Set("value", "")
		}).PreventDefault()
}
