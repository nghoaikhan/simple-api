package fieldnames

type accounts struct {
	collection string
	AccountID  string
	Role       string
	Status     string
}
type User struct {
	collection string
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

func GetUser() (User User) {
	User.collection = "users"
	User.ID = "_id"
	User.Email = "email"
	User.Password = "password"
	User.FirstName = "firstName"
	User.LastName = "lastName"
	User.FullName = "fullName"
	User.Phone = "phone"
	User.Status = "status"
	User.State = "state"
	User.LastLogin = "lastLogin"
	User.Accounts.Status = "accounts.status"
	User.Accounts.AccountID = "accounts.accountID"
	User.Accounts.Role = "accounts.role"
	User.Created = "created"
	User.CreatedBy = "createdBy"
	User.Modified = "modified"
	User.ModifiedBy = "modifiedBy"
	return User
}
