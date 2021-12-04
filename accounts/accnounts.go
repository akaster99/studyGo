package accounts

import (
	"errors"
	"fmt"
)

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
func (a account) String() string {
	return fmt.Sprint(a.Owner(), "s account.\nHas: ", a.Balance())
}
func (a account) Owner() string {
	return a.owner
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

func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}
