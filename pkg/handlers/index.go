package handlers

import (
	"net/http"
)

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (ih *IndexHandler) HandleChannelA(w http.ResponseWriter, r *http.Request) {
	fname := "html/channelA.html"
	http.ServeFile(w, r, fname)
}

func (ih *IndexHandler) HandleChannelB(w http.ResponseWriter, r *http.Request) {
	fname := "html/channelB.html"
	http.ServeFile(w, r, fname)
}

func (ih *IndexHandler) HandleChannelC(w http.ResponseWriter, r *http.Request) {
	fname := "html/channelC.html"
	http.ServeFile(w, r, fname)
}
