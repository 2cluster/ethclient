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

	abi "github.com/2cluster/ethclient/client/contract"
)


const infuraURL		= "https://rinkeby.infura.io/v3/634d2ee71c3e44a4ab4990f90f561398"

type Contract struct {
	Name string
	Instance abi.SUSD
	Address common.Address
	Tx *types.Transaction
}



func (c *Client) DeployContract() *Contract {

	var contract Contract


	nonce, err := c.eth.PendingNonceAt(context.Background(), c.Account.address)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := c.eth.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(c.Account.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice


	address, tx, instance, err := abi.DeploySUSD(auth, c.eth)
	if err != nil {
		log.Fatal(err)
	}

	contract.Name = "SUSD"
	contract.Instance = *instance
	contract.Address = address
	contract.Tx = tx

	log.Println(tx.Hash().Hex())

	return &contract
}

