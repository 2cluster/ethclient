package main


import (
	"fmt"
	eth "github.com/2cluster/ethclient/client"
	"github.com/ethereum/go-ethereum/common"
)


func main() {

	client, err:= eth.NewClient("Ruud", "6c0081a5b9511910a6cec018a99d3031197f079cde51c1a78124750a990cdd08")
	if err != nil {
		fmt.Println(err)
	}

	// client.DeployContract()
	// fmt.Println(client.Contract.Address)


	err = client.BindContract(common.HexToAddress("0x9c6947E8F72228F0D278763EE8efEC4fB7088c29"))
	if err != nil {
		fmt.Println(err)
	}
	balance, err := client.QueryBalance(client.Account.Address)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balance)
	
	// allowance := client.QueryAllowance(client.Account.Address,  common.HexToAddress("0x559BC07434C89c5496d790DFD2885dC966F9113a"))


	err = client.Transfer(common.HexToAddress("0x6dC89393FA30b64c56DEFF31dAAcf10cEdcD852D"),  60)
	if err != nil {
		fmt.Println(err)
	}
	// client.TransferFrom(client.Account.Address, client.Account.Address,  10)
}