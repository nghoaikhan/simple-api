package ioc

import (
	"simple-api/src/app/controllers/usercontrl"
	"simple-api/src/app/repositories/userrepo"
	"simple-api/src/app/schemas/usersche"
	"simple-api/src/app/services/userserv"
)

//Start point of Repositories

//InitUserRepo this function is used to create a user repository instant
func InitUserRepo() *userrepo.Repository {
	return userrepo.NewRepo(usersche.Schema, usersche.Session)
}

//End point of Repositories

//Start point for Services

//InitUserServ is
func InitUserServ() *userserv.Service {

	return userserv.NewUserServ(InitUserRepo())
}

//End point for Services

//Start point for Controllers

func InitUserContrl() *usercontrl.Controller {

	return usercontrl.NewUserContrl(InitUserServ())
}

//End point for Controllers
