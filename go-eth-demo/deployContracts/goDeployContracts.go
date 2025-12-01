package deploycontracts

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	store "github.com/local/go-eth-demo/store"
)

func Run() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// privateKey, err := crypto.GenerateKey()
	// privateKeyBytes := crypto.FromECDSA(privateKey)
	// privateKeyHex := hex.EncodeToString(privateKeyBytes)
	// fmt.Println("Private Key:", privateKeyHex)
	privateKey, err := crypto.HexToECDSA("账户私钥")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyHex := hex.EncodeToString(publicKeyBytes)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) //
	log.Println("Public Key:", publicKeyHex, "Address:", fromAddress)

	// 获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Nonce:", nonce)
	// 获取gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Gas Price:", gasPrice)
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Chain ID:", chainId)
	// NewKeyedTransactor 是一个实用方法，用于轻松地从单个私钥创建交易签名者。
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contract Address:", address.Hex())
	log.Println("Transaction Hash:", tx.Hash().Hex())

	_ = instance

}
