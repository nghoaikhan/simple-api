package userserv

import (
	"reflect"
	"simple-api/src/app/dtos/userdto"
	"simple-api/src/app/models/usermod"
	"simple-api/src/app/repositories/userrepo"

	"gopkg.in/mgo.v2/bson"
)

// Service is a
type Service struct {
	userRepo *userrepo.Repository
}

func NewUserServ(userRepo *userrepo.Repository) *Service {
	return &Service{userRepo}
}

func (serv *Service) CreateUser(createdDto userdto.CreateDto) (user usermod.Model, err error) {
	defer serv.userRepo.Session.Close()
	userModel := usermod.Model{}
	userModel.InitFromCreateDto(createdDto)

	err = serv.userRepo.Create(&userModel)
	if err == nil {
		user = userModel
	}
	return
}
func (serv *Service) GetUsers() (users []usermod.Model, err error) {
	defer serv.userRepo.Session.Close()

	result, err := serv.userRepo.Find()
	return
}
func CreateArray(t reflect.Type, length int) reflect.Value {
	var arrayType reflect.Type
	arrayType = reflect.ArrayOf(length, t)
	return reflect.Zero(arrayType)
}
func (serv *Service) GetUserByID(ID bson.ObjectId) (user usermod.Model, err error) {
	defer serv.userRepo.Session.Close()

	return

}
func (serv *Service) UpdateUser(dto userdto.UpdateDto) (users []usermod.Model, err error) {

	return
}
func (serv *Service) DeleteUserByID(ID bson.ObjectId) (err error) {
	return
}
