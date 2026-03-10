package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner string
	balance int
}

var noMoney = errors.New("Can't withdraw you are poor")

// NewAccount account 만드는 함수
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance get balance
func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {

	if a.balance < amount {
		return noMoney
	}

	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}