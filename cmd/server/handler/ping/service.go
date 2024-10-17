package ping

import "net/http"

type pongService struct{}

type InterfacePong interface {
	GetPong(w http.ResponseWriter, r *http.Request)
}

func NewHandler() InterfacePong {
	return &pongService{}
}
