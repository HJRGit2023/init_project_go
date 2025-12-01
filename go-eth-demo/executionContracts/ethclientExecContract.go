package executioncontracts

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAddress = "0x8129D1357e64b3ef220503bb8B7f6074C58160AA"
)

func Run1() {
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
	contractABI, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		log.Fatal(err)
	}
	methodName := "setItem"
	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_save_key_use_abi"))
	copy(value[:], []byte("demo_save_value_use_abi_11111"))
	input, err := contractABI.Pack(methodName, key, value)
	if err != nil {
		log.Fatal(err)
	}
	// 创建并签名交易，input不为空表示调用合约方法
	chainID := big.NewInt(int64(11155111))
	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddress), big.NewInt(0), 200000, gasPrice, input)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// auth := bind.NewKeyedTransactor(privateKey)
	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)     // in wei
	// auth.GasLimit = uint64(3000000) // in units
	// auth.GasPrice = gasPrice
	// address := common.HexToAddress("0x14b0Ed2a7C4cC60DD8F672059828eB19D0617F63")
	// var amount uint64 = 10000000
	// var data []byte
	// amount不为0，data不为空时，表示转账交易
	// tx := types.NewTransaction(nonce, address, amount, gasLimit, gasPrice, data)
	// signedTx, err := auth.Sign(tx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// 发送签名好的交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	// 等待交易被打包进区块
	_, err = waitForReceipt(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transaction mined: %s\n", signedTx.Hash().Hex())
	// 创建 call 查询
	callInput, err := contractABI.Pack("items", key)
	if err != nil {
		log.Fatal(err)
	}
	to := common.HexToAddress(contractAddr)
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
	err = contractABI.UnpackIntoInterface(&unpacked, "items", result)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Value of key1: %s\n", string(unpacked[:]))
	log.Println("is value saving in contract equals to origin value:", unpacked == value)
}

func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
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
