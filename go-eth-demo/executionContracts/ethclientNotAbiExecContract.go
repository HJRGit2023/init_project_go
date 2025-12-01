package executioncontracts

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAdr = "0x8129D1357e64b3ef220503bb8B7f6074C58160AA"
)

func Run2() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// 通过私钥hash获取私钥对象
	privateKey, err := crypto.HexToECDSA("账户私钥")
	if err != nil {
		log.Fatal(err)
	}
	// 从私钥实例获取公开地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type of publicKey")
	}
	// 账户地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 使用地址获取地址的 nonce 值:
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 估算 gas 价格：
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 准备交易 calldata
	methodSignature := []byte("setItem(bytes32,bytes32)")
	methodSelector := crypto.Keccak256(methodSignature)[:4] // 获取函数选择器

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_save_key_no_use_abi"))
	copy(value[:], []byte("demo_save_value_no_use_abi_11111"))

	// 组合调用数据
	var input []byte
	input = append(input, methodSelector...)
	input = append(input, key[:]...)
	input = append(input, value[:]...)

	// 创建并签名交易，input不为空表示调用合约方法
	chainID := big.NewInt(int64(11155111))
	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddress), big.NewInt(0), 200000, gasPrice, input)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送签名好的交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	// 等待交易被打包进区块
	_, err = waitForReceipt2(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transaction mined: %s\n", signedTx.Hash().Hex())
	// 创建 call 查询
	itemSignature := []byte("items(bytes32)")
	itemsSelector := crypto.Keccak256(itemSignature)[:4] // 获取函数items(bytes32)选择器

	var callInput []byte
	callInput = append(callInput, itemsSelector...)
	callInput = append(callInput, key[:]...)
	to := common.HexToAddress(contractAdr)
	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: callInput,
	}
	// 解析返回值
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}
	var unpacked [32]byte
	copy(unpacked[:], result)
	log.Printf("Value of key1: %s\n", string(unpacked[:]))
	log.Println("is value saving in contract equals to origin value:", unpacked == value)
}

func waitForReceipt2(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if receipt != nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		// 等待一段时间后再次查询
		time.Sleep(1 * time.Second)
	}
}
