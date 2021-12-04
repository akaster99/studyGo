package main

import (
	"fmt"

	"github.com/akaster99/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("akaster99")
	fmt.Println(account)
	account.Deposiot(100)
	error := account.Withdraw(110)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(account.Balance())
	}

}
