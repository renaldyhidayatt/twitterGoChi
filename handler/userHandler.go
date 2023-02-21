package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/twittersqlc/dto/request"
	"github.com/renaldyhidayatt/twittersqlc/dto/response"
	"github.com/renaldyhidayatt/twittersqlc/services"
	"github.com/renaldyhidayatt/twittersqlc/utils"
)

type userHandler struct {
	services services.UserService
}

func NewUserHandler(services services.UserService) *userHandler {
	return &userHandler{services: services}
}

func (h *userHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.services.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *userHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("decoded")
	tokenStr := token.(string)
	var decodedToken map[string]interface{}
	json.Unmarshal([]byte(tokenStr), &decodedToken)
	email, ok := decodedToken["sub"].(string)
	if !ok {
		response.ResponseError(w, http.StatusBadRequest, fmt.Errorf("failed convert"))
	}
	res, err := h.services.GetCurrentUser(email)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	id_int, err := strconv.Atoi(id)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	username_2 := r.FormValue("username2") //untuk ganti
	myupload := utils.UploadFileImage(w, r)
	bio := r.FormValue("bio")
	country := r.FormValue("country")
	website := r.FormValue("website")

	var updateUser request.UpdateUserRequest

	updateUser.Firstname = firstname
	updateUser.Lastname = lastname
	updateUser.Username = username
	updateUser.Email = email
	updateUser.Password = password
	updateUser.Username_2 = username_2
	updateUser.Profileimage = myupload.ProfileImage
	updateUser.Profilecover = myupload.ProfileCover
	updateUser.Bio = bio
	updateUser.Country = country
	updateUser.Website = website

	if err = updateUser.Validate(); err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.services.UpdateUser(id_int, updateUser)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil menupdate data", res, http.StatusOK)
	}

}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	user_id := chi.URLParam(r, "user_id")

	user_int, err := strconv.Atoi(user_id)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.services.DeleteUser(user_int)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendelete data", res, http.StatusOK)
	}

}
