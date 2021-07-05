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

	client, err := eth.NewClient("dealblock", "f238a37e42b7062bdbc062a1833a6361f9a6d0e324a95ca2f7c4c3034e67ee5c", "INFURA")
	if err != nil {
		fmt.Println(err)
	}

	err = client.BindContract(common.HexToAddress("0x9c6947e8f72228f0d278763ee8efec4fb7088c29"))
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

	// receipt, err := client.ApproveAllowance(common.HexToAddress("0x54806DD512b21814aa560D627432a75720ed6bB3"),  100000000)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	receipt, err := client.Transfer(common.HexToAddress("0x54806DD512b21814aa560D627432a75720ed6bB3"), 100)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(receipt.Status)
	fmt.Printf("%T\n", receipt.Status)


	fmt.Println(receipt.ContractAddress.Hex())
	fmt.Printf("%T\n", receipt.ContractAddress.String())

	fmt.Println(receipt.TxHash.Hex())
	fmt.Println(receipt.BlockHash.Hex())


	fmt.Println(receipt.BlockNumber)
	fmt.Printf("%T\n", receipt.BlockNumber)
	fmt.Println(receipt.TransactionIndex)
	fmt.Printf("%T\n", receipt.TransactionIndex)

}