package executioncontracts

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/local/go-eth-demo/store"
)

const (
	contractAddr = "0x8129D1357e64b3ef220503bb8B7f6074C58160AA"
)

func Run() {
	// 创建 ethclient 实例
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		log.Fatal(err)
	}
	// 导入私钥, 根据 hex 创建私钥实例
	privateKey, err := crypto.HexToECDSA("账户私钥")
	if err != nil {
		log.Fatal(err)
	}
	// 创建合约实例:创建一个新的实例执行一个store合约
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	// 调用合约方法
	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value11111"))
	//
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("tx hash: %s", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ret: %s", string(valueInContract[:]))
	log.Println("is value saving in contract equals to origin value:", valueInContract == value)
}
