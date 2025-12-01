package blockhead

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Run() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		log.Fatal(err)
	}
	// blockNumber := big.NewInt(5671744)
	// client.HeaderByNumber(context.Background(), blockNumber)
	// blockNumber is nil表示获取最新区块的header
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("header number:", header.Number.String())
	fmt.Println("header hash:", header.Hash)
	fmt.Println("header parent hash:", header.ParentHash.Hex())
	fmt.Println("header uncle hash:", header.UncleHash.Hex())
	fmt.Println("header coinbase:", header.Coinbase.Hex())
	fmt.Println("header root:", header.Root.Hex())
	fmt.Println("header tx hash:", header.TxHash.Hex())
	fmt.Println("header receipt hash:", header.ReceiptHash.Hex())
	fmt.Println("header bloom:", header.Bloom)
	fmt.Println("header difficulty:", header.Difficulty.String())
	fmt.Println("header Number:", header.Number.Int64())
	fmt.Println("header gas limit:", header.GasLimit)
	fmt.Println("header gas used:", header.GasUsed)
	fmt.Println("header time:", header.Time)
	fmt.Println("header extra:", header.Extra)
	fmt.Println("header mix digest:", header.MixDigest.Hex())

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block hash:", block.Hash().Hex())                // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	fmt.Println("block number:", block.Number().Uint64())         // 5671744
	fmt.Println("block difficulty:", block.Difficulty().String()) // 0
	fmt.Println("block timestamp:", block.Time())                 // 1712798400
	fmt.Println("block transactions:", block.Transactions())      // [很多值]
	fmt.Println("block uncles:", block.Uncles())                  // []
	fmt.Println("block size:", block.Size())                      // 73003
	fmt.Println("block gas limit:", block.GasLimit())             // 30000000
	fmt.Println("block gas used:", block.GasUsed())               // 10457625
	fmt.Println("block nonce:", block.Nonce())                    // 0
	fmt.Println("block extra:", block.Extra())                    // [216 131 1 13 11 132 103 101 116 104 136 103 111。。。]
	// 通过block hash获取交易数量
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(" block transactions count:", count) // 70
}
