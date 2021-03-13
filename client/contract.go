package client


import (
	"fmt"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	abi "github.com/2cluster/ethclient/client/contract"
)

var VALUE_URL = map[string]string{
	"INFURA"		: "https://rinkeby.infura.io/v3/634d2ee71c3e44a4ab4990f90f561398",
	"LOCAL" 		: "http://127.0.0.1:8545/",
}

type Contract struct {
	Name string
	Instance abi.SUSD
	Address common.Address
	Tx *types.Transaction
}


func (c *Client) DeployContract() error {
	ctx := context.Background()
	nonce, err := c.eth.PendingNonceAt(ctx, c.Account.Address)
	if err != nil {
		return err
	}

	gasPrice, err := c.eth.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(c.Account.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	address, tx, instance, err := abi.DeploySUSD(auth, c.eth)
	if err != nil {
		return err
	}

	c.Contract.Name = "SUSD"
	c.Contract.Instance = *instance
	c.Contract.Address = address
	c.Contract.Tx = tx

	return nil
}


func (c *Client) BindContract(address common.Address) error {
	instance, err := abi.NewSUSD(address, c.eth)
	if err != nil {
		return fmt.Errorf("Failed to bind contract: %v", err)
	}

	c.Contract.Name = "SUSD"
	c.Contract.Instance = *instance
	c.Contract.Address = address

	return nil
}


func (c *Client) QueryBalance(address common.Address) (int64, error) {
	balance, err := c.Contract.Instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return 0, fmt.Errorf("Failed to query balance: %v", err)
	}
	return balance.Int64(), nil
}

func (c *Client) QueryAllowance(from common.Address, spender common.Address) (int64, error){

	allowance, err := c.Contract.Instance.Allowance(&bind.CallOpts{}, from, spender)
	if err != nil {
		return 0, fmt.Errorf("Failed to query allowance: %v", err)
	}
	return allowance.Int64(), nil

}

func (c *Client) AproveAllowance(spender common.Address, amount int64) error {

	nonce, err := c.eth.PendingNonceAt(context.Background(), c.Account.Address)
	if err != nil {
		return fmt.Errorf("Failed to AproveAllowance: %v", err)
	}

	gasPrice, err := c.eth.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to AproveAllowance: %v", err)
	}

	auth := bind.NewKeyedTransactor(c.Account.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) 
	auth.GasLimit = uint64(3000000) 
	auth.GasPrice = gasPrice

	c.Contract.Instance.Approve(auth, spender, big.NewInt(amount))
	return nil
}

func (c *Client) Transfer(to common.Address, amount int64) (common.Hash, error) {
	nonce, err := c.eth.PendingNonceAt(context.Background(), c.Account.Address)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Failed to Transfer: %v", err)
	}

	gasPrice, err := c.eth.SuggestGasPrice(context.Background())
	if err != nil {
		return common.Hash{}, fmt.Errorf("Failed to Transfer: %v", err)
	}

	auth := bind.NewKeyedTransactor(c.Account.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) 
	auth.GasLimit = uint64(3000000) 
	auth.GasPrice = gasPrice

	byts, err := c.Contract.Instance.Transfer(auth, to, big.NewInt(amount))
	if nil != err {
		return common.Hash{}, fmt.Errorf("err: %v \n", err)
	}

	var tx types.Transaction
	err = rlp.DecodeBytes(byts, &tx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("err: %v \n", err)
	}

	fmt.Printf("result: %v\n", tx)
	return tx, nil

}


func (c *Client) TransferFrom(from common.Address, to common.Address, amount int64) error {

	nonce, err := c.eth.PendingNonceAt(context.Background(), c.Account.Address)
	if err != nil {
		return fmt.Errorf("Failed to TransferFrom: %v", err)
	}

	gasPrice, err := c.eth.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to TransferFrom: %v", err)
	}

	auth := bind.NewKeyedTransactor(c.Account.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) 
	auth.GasLimit = uint64(3000000) 
	auth.GasPrice = gasPrice

	c.Contract.Instance.TransferFrom(auth, from, to, big.NewInt(amount))

	return nil
}
