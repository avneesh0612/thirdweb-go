package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

type Body struct {
	Address string `json:"address"`
}

func GenerateSignature(c *gin.Context) {
	// Get contract
	contract := getContract()

	// Get address
	var address Body
	c.BindJSON(&address)

	println(address.Address)

	payload := &thirdweb.Signature721PayloadInput{
		To: "0x39Ab29fAfb5ad19e96CFB1E1c492083492DB89d4",
		Metadata: &thirdweb.NFTMetadataInput{
			Name:        "My NFT",
			Description: "This is my cool NFT",
			Image:       "https://gateway.ipfscdn.io/ipfs/QmXXjx3aJCs9W9mN35Aade6etSoceqMk8ykkasbB87MaLt/0.png",
		},
	}

	signedPayload, err := contract.Signature.Generate(context.Background(), payload)

	// Get all NFTs
	if err != nil {
		panic(err)
	}

	println(signedPayload)

	c.JSON(200, gin.H{
		"signedPayload": signedPayload,
	})
}
