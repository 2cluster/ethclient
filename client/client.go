package client


import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	Account *Account
	eth *ethclient.Client
	Contract *Contract
}


func NewClient(name string, pk string) (Client, error) {
	acount, err := genAccount(name, pk)
	if err != nil {
		return Client{}, err
	}

	connection, err := ethclient.Dial(infuraURL)
	if err 	!= nil {
		return Client{}, err
	}

	client := Client{acount, connection, new(Contract)}

	return client, nil
}