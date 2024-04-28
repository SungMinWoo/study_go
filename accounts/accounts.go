package accounts

import (
	"errors"
)

type Account struct {
	owner string
	balance int
}

var errNoMoney error = errors.New("Can't withdraw you are poor")

func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
}

// Account의 method Account.Deposit()으로 호출 가능, a는 receiver
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// try-catch같은 예외처리가 존재하지 않아 따로 처리해줘야함, error가 없다면 nil을 반환해주면됨
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}


func (a *Account) ChangeOwner(newOwner string){
	a.owner = newOwner
}


func (a Account) Owner() string {
	return a.owner
}

// like python str method
func (a Account) String() string {
	return "do something"
}