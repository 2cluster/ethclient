package main


import (
	"fmt"
	// "context"
	eth "github.com/2cluster/ethclient/client"
	"github.com/ethereum/go-ethereum/common"
)


func main() {

	// client, err := eth.NewClient("Ruud", "f238a37e42b7062bdbc062a1833a6361f9a6d0e324a95ca2f7c4c3034e67ee5c", "INFURA")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	client, err := eth.NewClient("NietRuud", "6c0081a5b9511910a6cec018a99d3031197f079cde51c1a78124750a990cdd08", "INFURA")
	if err != nil {
		fmt.Println(err)
	}

	err = client.BindContract(common.HexToAddress("0x5df43ee3333ca8f117f4329de23502df60362f16"))
	if err != nil {
		fmt.Println(err)
	}
	balance, err := client.QueryBalance(client.Account.Address)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balance)
	
	allowance, err := client.QueryAllowance(common.HexToAddress("0x559BC07434C89c5496d790DFD2885dC966F9113a"),  common.HexToAddress("0x54806DD512b21814aa560D627432a75720ed6bB3"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(allowance)

	receipt, err := client.ApproveAllowance(common.HexToAddress("0x54806DD512b21814aa560D627432a75720ed6bB3"),  100000000)
	if err != nil {
		fmt.Println(err)
	}

	receipt, err := client.TransferFrom(common.HexToAddress("0x559BC07434C89c5496d790DFD2885dC966F9113a"), common.HexToAddress("0x6dC89393FA30b64c56DEFF31dAAcf10cEdcD852D"), 100)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(receipt.Status)
	fmt.Println(receipt.TxHash)
	fmt.Println(receipt.BlockHash)
	fmt.Println(receipt.BlockNumber)
	fmt.Println(receipt.TransactionIndex)

}