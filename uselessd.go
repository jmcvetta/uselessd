package main

import (
	"code.google.com/p/go.net/websocket"
	"github.com/jmcvetta/useless"
	"net/http"
)

func main() {
	http.Handle("/useless", websocket.Handler(func(ws *websocket.Conn) {
		ws.Write([]byte(useless.Foobar()))
	}))
	http.ListenAndServe(":3000", nil)
}
