package ping

import (
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler"
	"net/http"
)

func (p pongService) GetPong(w http.ResponseWriter, r *http.Request) {
	handler.SetResponse(w, http.StatusOK, string([]byte("pong!")), true, nil, nil)
}
