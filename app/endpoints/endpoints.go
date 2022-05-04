package endpoints

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"grillo.com.br/bank-operation/contas"
	"log"
)

func DefineEndPoints() {

	path := "contas/contas.txt"

	engine := gin.Default()
	engine.POST("/create-account", func(context *gin.Context) {
		rawData, err := context.GetRawData()
		if err != nil {
			log.Fatal(err)
		}
		msg := StructData(context, path, rawData)
		context.JSON(200, msg)
	})

	engine.GET("/search", func(context *gin.Context) {
		rawData, err := context.GetRawData()
		if err != nil {
			log.Fatal(err)
		}
		conta := contas.ContaCorrente{}
		json.Unmarshal(rawData, &conta)
		//if conta.Titular == "Guilhe"
	})
	err := engine.Run()
	if err != nil {
		log.Fatal(err)
	}
}
