package main

import (
	customwebsocket "chatapp/websocket"
	"log"
	"net/http"
)

func serveWs(pool *customwebsocket.Pool ,w http.ResponseWriter, r *http.Request) {
	conn, err := customwebsocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	client := &customwebsocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes(){
	log.Println("This is working")
	pool := customwebsocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}


func main() {

	setupRoutes()

	http.ListenAndServe(":8080", nil)
}