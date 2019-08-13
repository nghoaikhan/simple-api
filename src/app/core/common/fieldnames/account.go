package fieldnames

type users struct {
	collection string
	UserId     string
	LastAccess string
}
type Account struct {
	collection    string
	ID            string
	Name          string
	Status        string
	Timezone      string
	Features      string
	AddOnFeatures string
	Created       string
	CreatedBy     string
	Modified      string
	ModifiedBy    string
	Users         users
}

func GetAccount() (Account Account) {

	Account.collection = "accounts"
	Account.ID = "_id"
	Account.AddOnFeatures = "addOnFeatures"
	Account.Created = "created"
	Account.CreatedBy = "createdBy"
	Account.Features = "features"
	Account.Modified = "modified"
	Account.ModifiedBy = "modifiedBy"
	Account.Name = "name"
	Account.Status = "status"
	Account.Timezone = "timezone"
	Account.Users.collection = "users"
	Account.Users.UserId = "userId"
	Account.Users.LastAccess = "lastAccess"
	return Account

}
