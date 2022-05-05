package endpoints

import (
	"github.com/gin-gonic/gin"
	"grillo.com.br/bank-operation/accounts"
	"grillo.com.br/bank-operation/accounts/transactions"
	"log"
)

func DefineEndPoints() {

	path := "accounts/accounts.txt"

	engine := gin.Default()
	engine.POST("/create-account", func(context *gin.Context) {
		rawData, err := context.GetRawData()
		if err != nil {
			log.Fatal(err)
			context.JSON(400, gin.H{"message": "Error occurred while reading raw data.", "Error": err})
		}
		msg, err := accounts.CreateAccount(context, path, rawData)
		if err != nil {

		} else {
			context.JSON(200, gin.H{"message": msg})
		}

	})

	engine.POST("transaction/withdraw-money", func(context *gin.Context) {
		rawData, err := context.GetRawData()
		if err != nil {
			log.Fatal(err)
			context.JSON(400, gin.H{"message": "Error occurred while reading raw data.", "Error": err})
		}
		balance, err := transactions.WithdrawMoney(context, path, rawData)
		context.JSON(200, gin.H{"balance": balance})
	})

	engine.POST("transaction/cash-deposit", func(context *gin.Context) {
		rawData, err := context.GetRawData()
		if err != nil {
			log.Fatal(err)
			context.JSON(400, gin.H{"message": "Error occurred while reading raw data.", "Error": err})
		}
		balance, err := transactions.CashDeposit(context, path, rawData)
		if err != nil {
			log.Fatal(err)
			context.JSON(400, gin.H{"message": "Error occurred while doing deposit transaction.", "Error": err})
		}
		context.JSON(200, gin.H{"balance": balance})
	})

	engine.GET("/balance", func(context *gin.Context) {
		rawData, err := context.GetRawData()
		if err != nil {
			log.Panic(err)
			context.JSON(400, gin.H{"message": "Error occurred while reading raw data.", "Error": err})
		}
		transactions.CashBalance(context, path, rawData)

	})
	engine.GET("/search-client", func(context *gin.Context) {
		rawData, err := context.GetRawData()
		if err != nil {
			log.Panic(err)
			context.JSON(400, gin.H{"message": "Error occurred while reading raw data.", "Error": err})
		}
		msg, err := SearchClientData(context, path, rawData)
		if err != nil {
			log.Panic(err)
			context.JSON(400, gin.H{"message": "Error occurred while searching client.", "Error": err})
		}
		context.JSON(200, gin.H{"message": msg, "Error": err})
	})
	err := engine.Run()
	if err != nil {
		panic(err.Error())
	}
}
