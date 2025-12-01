package loadcontracts

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/local/go-eth-demo/store"
)

const (
	contractAddr = "0x8129D1357e64b3ef220503bb8B7f6074C58160AA"
)

func Run() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		log.Fatal(err)
	}
	// 使用store.go中的NewStore方法加载合约
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	// 使用%s打印storeContract
	log.Printf("Store contract loaded: %s", storeContract)
	_ = storeContract
}
