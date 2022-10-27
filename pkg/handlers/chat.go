package handlers

import (
	"log"
	"net/http"

	"chat/pkg/service"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type ChatHandler struct {
	Upgrader    websocket.Upgrader
	PoolChannel service.PoolChannel
}

func NewChatHandler(poolChannel service.PoolChannel) *ChatHandler {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	return &ChatHandler{
		Upgrader:    upgrader,
		PoolChannel: poolChannel,
	}
}

func (c *ChatHandler) Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := c.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error in websocket conn")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	channel := vars["channel"]

	client := service.Client{
		ClientID: conn.RemoteAddr().String(),
		Conn:     conn,
		Pool:     c.PoolChannel.GetPool(channel),
	}

	log.Println("register client")

	client.Pool.Register <- &client

	client.Read()
}
