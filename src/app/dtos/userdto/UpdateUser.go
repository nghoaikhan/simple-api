package userdto

type UpdateDto struct {
	Username string
	Password string
	FullName string
	Role     string
	Account  []string
}
