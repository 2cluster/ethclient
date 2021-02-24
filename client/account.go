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

)

type Account struct {
	holder string
	privateKey *ecdsa.PrivateKey
	publicKey *ecdsa.PublicKey
	address common.Address
}


func genAccount(owner string, privateKey string) (*Account, error) {

	var account Account
	
	if len(privateKey) != 64 {
		return nil, fmt.Errorf("invalid length of private key")
	}
	privk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("error")
	}

	pubk := privk.Public()
	publicKeyECDSA, ok := pubk.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	adr := crypto.PubkeyToAddress(*publicKeyECDSA)
	
	account.holder = owner
	account.privateKey = privk
	account.publicKey = publicKeyECDSA
	account.address = adr


    return &account, nil
}