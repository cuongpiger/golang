package api

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "Hello, World!"}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}