package main

import (
	"fmt"

	"github.com/akaster99/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("akaster99")
	fmt.Println(account)
	account.Deposiot(100)
	fmt.Println(account.Balance())
}
