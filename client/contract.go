package client


import (
	"fmt"
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	abi "github.com/2cluster/ethclient/client/contract"
)

var VALUE_URL = map[string]string{
	"INFURA"		: "https://rinkeby.infura.io/v3/634d2ee71c3e44a4ab4990f90f561398",
	"LOCAL1" 		: "127.0.0.1:8545",
	"LOCAL2" 		: "localhost:8545",
	"LOCAL3" 		: "devchain:8545",
}

var CONFIRMATIONS = uint64(1)

type Contract struct {
	Name string
	Instance abi.SUSD
	Address common.Address
	Tx *ethtypes.Transaction
}


func (c *Client) DeployContract() error {
	ctx := context.Background()
	nonce, err := c.Eth.PendingNonceAt(ctx, c.Account.Address)
	if err != nil {
		return err
	}

	gasPrice, err := c.Eth.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(c.Account.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	address, tx, instance, err := abi.DeploySUSD(auth, c.Eth)
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
	instance, err := abi.NewSUSD(address, c.Eth)
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

func (c *Client) QueryAllowance(from common.Address, spender common.Address) (int64, error) {

	allowance, err := c.Contract.Instance.Allowance(&bind.CallOpts{}, from, spender)
	if err != nil {
		return 0, fmt.Errorf("Failed to query allowance: %v", err)
	}
	return allowance.Int64(), nil

}

func (c *Client) ApproveAllowance(spender common.Address, amount int64) (string, error) {

	nonce, err := c.Eth.PendingNonceAt(context.Background(), c.Account.Address)
	if err != nil {
		return "", fmt.Errorf("Failed to AproveAllowance: %v", err)
	}

	gasPrice, err := c.Eth.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("Failed to AproveAllowance: %v", err)
	}

	auth := bind.NewKeyedTransactor(c.Account.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) 
	auth.GasLimit = uint64(3000000) 
	auth.GasPrice = gasPrice

	tx, err := c.Contract.Instance.Approve(auth, spender, big.NewInt(amount))
	if err != nil {
		return "", fmt.Errorf("err: %v \n", err)
	}

	receipt, err := waitMined(context.Background(), c.Eth, tx)
	if err != nil {
		return "", fmt.Errorf("err: %v \n", err)
	}

	return receipt, nil
}

func (c *Client) Transfer(to common.Address, amount int64) (string, error) {
	nonce, err := c.Eth.PendingNonceAt(context.Background(), c.Account.Address)
	if err != nil {
		return "", fmt.Errorf("Failed to Transfer: %v", err)
	}

	gasPrice, err := c.Eth.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("Failed to Transfer: %v", err)
	}

	auth := bind.NewKeyedTransactor(c.Account.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) 
	auth.GasLimit = uint64(3000000) 
	auth.GasPrice = gasPrice

	tx, err := c.Contract.Instance.Transfer(auth, to, big.NewInt(amount))
	if err != nil {
		return "", fmt.Errorf("err: %v \n", err)
	}

	receipt, err := waitMined(context.Background(), c.Eth, tx)
	if err != nil {
		return "", fmt.Errorf("err: %v \n", err)
	}

	return receipt, nil

}


func (c *Client) TransferFrom(from common.Address, to common.Address, amount int64) (string, error) {

	nonce, err := c.Eth.PendingNonceAt(context.Background(), c.Account.Address)
	if err != nil {
		return "", fmt.Errorf("Failed to TransferFrom: %v", err)
	}

	gasPrice, err := c.Eth.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("Failed to TransferFrom: %v", err)
	}

	auth := bind.NewKeyedTransactor(c.Account.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) 
	auth.GasLimit = uint64(3000000) 
	auth.GasPrice = gasPrice

	tx, err := c.Contract.Instance.TransferFrom(auth, from, to, big.NewInt(amount))
	if err != nil {
		return "", fmt.Errorf("err: %v \n", err)
	}

	receipt, err := waitMined(context.Background(), c.Eth, tx)
	if err != nil {
		return "", fmt.Errorf("err: %v \n", err)
	}

	return receipt, nil
}

func waitMined(ctx context.Context, conn *ethclient.Client, tx *ethtypes.Transaction) (string, error) {
	receipt, err := WaitMinedWithTxHash(ctx, conn, tx.Hash().Hex(), CONFIRMATIONS)
	if err != nil {
		fmt.Errorf("err: %v \n", err)
	}
	if receipt.Status == 0 {
		return "", fmt.Errorf("Transaction Failed")
	}

	return tx.Hash().Hex(), nil
}

// WaitMined waits for tx to be mined on the blockchain
// It returns tx receipt when the tx has been mined and enough block confirmations have passed
func WaitMinedWithTxHash(ctx context.Context, ec *ethclient.Client,
	txHash string, blockDelay uint64) (*ethtypes.Receipt, error) {
	// an error possibly returned when a transaction is pending
	const missingFieldErr = "missing required field 'transactionHash' for Log"

	if ec == nil {
		return nil, fmt.Errorf("Nill client")
	}
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	// wait tx to be mined
	txHashBytes := common.HexToHash(txHash)
	fmt.Printf("\n")
	fmt.Printf("Waiting for transaction to be mined.")
	for {
		receipt, rerr := ec.TransactionReceipt(ctx, txHashBytes)
		fmt.Printf(".")
		if rerr == nil {
			fmt.Printf("\n")
			fmt.Printf("Transaction written to block.\n")
			fmt.Printf("\n")
			fmt.Printf("https://rinkeby.etherscan.io/tx/" + txHash + "\n")
			fmt.Printf("\n")
			fmt.Printf("Waiting for %d block confirmations.", blockDelay)

			if blockDelay == 0 {
				return receipt, rerr
			}
			break
		} else if rerr == ethereum.NotFound || rerr.Error() == missingFieldErr {
			// Wait for the next round
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return nil, ctx.Err()
			case <-queryTicker.C:
			}
		} else {
			return receipt, rerr
		}
	}
	// wait for enough block confirmations
	ddl := big.NewInt(0)
	latestBlockHeader, err := ec.HeaderByNumber(ctx, nil)
	if err == nil {
		ddl.Add(new(big.Int).SetUint64(blockDelay), latestBlockHeader.Number)
	}
	for {
		latestBlockHeader, err := ec.HeaderByNumber(ctx, nil)
		fmt.Printf(".")
		if err == nil && ddl.Cmp(latestBlockHeader.Number) < 0 {
			receipt, rerr := ec.TransactionReceipt(ctx, txHashBytes)
			if rerr == nil {
				fmt.Println("\ntx confirmed!\n")
				return receipt, rerr
			} else if rerr == ethereum.NotFound || rerr.Error() == missingFieldErr {
				return nil, fmt.Errorf("\ntx is dropped due to chain re-org\n")
			} else  {
				return receipt, rerr
			}
		}
		select {
		case <-ctx.Done():
			fmt.Errorf("err: %v \n", err)
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}
