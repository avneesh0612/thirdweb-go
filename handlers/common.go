package handler

import (
	"os"

	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

func getContract() *thirdweb.NFTDrop {
	sdk := initSdk()
	collectionAddress := "0x05B8aab3fd77580C29c6510d8C54D9E6be4262d2"

	contract, _ := sdk.GetNFTDrop(collectionAddress)

	return contract
}

func initSdk() *thirdweb.ThirdwebSDK {
	privateKey := os.Getenv("PRIVATE_KEY")

	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
		PrivateKey: privateKey,
	})
	if err != nil {
		panic(err)
	}

	return sdk
}
