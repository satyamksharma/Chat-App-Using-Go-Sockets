package customwebsocket

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Pool *Pool 
	mu sync.Mutex
}

type Message struct {
	Type int `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		msgType, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		message := Message{Type: msgType, Body: string(msg)}
		c.Pool.Broadcast <- message

		fmt.Printf("Message Received:-----> %+v\n", message)
	}
}

func (c *Client) Write(messageType int, payload []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Conn.WriteMessage(messageType, payload)
}