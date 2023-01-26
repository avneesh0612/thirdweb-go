package handler

import (
	"os"

	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

func getContract() *thirdweb.NFTCollection {
	sdk := initSdk()
	collectionAddress := "0xc9cD8FB4D61204171af7dE5B7aB7aB7c948597CB"

	contract, _ := sdk.GetNFTCollection(collectionAddress)

	return contract
}

func initSdk() *thirdweb.ThirdwebSDK {
	privateKey := os.Getenv("PRIVATE_KEY")

	sdk, err := thirdweb.NewThirdwebSDK("goerli", &thirdweb.SDKOptions{
		PrivateKey: privateKey,
	})
	if err != nil {
		panic(err)
	}

	return sdk
}
