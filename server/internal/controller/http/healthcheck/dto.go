package healthcheck

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func WriteJson(w http.ResponseWriter, code int, data string) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(Response{Message: data})
}
