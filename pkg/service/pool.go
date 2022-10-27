package service

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Pool struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan string
}

func newPool() *Pool {
	return &Pool{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan string),
	}
}

func (p *Pool) run() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			msg := fmt.Sprintf("%s has joined", client.ClientID)
			p.broadcastMessage(msg)
		case client := <-p.Unregister:
			delete(p.Clients, client)
			msg := fmt.Sprintf("%s has left", client.ClientID)
			p.broadcastMessage(msg)
		case msg := <-p.Broadcast:
			p.broadcastMessage(msg)
		}
	}
}

func (p *Pool) broadcastMessage(msg string) {
	for client := range p.Clients {
		client.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
	}
}

type PoolChannel map[string]*Pool

const defaultChannel = "general"

func NewPoolChannel() PoolChannel {
	once.Do(func() {
		pc = make(map[string]*Pool)
	})
	return pc
}

func (pc PoolChannel) GetPool(channelName ...string) *Pool {
	channel := defaultChannel
	if len(channelName) > 0 {
		channel = channelName[0]
	}
	if p, ok := pc[channel]; ok {
		return p
	}
	pool := newPool()
	pc[channel] = pool
	go pool.run()
	return pool
}

var (
	pc   map[string]*Pool
	once sync.Once
)
