package main

import (
	"chat/pkg/handlers"
	"chat/pkg/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	poolChannel := service.NewPoolChannel()
	chatHandler := handlers.NewChatHandler(poolChannel)
	indexHandler := handlers.NewIndexHandler()

	mu := mux.NewRouter()
	mu.HandleFunc("/channelA", indexHandler.HandleChannelA)
	mu.HandleFunc("/channelB", indexHandler.HandleChannelB)
	mu.HandleFunc("/channelC", indexHandler.HandleChannelB)

	mu.HandleFunc("/chat/{channel}", chatHandler.Handle)
	log.Fatal(http.ListenAndServe(":8080", mu))
}
