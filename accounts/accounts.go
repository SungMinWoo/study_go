package accounts

type Account struct {
	ower string
	balance int
}

func NewAccount(ower string) *Account{
	account := Account{ower: ower, balance: 0}
	return &account
}