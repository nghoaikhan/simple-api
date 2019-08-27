package userserv

import (
	"simple-api/src/app/dtos/userdto"
	"simple-api/src/app/models/usermod"
	"simple-api/src/app/repositories/userrepo"
)

// Service is a
type Service struct {
	userRepo *userrepo.Repository
}

func NewUserServ(userRepo *userrepo.Repository) *Service {
	return &Service{userRepo}
}

func (serv *Service) CreateUser(createdDto userdto.CreateDto) (user usermod.Model, err error) {

	userModel := usermod.Model{}
	userModel.InitFromCreateDto(createdDto)

	err = serv.userRepo.Create(&userModel)
	if err == nil {
		user = userModel
	}
	return
}
func (serv *Service) GetUsers() (users []usermod.Model, err error) {

	result, err := serv.userRepo.Find()
	users = usermod.Model{}.ConvertFromArrInter(result)

	return
}

func (serv *Service) GetUserByID(id string) (user usermod.Model, err error) {

	result, err := serv.userRepo.FindByID(id)
	user.InitFromInter(result)
	return

}
func (serv *Service) UpdateUserByID(id string, dto userdto.UpdateDto) (user usermod.Model, err error) {
	bsonUpdate, err := dto.ToBsonMap()
	if err != nil {
		return
	}
	err = serv.userRepo.UpdateByID(id, bsonUpdate)
	if err != nil {
		return
	}
	user, err = serv.GetUserByID(id)
	return
}
func (serv *Service) DeleteUserByID(id string) (err error) {
	err = serv.userRepo.DeleteByID(id)
	return
}

func (serv *Service) CloseDBSession() {
	serv.userRepo.Session.Close()
}
