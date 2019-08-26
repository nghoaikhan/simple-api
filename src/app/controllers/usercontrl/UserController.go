package usercontrl

import (
	"net/http"
	"simple-api/src/app/core/utils"
	"simple-api/src/app/dtos/userdto"
	"simple-api/src/app/models/usermod"
	"simple-api/src/app/services/userserv"
)

type Controller struct {
	userServ *userserv.Service
}

func NewUserContrl(userServ *userserv.Service) *Controller {
	return &Controller{userServ}
}

func (contrl *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto userdto.CreateDto
	var user usermod.Model
	err := utils.DecodeJson(r.Body, &dto)
	if err != nil {
		//handle
		return
	}
	user, err = contrl.userServ.CreateUser(dto)
	if err != nil {
		//handle
		return
	}
	json, err := utils.EndcodeJson(&user)
	if err != nil {
		//handle
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)

}
