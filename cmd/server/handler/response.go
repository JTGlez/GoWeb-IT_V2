package handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Error   error       `json:"error"`
	Count   *int        `json:"count,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SetResponse(w http.ResponseWriter,
	responseCode int,
	data interface{},
	success bool,
	err error,
	count *int) {

	r := Response{
		Success: success,
		Error:   err,
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
