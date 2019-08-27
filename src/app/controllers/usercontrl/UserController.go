package usercontrl

import (
	"net/http"
	"simple-api/src/app/core/utils"
	"simple-api/src/app/dtos/userdto"
	"simple-api/src/app/services/userserv"

	"github.com/gorilla/mux"
)

type Controller struct {
	userServ *userserv.Service
}

func NewUserContrl(userServ *userserv.Service) *Controller {
	return &Controller{userServ}
}

func (contrl *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto userdto.CreateDto

	// decode body request
	err := utils.DecodeJson(r.Body, &dto)
	if err != nil {
		//handle
		return
	}
	// call service to create a user
	user, err := contrl.userServ.CreateUser(dto)
	if err != nil {
		//handle
		return
	}

	// endcode to json to send respone to client
	json, err := utils.EndcodeJson(&user)
	if err != nil {
		//handle
		return
	}

	// send result to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)

}

func (contrl *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	// call service to get all users
	users, err := contrl.userServ.GetUsers()
	if err != nil {
		//handle
		return
	}

	// endcode to json to send respone to client
	json, err := utils.EndcodeJson(&users)
	if err != nil {
		//handle
		return
	}

	// send result to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}

func (contrl *Controller) GetUserByID(w http.ResponseWriter, r *http.Request) {

	// Get Id from params
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// call service to get a user by Id
	user, err := contrl.userServ.GetUserByID(id)
	if err != nil {
		//handle
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// endcode to json to send respone to client
	json, err := utils.EndcodeJson(&user)
	if err != nil {
		//handle
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// send result to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}
func (contrl *Controller) UpdateUserByID(w http.ResponseWriter, r *http.Request) {

	// Get Id from params
	vars := mux.Vars(r)
	id := vars["id"]

	// Get update Object from request
	var dto userdto.UpdateDto

	// decode body request
	err := utils.DecodeJson(r.Body, &dto)
	if err != nil {
		//handle
		return
	}

	// call service to update a user by Id
	user, err := contrl.userServ.UpdateUserByID(id, dto)
	if err != nil {
		//handle
		return
	}

	// endcode to json to send respone to client
	json, err := utils.EndcodeJson(&user)
	if err != nil {
		//handle
		return
	}

	// send result to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}

func (contrl *Controller) DeleteByID(w http.ResponseWriter, r *http.Request) {

	// Get Id from params
	vars := mux.Vars(r)
	id := vars["id"]

	// call service to delete a user by Id
	err := contrl.userServ.DeleteUserByID(id)
	if err != nil {
		//handle
		return
	}

	// send result to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
