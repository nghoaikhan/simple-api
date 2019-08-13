package fieldnames

type role struct {
	Collection string
	ID         string
	Account    string
	Name       string
	SystemRole string
	Created    string
	CreatedBy  string
	Modified   string
	ModifiedBy string
}

func GetRole() (role role) {

	role.Collection = "roles"
	role.ID = "_id"
	role.Account = "account"
	role.Name = "name"
	role.SystemRole = "systemRole"
	role.Created = "created"
	role.CreatedBy = "createdBy"
	role.Modified = "modified"
	role.ModifiedBy = "modifiedBy"
	return role
}
