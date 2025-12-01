package aWork

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
	// 获取ethclient客户端,使用 ethclient 连接到 Sepolia 测试网络
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		log.Fatal(err)
	}
	// 查询区块信息
	blockNumber := big.NewInt(9732767) // 区块高度
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("区块高度：", block.Number().Int64())
	log.Println("区块哈希：", block.Hash().Hex())
	log.Println("区块父哈希：", block.ParentHash().Hex())
	log.Println("区块难度：", block.Difficulty().String())
	log.Println("区块时间戳：", block.Time())
	log.Println("区块交易数量：", len(block.Transactions()))
	// 查询账户信息
	// 使用私钥地址，获取私钥对象
	privateKey, err := crypto.HexToECDSA("账户私钥")
	if err != nil {
		log.Fatal(err)
	}
	// 获取公钥对象
	publicKey := privateKey.Public()
	// 类型断言，将公钥对象转换为ecdsa.PublicKey类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // ok为true表示转换成功
	if !ok {                                           // bool类型断言，ok为false表示转换失败
		log.Fatal(err)
	}
	// 将ecdsa.PublicKey类型公钥对象转换为地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Println("fromAddress:", fromAddress.Hex())
	toAddress := common.HexToAddress("0xb500295aa4ae9e36f420601603634f11ff21a6c9")
	// 转账金额
	value := big.NewInt(1000000000000000) // in wei (1 eth) = 10^18 wei
	// nonce值，用于防止重放攻击
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 获取gasPrice，用于设置交易的gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gaslimit := uint64(21000) // 21000 是默认的gasLimit值
	// 构建交易对象
	tx := types.NewTransaction(nonce, toAddress, value, gaslimit, gasPrice, nil)
	// 签名交易对象
	// 链ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 使用私钥和链ID对交易对象进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("交易发送成功：", signedTx.Hash().Hex())

}
