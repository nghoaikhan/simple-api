package userserv

type IService interface {
	CreateUser()
	GetUsers()
	GetUserByID()
	UpdateUser()
	DeleteUserByID()
}
