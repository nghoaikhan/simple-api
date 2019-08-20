package user

type createDto struct {
	Email string
	Password string
	FullName string
	Role string
	Account []string
}