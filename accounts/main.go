package main

import (
	"fmt"

	"github.com/hyunwoomemo/learngo/accounts"
)

func main() {

	account := accounts.NewAccount("hyunwoo")
	account.Deposit(10)
	err := account.Withdraw(20)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account.Balance())
}