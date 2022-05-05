package transactions

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"grillo.com.br/bank-operation/endpoints"
)

func WithdrawMoney(context *gin.Context, path string, rawdata []byte) (float64, error) {
	acc := Account{}
	trs := Transaction{}
	err := json.Unmarshal(rawdata, trs)
	if err != nil {
		return 0, err
	}
	js, err := endpoints.SearchClientData(context, path, rawdata)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal([]byte(js), acc)
	if err != nil {
		return 0, err
	}

	allowWithdraw := trs.Value <= acc.Balance && trs.Value > 0
	if allowWithdraw {
		acc.Balance -= trs.Value
		return acc.Balance, err
	} else {
		return 0, errors.New("Saldo Insuficiente")
	}
}

func CashDeposit(context *gin.Context, path string, rawdata []byte) (float64, error) {
	acc := Account{}
	trs := Transaction{}
	err := json.Unmarshal(rawdata, trs)
	if err != nil {
		return 0, err
	}
	js, err := endpoints.SearchClientData(context, path, rawdata)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal([]byte(js), acc)
	if err != nil {
		return 0, err
	}
	acc.Balance += trs.Value
	return acc.Balance, err
}

func CashBalance(context *gin.Context, path string, rawdata []byte) (float64, error) {
	acc := Account{}
	js, err := endpoints.SearchClientData(context, path, rawdata)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal([]byte(js), acc)
	if err != nil {
		return 0, err
	}
	return acc.Balance, err
}
