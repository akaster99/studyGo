package accounts

import "errors"

//acount struct
type account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("you don't have enough money")

//NewAccount creates account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

func (a *account) Deposiot(amount int) {
	a.balance += amount
}

func (a account) Balance() int {
	return a.balance
}

func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	} else {
		a.balance -= amount
		return nil
	}
}
