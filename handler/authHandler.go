package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/dto/response"
	"github.com/renaldyhidayatt/twittersqlc/security"
	"github.com/renaldyhidayatt/twittersqlc/services"
)

type authHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *authHandler {
	return &authHandler{service: service}
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerRequest request.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&registerRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = registerRequest.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return

	}

	res, err := h.service.RegisterUser(registerRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var authRequest request.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&authRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err = authRequest.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return

	}

	res, err := h.service.Login(authRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if res.UserID > 0 {
		hashPwd := res.Password
		pwd := authRequest.Password

		hash := security.VerifyPassword(hashPwd, pwd)

		fmt.Println(res.Email)

		if hash == nil {
			token, err := security.GenerateToken(res.Email)

			if err != nil {
				response.ResponseError(w, http.StatusInternalServerError, err)
				return
			}

			response.ResponseToken(w, "Login Berhasil", token, res, http.StatusOK)
		} else {
			response.ResponseError(w, http.StatusBadRequest, errors.New("password tidak sesuai"))
			return
		}
	}
}
