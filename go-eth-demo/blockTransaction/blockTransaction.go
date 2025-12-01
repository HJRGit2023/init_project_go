package blocktransaction

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Run() {
	// 两个地址获取到的client是一样的，
	// client, err := ethclient.Dial("https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1")
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/Hb2vSp12j64yRzffi691S")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// blockNumber, err := client.BlockNumber(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	blockNumber1 := big.NewInt(5671744)
	// 需要把blockNumber转换成big.Int类型才能使用
	block, err := client.BlockByNumber(context.Background(), blockNumber1) // 这里很慢
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block number: %d", block.NumberU64())             // 5671744
	log.Printf("Block hash: %s", block.Hash().Hex())              //0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	log.Printf("Block parent hash: %s", block.ParentHash().Hex()) // 0xf6e3e5c046be170caffa2f330cb3fc5264cf704ddd9089cb3c70d55f6f938816
	log.Printf("Block nonce: %d", block.Nonce())                  // 0
	log.Printf("Block difficulty: %d", block.Difficulty())        // 0
	log.Printf("Block gas limit: %d", block.GasLimit())           // 30000000
	log.Printf("Block gas used: %d", block.GasUsed())             // 10457625
	// block.Transactions()返回一个交易列表，循环遍历集合并获取交易的信息。
	for idx, tx := range block.Transactions() {
		log.Printf("Transaction hash: %s", tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5

		// 先从客户端获取chainID，然后使用chainID创建签名者，再使用签名者发送交易
		chainID, err := client.ChainID(context.Background()) // 11155111
		if err != nil {
			log.Fatal(err)
		}
		// transaction type not supported,这里用EIP155Signer签名，不能处理所有交易类型，
		// 改为NewLondonSigner签名，可以处理所有交易类型。
		// if sender, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
		// 	log.Println("tx type :", tx.Type(), " sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
		// } else {
		// 这里有些交易无法解析，跳过
		// log.Println("tx type :", tx.Type(), " sender", "unknown")
		// continue
		// log.Fatal(err)
		// }
		var signer types.Signer
		switch tx.Type() {
		case types.LegacyTxType:
			signer = types.NewEIP155Signer(chainID)
		case types.AccessListTxType:
			signer = types.NewEIP155Signer(chainID)
		case types.DynamicFeeTxType:
			signer = types.NewLondonSigner(chainID)
		case types.BlobTxType:
			signer = types.NewCancunSigner(chainID)
		default:
			log.Fatal("tx type not supported")
		}
		sender, err := types.Sender(signer, tx)
		if err == nil {
			log.Printf("tx type : %d, sender: %s", tx.Type(), sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
		} else {
			log.Printf("Failed to get sender for tx %d: %v", idx, err)
			continue
		}

		// 传入链ID和交易类型，自动生成对应签名器, AsMessage这个也不行
		// signer1 := types.MakeSigner(types.NewChainConfig(chainID, nil), block.Number())

		// 在使用 EIP155Signer 还原出 sender 地址：AsMessage这个也不行
		// if msg, err := signer.AsMessage(tx, nil); err == nil {
		// 	log.Printf("From: %s, To: %s, GasLimit: %d, GasPrice: %d, Value: %d", msg.From().Hex(), msg.To().Hex(), msg.GasLimit(), msg.GasPrice(), msg.Value())
		// } else {
		// 	log.Fatal(err)
		// }
		log.Printf("From: %s, To: %s, GasLimit: %d, GasPrice: %d, Value: %d", sender, tx.To().Hex(), tx.Gas(), tx.GasPrice(), tx.Value())
		// 解析交易数据 每个交易都有一个收据，其中包含执行交易的结果，
		// 例如所有的返回值和日志， 以及“1”（成功）或“0”（失败）的交易结果状态。
		if receipt, err := client.TransactionReceipt(context.Background(), tx.Hash()); err == nil {
			log.Printf("receipt logs: %v\n, Transaction status: %d", receipt.Logs, receipt.Status) // [], 1
		} else {
			log.Fatal(err)
		}
		break // 这里只取第一个交易
	}
	// 这里直接通过blockhash获取交易数量
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		// 调用客户端的 TransactionInBlock 方法，这里通过blockhash和交易索引获取交易
		// 此方法仅接受块哈希和块内事务的索引值。
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	}

	// 使用 TransactionByHash 在给定具体事务哈希值的情况下直接查询单个事务。
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transaction hash: %s, is pending: %t", tx.Hash().Hex(), isPending) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2, false
}
