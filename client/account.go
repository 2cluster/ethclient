package client


import (
	"fmt"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

)

type Account struct {
	Holder string
	PrivateKey *ecdsa.PrivateKey
	PublicKey *ecdsa.PublicKey
	Address common.Address
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
	
	account.Holder = owner
	account.PrivateKey = privk
	account.PublicKey = publicKeyECDSA
	account.Address = adr


    return &account, nil
}