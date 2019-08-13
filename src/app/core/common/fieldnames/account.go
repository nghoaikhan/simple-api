package fieldnames

type users struct {
	Collection string
	UserId     string
	LastAccess string
}
type account struct {
	Collection    string
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

var Account account

func init() {
	Account = getAccount()
}
func getAccount() (account account) {

	account.Collection = "accounts"
	account.ID = "_id"
	account.AddOnFeatures = "addOnFeatures"
	account.Created = "created"
	account.CreatedBy = "createdBy"
	account.Features = "features"
	account.Modified = "modified"
	account.ModifiedBy = "modifiedBy"
	account.Name = "name"
	account.Status = "status"
	account.Timezone = "timezone"
	account.Users.Collection = "users"
	account.Users.UserId = "userId"
	account.Users.LastAccess = "lastAccess"
	return account

}
