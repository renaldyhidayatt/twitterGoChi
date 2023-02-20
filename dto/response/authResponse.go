package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type AuthResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Token   string      `json:"token"`
	Data    interface{} `json:"data"`
}

func ResponseToken(w http.ResponseWriter, message string, token string, data interface{}, status int) {
	res := AuthResponse{
		Status:  status,
		Message: message,
		Token:   token,
		Data:    data,
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}
