package transacoes

import (
	"grillo.com.br/bank-operation/contas"
	"testing"
)

func TestAccount_Depositar(t *testing.T) {
	type fields struct {
		ContaCorrente contas.ContaCorrente
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
			conta := &Account{
				ContaCorrente: tt.fields.ContaCorrente,
			}
			if got := conta.Depositar(tt.args.deposito); got != tt.want {
				t.Errorf("Depositar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Saque(t *testing.T) {
	type fields struct {
		ContaCorrente contas.ContaCorrente
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
				ContaCorrente: tt.fields.ContaCorrente,
			}
			if got := conta.Saque(tt.args.saque); got != tt.want {
				t.Errorf("Saque() = %v, want %v", got, tt.want)
			}
		})
	}
}
