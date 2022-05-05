package accounts

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func SaveData(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panic(err)
		return err
	}
	n, err := file.Write(data)
	if err != nil {
		log.Panic(err, n)
		return err
	}
	if _, err = file.WriteString("\n"); err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		log.Panic(err)
		return err
	}
	return err
}

// CheckFileExistence Check if the file exists in the file path
func CheckFileExistence(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		// path/to/whatever exists
		return true, err
	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		return false, err
	} else {
		// Schrödinger: file may or may not exist. See err for details.
		log.Panic(err)
		return false, err
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}
}

//Check if the client is registered in the bank
func CheckClientExistence(fileContent []byte, rawdata []byte) (bool, CheckingAccount, error) {
	rawDataAccount := CheckingAccount{}
	err := json.Unmarshal(rawdata, &rawDataAccount)
	fileSplit := strings.Split(string(fileContent), "\n")
	for _, content := range fileSplit {
		account := CheckingAccount{}
		err := json.Unmarshal([]byte(content), &account)
		if err != nil {
			return false, CheckingAccount{}, err
		}
		clientExists := strings.ToLower(account.Holder) == strings.ToLower(rawDataAccount.Holder)
		if clientExists {
			return clientExists, account, err
		}
	}
	return false, CheckingAccount{}, err
}

func CreateAccount(context *gin.Context, path string, rawdata []byte) (string, error) {
	var fileContent []byte
	var err error
	account := CheckingAccount{}
	var data []byte

	err = json.Unmarshal(rawdata, &account)
	if err != nil {
		log.Panic(err)
		return "Error occurred during unmarshal rawdata.", err
	}
	newData, err := json.Marshal(account)
	if err != nil {
		log.Panic(err)
		return "Error occurred during marshal data.", err
	}

	//Verify if a file exists
	fileExists, err := CheckFileExistence(path)
	if fileExists {
		if err != nil {
			return "Error occurred during checking file existence.", err
		}
		fileContent, err = ioutil.ReadFile(path)
		if err != nil {
			log.Panic(err)
			return "Error occurred while reading file.", err
		}
		clientExists, _, err := CheckClientExistence(fileContent, rawdata)
		if err != nil {
			return "Error occurred during validation client data.", err
		}
		if !clientExists {
			data = append(fileContent[:], newData[:]...)
		} else {
			return "Client already registered.", err
		}

	} else {
		data = newData
	}

	err = SaveData(path, data)
	if err != nil {
		return "Error occurred while saving data.", err
	}
	return "Client registered successfully.", err
}
