package transacoes

import "grillo.com.br/bank-operation/contas"

type Account struct {
	contas.ContaCorrente
}

func (conta *Account) Saque(saque float64) string {
	permitidoSacar := saque <= conta.Saldo && saque > 0
	if permitidoSacar {
		conta.Saldo -= saque
		return "Saque realizado com sucesso."
	} else {
		return "Saldo insuficiente"
	}
}

func (conta *Account) Depositar(deposito float64) string {
	conta.Saldo += deposito
	return "Depósito realizado com sucesso."
}
