package querybalance

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Run() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		log.Fatal(err)
	}
	account := common.HexToAddress("0xd865a5913887b5790137b589ab81ba3fef6bcf0a")
	// nil is the block number, which means the latest block，传nil表示获取最新块的余额
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Account balance at latest block:", balance.String())

	blockNumber := big.NewInt(5532993)
	balance1, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Account balance at block 5532993:", balance1.String())
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethvalue := new(big.Float).Quo(fbalance, big.NewFloat(1e18))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Account balance at latest block  in ETH:", ethvalue.String())
	// 待处理账户余额
	pendingBalanceAt, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Pending balance at latest block:", pendingBalanceAt.String())

}
