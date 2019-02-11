package main

import (
	"net/http"
	"github.com/gorilla/websocket"
)
	var(
	upgrader = websocket.Upgrader{
		CheckOrigin:func(r *http.Request) bool{
			return true
		},
	}
)

func wsHandler(w http.ResponseWrite, *http.Request){
	var(
		conn *websocket.Conn
		err error
	)

}
func main()  {
	http.HandleFunc("/ws",wsHandler)
	http.ListtenAndServer("0.0.0.0:7777",nil)
}