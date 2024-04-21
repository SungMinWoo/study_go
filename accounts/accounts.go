package accounts

import (
	"errors"
)

type Account struct {
	ower string
	balance int
}

var errNoMoney error = errors.New("Can't withdraw you are poor")

func NewAccount(ower string) *Account{
	account := Account{ower: ower, balance: 0}
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