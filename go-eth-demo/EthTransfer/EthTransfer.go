package ethtransfer

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Run() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		log.Fatal(err)
	}
	// privateKey, err := crypto.GenerateKey()
	// 通过私钥地址获取私钥对象,这里是自己metamask钱包的私钥，不能公开，后续打乱输入
	privateKey, err := crypto.HexToECDSA("账户私钥")
	if err != nil {
		log.Fatal(err)
	}

	// 打印私钥
	log.Printf("私钥: %s\n", privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("From address: %s\n", fromAddress) // 0xd865a5913887b5790137B589Ab81ba3Fef6BCf0a
	// 构建转账金额
	value := big.NewInt(1000000000000000) // in wei (1 eth) = 10^18 wei
	// 通过PendingNonceAt方法获取应该使用的下一个nonce值，是我们用于帐户交易的随机数。
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Nonce: %d\n", nonce) // 2
	gasLimit := uint64(21000)        // in units
	// 根据SuggestGasPrice方法获取建议的gas价格，单位是wei。根据'x'个先前块来获得平均燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Gas price: %d\n", gasPrice) // 1000001
	// 构建接收方地址
	toAddress := common.HexToAddress("0xd865a5913887b5790137b589ab81ba3fef6bcf0a")
	// 构建交易对象
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	// 签名交易对象
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 发送交易,将已签名的事务广播到整个网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
}
