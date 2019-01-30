package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/anikhasibul/chatter/api/util"
	"github.com/anikhasibul/push"
	"github.com/gorilla/websocket"
)

// nolint: gochecknoglobals
var socket = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type inputModel struct {
	Type int    `json:"type"`
	Me   string `json:"me"`
	Msg  string `json:"msg"`
}

type outputModel struct {
	Type   int    `json:"type"`
	Active int    `json:"active"`
	From   string `json:"from"`
	Me     string `json:"me"`
	Msg    string `json:"msg"`
}

// var db = database.ConnectPSQL(os.Getenv("DATABASE_URL"))

// ChatSocket handles the websocket connection for chat service.
func ChatSocket(w http.ResponseWriter, r *http.Request) {
	// recover from panic
	defer func() {
		if p := recover(); p != nil {
			log.Error("+-PANIC-+", p)
		}
	}()

	// upgrade the HTTP connection
	// to websocket connection
	ws, err := socket.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	// socket closing mechanism
	quit := make(chan bool, 1)
	defer func() {
		// close the connection
		// if remains open
		util.SafeClose(quit)
	}()
	go func() {
		// wait for closing signal
		<-quit
		// close the ws connection
		closeErr := ws.Close()
		if closeErr != nil {
			log.Error(closeErr)
		}
	}()

	// get chat session
	sessID := util.ExtractSession(
		r.URL.Path,
		"/socket/",
	)
	// random ID for each connection
	clID := util.GenerateID()

	// session and client management
	// each registered client
	// on this session/url will
	// receive push messages, when
	// a session pushes a message
	sess := push.NewSession(sessID)
	client := sess.NewClient(clID)
	defer client.DeleteSelf()

	// send a welcome message to
	// notify everyone that
	// a new client has been joined
	// to this url.
	var out = outputModel{
		Type:   0,
		Active: sess.Len(),
		From:   "0",
		Me:     client.KeyString(),
		Msg:    "A new person added.",
	}
	// write the message to the client personally via websocket
	err = ws.WriteJSON(out)
	if err != nil {
		log.Error(err)
		return
	}
	// now push the message to all over the session.
	sess.Push(inputModel{
		Me:  "0",
		Msg: "A new person joined.",
	})
	// receive messages.
	// here we created another
	// goroutine.
	// because, we will send and receive in real time from this single connection.
	// so, one goroutine for send.
	// one for receive.
	// one for ping.
	go func() {
		// defer will run,
		// when the ws conn ended up.
		// so, we've to put all
		// closing signal logic here.
		defer func() {
			// send ws closing signal
			util.SafeClose(quit)
			// close the push
			// message client too.
			client.Close()
		}()
		// read the messages
		for {
			// receive inpur
			var in inputModel
			err = ws.ReadJSON(&in)
			if err != nil {
				// ignore close error message
				if !websocket.IsCloseError(err, 1006, 1001) {
					log.Error(err)
					return
				}
			}
			// Unauthorized user
			if in.Me != client.KeyString() {
				in.Me = "0"
				in.Msg = "A person has been removed!"
				// push the remove message to
				// all active clients on this session
				sess.Push(in)
				return
			}
			// log to database
			/*if in.Type == 1 {
				err = db.InsertLog(
					in.Me,
					in.Msg,
				)
				if err != nil {
					log.Error(err)
					return
				}
			}*/
			// else push the
			// received message
			if in.Type != 0 {
				sess.Push(in)
			}
		}
	}()

	// ping goroutine
	// this sends a ping message to the current client.
	// it ensures that the client is reachable.
	go func() {
		// defer will run,
		// when the ws conn ended up.
		// so, we've to put all
		// closing signal logic here.
		defer func() {
			// send ws closing signal
			util.SafeClose(quit)
			// close the push
			// message client too.
			client.Close()
		}()
		// send ping message
		for {
			// wait 2 seconds
			time.Sleep(2 * time.Second)
			// prepare the message
			var out = outputModel{
				Type:   0,
				Active: sess.Len(),
				From:   "0",
				Me:     client.KeyString(),
				Msg:    "",
			}
			// write the message
			// to websocket.
			// NOTE: look closer, we only pinged the current connection, not the whole session.
			err = ws.WriteJSON(out)
			if err != nil {
				// ignore closed conn error
				if strings.Contains(err.Error(), "use of closed network connection") {
					return
				}
				// ignore closed sent error
				if strings.Contains(err.Error(), "websocket: close sent") {
					return
				}
				// log any else error
				log.Error(err)
				return
			}
		}
	}()

	// send message.
	// we will use this loop for
	// sending a message.
	// our other goroutines are doing the receiving and pinging job.
	// so this loop will only send messages.
	for {
		// wait and pull the input
		// that pushed by our session
		input, err := client.Pull()
		if err != nil {
			// ignore closed client error
			//  ut log other errors.
			if err.Error() != "push: client closed" {
				log.Error(err)
			}
			return
		}

		// assert the input to output model
		msg, ok := input.(inputModel)
		// assertion error :/
		if !ok {
			msg.Me = "0"
			msg.Msg = "Whoops!"
		}
		// prepare the  output
		var out = outputModel{
			Type:   1,
			Active: sess.Len(),
			From:   msg.Me,
			Me:     client.KeyString(),
			Msg:    msg.Msg,
		}
		// write the output
		// to the websocket
		err = ws.WriteJSON(out)
		if err != nil {
			// ignore closed conn error
			if strings.Contains(err.Error(), "use of closed network connection") {
				return
			}
			// ignore closed sent error
			if strings.Contains(err.Error(), "websocket: close sent") {
				return
			}
			// log other error
			log.Error(err)
			return
		}
		// loop ends, wait and pull another message....
	}
}
