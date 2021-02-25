package client


import (
	"log"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	abi "github.com/2cluster/ethclient/client/contract"
)


const infuraURL		= "https://rinkeby.infura.io/v3/634d2ee71c3e44a4ab4990f90f561398"

type Contract struct {
	Name string
	Instance abi.SUSD
	Address common.Address
	Tx *types.Transaction
}



func (c *Client) DeployContract() {

	nonce, err := c.eth.PendingNonceAt(context.Background(), c.Account.Address)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := c.eth.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(c.Account.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice


	address, tx, instance, err := abi.DeploySUSD(auth, c.eth)
	if err != nil {
		log.Fatal(err)
	}

	c.Contract.Name = "SUSD"
	c.Contract.Instance = *instance
	c.Contract.Address = address
	c.Contract.Tx = tx

}


func (c *Client) BindContract(address common.Address) {
	instance, err := abi.NewSUSD(address, c.eth)
	if err != nil {
		log.Fatalf("Failed to bind contract: %v", err)
	}

	c.Contract.Name = "SUSD"
	c.Contract.Instance = *instance
	c.Contract.Address = address
}


func (c *Client) QueryBalance(address common.Address) int64 {
	balance, err := c.Contract.Instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatalf("Failed to query balance: %v", err)
	}
	return balance.Int64()
}

func (c *Client) QueryAllowance(from common.Address, spender common.Address) int64 {

	allowance, err := c.Contract.Instance.Allowance(&bind.CallOpts{}, from, spender)
	if err != nil {
		log.Fatalf("Failed to query allowance: %v", err)
	}
	return allowance.Int64()

}
 
func (c *Client) AproveAllowance(spender common.Address, amount int64) {

	nonce, err := c.eth.PendingNonceAt(context.Background(), c.Account.Address)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := c.eth.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(c.Account.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) 
	auth.GasLimit = uint64(3000000) 
	auth.GasPrice = gasPrice

	c.Contract.Instance.Approve(auth, spender, big.NewInt(amount))

}

func (c *Client) Transfer(to common.Address, amount int64) {
	nonce, err := c.eth.PendingNonceAt(context.Background(), c.Account.Address)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := c.eth.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(c.Account.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) 
	auth.GasLimit = uint64(3000000) 
	auth.GasPrice = gasPrice

	c.Contract.Instance.Transfer(auth, to, big.NewInt(amount))

}


func (c *Client) TransferFrom(from common.Address, to common.Address, amount int64) {

	nonce, err := c.eth.PendingNonceAt(context.Background(), c.Account.Address)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := c.eth.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(c.Account.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) 
	auth.GasLimit = uint64(3000000) 
	auth.GasPrice = gasPrice

	c.Contract.Instance.TransferFrom(auth, from, to, big.NewInt(amount))
}
