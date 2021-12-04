package accounts

//acount struct
type account struct {
	owner   string
	balance int
}

//NewAccount creates account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}
