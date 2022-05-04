package endpoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"grillo.com.br/bank-operation/contas"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Account struct {
	contas.ContaCorrente
}

func SaveData(path string, data []byte) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panic(err)
	}
	n, err := file.Write(data)
	if err != nil {
		log.Panic(err, n)
	}
	if n, err = file.WriteString("\n"); err != nil {
		fmt.Println(n, err)
	}

	err = file.Close()
	if err != nil {
		log.Panic(err)
	}
}

//Check if the file exists in the file path
func CheckFileExistence(path string) bool {
	if _, err := os.Stat(path); err == nil {
		// path/to/whatever exists
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		return false
	} else {
		// Schrödinger: file may or may not exist. See err for details.
		log.Panic(err)
		return false
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}
}

//Check if the client is registered in the bank
func checkClientExistence(fileContent []byte, rawdata []byte) bool {
	rawDataAccount := Account{}
	json.Unmarshal(rawdata, &rawDataAccount)
	fileSplit := strings.Split(string(fileContent), "\n")
	for _, content := range fileSplit {
		account := Account{}
		json.Unmarshal([]byte(content), &account)
		clientExists := strings.ToLower(account.Titular) == strings.ToLower(rawDataAccount.Titular)
		if clientExists {
			return clientExists
		}
	}
	return false
}

func StructData(context *gin.Context, path string, rawdata []byte) map[string]interface{} {
	var fileContent []byte
	var err error
	account := Account{}
	var data []byte

	err = json.Unmarshal(rawdata, &account)
	if err != nil {
		log.Panic(err)
	}
	newData, err := json.Marshal(account)
	if err != nil {
		log.Panic(err)
	}

	//Verify if a file exists
	fileExists := CheckFileExistence(path)

	if fileExists {
		fileContent, err = ioutil.ReadFile(path)
		if err != nil {
			log.Panic(err)
		}
		clientExists := checkClientExistence(fileContent, rawdata)
		if !clientExists {
			data = append(fileContent[:], newData[:]...)
		} else {
			return gin.H{"message": "Client already registered."}
		}

	} else {
		data = newData
	}

	SaveData(path, data)
	return gin.H{"message": "Client registered successfully."}

}
