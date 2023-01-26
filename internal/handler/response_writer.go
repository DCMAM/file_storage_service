package handler

import (
	"encoding/json"
	"net/http"
)

// Response will write http response
func Response(rw http.ResponseWriter, data interface{}, code int) {
	apiResponseData := struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}{
		code,
		data,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)

	json.NewEncoder(rw).Encode(apiResponseData)

	return
}
