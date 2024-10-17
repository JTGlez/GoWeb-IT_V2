package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrorInvalidID = errors.New("invalid id input")
)

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Count   *int        `json:"count,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SetResponse(w http.ResponseWriter,
	responseCode int,
	data interface{},
	success bool,
	err error,
	count *int) {

	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	r := Response{
		Success: success,
		Error:   errMsg,
		Count:   count,
		Data:    data,
	}

	dataBytes, jsonErr := json.Marshal(r)
	if jsonErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"success":false,"error":"internal server error!"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	_, _ = w.Write(dataBytes)
}
