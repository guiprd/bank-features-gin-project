package transactions

import (
	"grillo.com.br/bank-operation/accounts"
	"testing"
)

func TestAccount_Depositar(t *testing.T) {
	type fields struct {
		ContaCorrente accounts.CheckingAccount
	}
	type args struct {
		deposito float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := &Account{
				CheckingAccount: tt.fields.ContaCorrente,
			}
			if got := account.CashDeposit(tt.args.deposito); got != tt.want {
				t.Errorf("Depositar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Saque(t *testing.T) {
	type fields struct {
		CheckingAccount accounts.CheckingAccount
	}
	type args struct {
		saque float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conta := &Account{
				CheckingAccount: tt.fields.CheckingAccount,
			}
			if got := conta.WithdrawMoney(tt.args.saque); got != tt.want {
				t.Errorf("Saque() = %v, want %v", got, tt.want)
			}
		})
	}
}
