package main


import (
	"fmt"
	eth "github.com/2cluster/ethclient/client"
	"github.com/ethereum/go-ethereum/common"
)


func main() {

	client, err:= eth.NewClient("Ruud", "f238a37e42b7062bdbc062a1833a6361f9a6d0e324a95ca2f7c4c3034e67ee5c")
	if err != nil {
		fmt.Println(err)
	}

	// client.DeployContract()
	// fmt.Println(client.Contract.Address)


	client.BindContract(common.HexToAddress("0xC382b4aF66EDb6Aa717B6A07330d41364b787B02"))
	balance := client.QueryBalance(client.Account.Address)
	fmt.Println(balance)
	
	allowance := client.QueryAllowance(client.Account.Address,  client.Account.Address)
	fmt.Println(allowance)
	fmt.Println(balance)

	client.AproveAllowance(client.Account.Address,  200)

	client.TransferFrom(client.Account.Address, client.Account.Address,  10)
}