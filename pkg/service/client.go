package service

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ClientID string
	Conn     *websocket.Conn
	Pool     *Pool
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		c.Pool.Broadcast <- string(p)
	}
}
