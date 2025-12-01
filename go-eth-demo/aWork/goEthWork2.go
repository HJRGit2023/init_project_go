package aWork

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Run2() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		log.Fatal(err)
	}
	// 获取私钥
	privateKey, err := crypto.HexToECDSA("账户私钥")
	if err != nil {
		log.Fatal(err)
	}
	// 获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type of publicKey to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("From address: %s", fromAddress.Hex())
	// 获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 获取gasPrice
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = big.NewInt(1000000)
	auth.GasLimit = uint64(300000) // in units
	// 部署合约
	contractAdr, tx, instance, err := DeployAWork(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Contract address: %s", contractAdr.Hex())
	log.Printf("Transaction hash: %s", tx.Hash().Hex())
	// 调用合约方法
	auth.GasPrice = big.NewInt(10000000)
	auth.GasLimit = uint64(10000000) // in units
	nonce1, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce1))
	auth.Value = big.NewInt(0) // in wei
	result, err := instance.CounterAdd(auth)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CounterAdd transaction hash: %s", result.Hash().Hex())
	// 等待交易被打包
	bind.WaitMined(context.Background(), client, result)
	// 查询合约状态
	callOpt := &bind.CallOpts{Context: context.Background()}
	result2, err := instance.Count(callOpt)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Counter value: %s", result2.String())

}
