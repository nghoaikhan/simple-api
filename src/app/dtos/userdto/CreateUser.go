package userdto

type CreateDto struct {
	Username string
	Password string
	FullName string
	Role     string
	Account  []string
}
