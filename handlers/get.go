package handler

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetNFTs(c *gin.Context) {
	// Get contract
	contract := getContract()

	// Get all NFTs
	nfts, err := contract.ERC721.GetAll(context.Background())
	if err != nil {
		panic(err)
	}

	// Format NFTs
	formattedNfts, _ := json.Marshal(nfts)
	var nftsArray []interface{}
	json.Unmarshal(formattedNfts, &nftsArray)

	// Return NFTs
	c.JSON(200, gin.H{
		"nfts": nftsArray,
	})
}
