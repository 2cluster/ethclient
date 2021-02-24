package main


import (
	"fmt"
	"log"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"contract"
)


const infuraURL		= "https://rinkeby.infura.io/v3/634d2ee71c3e44a4ab4990f90f561398"

type Contract struct {
	name string
	address common.Address
}


type Account struct {
	holder string
	privateKey *ecdsa.PrivateKey
	publicKey *ecdsa.PublicKey
	address common.Address
}

type Client {
	ac *Account
	c *Contract
	eth *ethclient
}


func genAccount(owner string, privateKey string) (Account, error) {

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


    return account, nil
}

func deployContract(caller *Account) Account {

	client, err := ethclient.Dial(infuraURL)
	if err 	!= nil {
		log.Fatal(err)
	}

	ac, err := genAccount(caller.holder, caller.privateKey)
	if err != nil {
		log.Fatalln(err)
	}

	client.account = ac

	nonce, err := client.con.PendingNonceAt(context.Background(), caller.address)
	if err != nil {
		fmt.Errorf(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(caller.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice


	address, tx, instance, err := SUSD.DeploySUSD(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	return instance
}


func main() {
	ac, err := genAccount("Ruud", "f238a37e42b7062bdbc062a1833a6361f9a6d0e324a95ca2f7c4c3034e67ee5c")

	if err != nil {
		log.Fatalln(err)
	}

	client = new(Client{ac})
	fmt.Println(client.ac.holder)
	fmt.Println(client.ac.privateKey)
	fmt.Println(client.ac.publicKey)
	fmt.Println(client.ac.address)
}