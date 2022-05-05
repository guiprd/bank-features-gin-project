package endpoints

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"grillo.com.br/bank-operation/accounts"
	"io/ioutil"
)

type Account struct {
	accounts.CheckingAccount
}

func SearchClientData(context *gin.Context, path string, rawdata []byte) (string, error) {
	var fileContent []byte
	var err error
	fileExists, err := accounts.CheckFileExistence(path)
	if fileExists {
		if err != nil {
			return "Error occurred during checking file existence.", err
		}
		fileContent, err = ioutil.ReadFile(path)
		if err != nil {
			return "Error occurred while reading data file.", err
		}
		clientExists, account, err := accounts.CheckClientExistence(fileContent, rawdata)
		if err != nil {
			return "Error occurred during marshal data.", err
		}
		if clientExists {
			js, err := json.Marshal(account)
			if err != nil {
				return "Error occurred during marshal data.", err
			}
			return string(js), err
		}
	}
	return "Client not found in the database", nil
}
