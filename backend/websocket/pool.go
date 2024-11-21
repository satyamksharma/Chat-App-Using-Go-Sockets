package customwebsocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for k, _ := range pool.Clients {
				fmt.Println(k)
				k.Conn.WriteJSON(Message{Type: 1, Body: "New User has Joined the chat..."})
			}
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for k, _ := range pool.Clients {
				fmt.Println(k)
				k.Conn.WriteJSON(Message{Type: 1, Body: "User has left the chat..."})
			}
		case msg := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for k, _ := range pool.Clients {
				if err := k.Conn.WriteJSON(msg); err != nil {
					fmt.Println("error: ", err)
					return
				}
			}
		}
	}
}