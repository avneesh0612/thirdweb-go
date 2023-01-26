package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func ClaimNFT(c *gin.Context) {
	contract := getContract()

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf("%v", data["address"])

	fmt.Println(address)
	fmt.Printf("%T", address)

	tx, err := contract.ClaimTo(context.Background(), address, 1)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"tx": tx,
	})
}
