package userdto

type UpdateDto struct {
	Username string   `json:"username,omitempty"`
	Password string   `json:"password,omitempty"`
	FullName string   `json:"fullname,omitemptty"`
	Role     string   `json:"role,omitempty"`
	Account  []string `json:"account,omitempty"`
}
