package client


import (
	"fmt"
	"log"
	"crypto/ecdsa"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"

	abi "github.com/2cluster/ethclient/contract"
)

type Client struct {
	Account *Account
	Contract *Contract
	eth *ethclient.Client
}



func newClient(string name, string privateKey) Client {
	acount, err := genAccount(name, privateKey )
	if err != nil {
		log.Fatalln(err)
	}

	connection, err := ethclient.Dial(infuraURL)
	if err 	!= nil {
		log.Fatal(err)
	}

	client := Client{acount, new(Contract), connection}

	contract := client.DeployContract()

	client.Contract = contract

	fmt.Println(client.Account.holder)
	fmt.Println(client.Contract.Name)

	return client
	
}

func (c Client) Test () {
	fmt.Println("hello")
}