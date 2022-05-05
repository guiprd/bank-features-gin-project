package transactions

import "grillo.com.br/bank-operation/accounts"

type Account struct {
	accounts.CheckingAccount
}

type Transaction struct {
	Client string
	Value  float64
}
