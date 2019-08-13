package fieldnames

type accounts struct {
	Collection string
	AccountID  string
	Role       string
	Status     string
}
type user struct {
	Collection string
	ID         string
	Email      string
	Password   string
	FirstName  string
	LastName   string
	FullName   string
	Phone      string
	Status     string
	State      string
	LastLogin  string
	Accounts   accounts
	Created    string
	CreatedBy  string
	Modified   string
	ModifiedBy string
}

var User user

func init() {
	User = getUser()
}
func getUser() (user user) {
	user.Collection = "users"
	user.ID = "_id"
	user.Email = "email"
	user.Password = "password"
	user.FirstName = "firstName"
	user.LastName = "lastName"
	user.FullName = "fullName"
	user.Phone = "phone"
	user.Status = "status"
	user.State = "state"
	user.LastLogin = "lastLogin"
	user.Accounts.Status = "accounts.status"
	user.Accounts.AccountID = "accounts.accountID"
	user.Accounts.Role = "accounts.role"
	user.Created = "created"
	user.CreatedBy = "createdBy"
	user.Modified = "modified"
	user.ModifiedBy = "modifiedBy"
	return user
}
