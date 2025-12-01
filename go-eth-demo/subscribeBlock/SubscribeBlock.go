package subscribeblock

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Run() {
	// https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1
	// wss://sepolia.infura.io/ws/v3/d9049ebd315048ab81699e12e1d1fac1
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// 创建一个通道，用于接收新区块的header
	headers := make(chan *types.Header)
	// ch := make(chan *ethclient.BlockHeader)
	// 订阅新区块的header
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	// 订阅将 推送新的区块头事件  到我们的通道，因此我们可以使用一个 select 语句来监听新消息
	for {
		select {
		case header := <-headers:
			log.Printf("New block number: %d", header.Number.Uint64())
			log.Printf("New block hash: %s", header.Hash().Hex())
			// 获得该区块的完整内容，我们可以将区块头的摘要传递给客户端的 BlockByHash 函数
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			// log.Printf("Block number: %d", block.Number().Uint64())
			log.Println(block.Hash().Hex())      // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			log.Println(block.Number().Uint64()) // 3477413
			log.Println(block.Time())            // 1529525947
			log.Println(block.Nonce())           // 130524141876765836fmt.Println(len(block.Transactions())) // 7
		case err := <-sub.Err():
			log.Fatal(err)
		}
	}

}
