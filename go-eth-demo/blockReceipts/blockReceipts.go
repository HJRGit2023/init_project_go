package blockreceipts

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func Run() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/Hb2vSp12j64yRzffi691S")
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber) // 这里很慢
	if err != nil {
		log.Fatal(err)
	}
	block1, err := client.BlockByHash(context.Background(), common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Block by number hash: %s", block1.Hash().Hex())
	log.Printf("Block by hash: %s", block.Hash().Hex())

	// blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	// 通过  区块hash  获取区块的收据 ，第二个参数onlyHash设为true，确保只按哈希取对应区块的收据
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(block.Hash(), true))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Receipts for block %s: %v", block.Hash().Hex(), receiptByHash)
	// 通过  区块高度  获取区块的收据
	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("receiptsByNum for block %s: %+v", block.Hash().Hex(), receiptsByNum[0])
	// false,虽然两个结构体的内容一样，但是地址是不同的，所以不相等
	log.Println(receiptByHash[0] == receiptsByNum[0])
	log.Println("receiptByHash[0]:", receiptByHash[0])
	log.Println("receiptsByNum[0]:", receiptsByNum[0])
	log.Printf("receiptByHash[0]:%p", receiptByHash[0])
	log.Printf("receiptsByNum[0]:%p", receiptsByNum[0])
	// 深比较两个结构体的内容
	if reflect.DeepEqual(receiptByHash[0], receiptsByNum[0]) {
		log.Println("内容一致")
	} else {
		log.Println("内容不一致")
	}
	data, _ := json.MarshalIndent(receiptByHash[0], "", "  ")
	log.Printf("Receips (JSON):\n%s", string(data))
	for _, receipt := range receiptByHash {
		log.Printf("Receipt status: %d, receipt.Logs: %v\n, receipt.TxHash: %s, receipt.TransactionIndex: %d", receipt.Status, receipt.Logs, receipt.TxHash.Hex(), receipt.TransactionIndex)
	}
	log.Println("----------------------------------------------")
	// for _, receipt := range receiptsByNum {
	// 	log.Printf("receiptsByNum status: %d, receipt.Logs: %v\n, receipt.TxHash: %s, receipt.TransactionIndex: %d", receipt.Status, receipt.Logs, receipt.TxHash.Hex(), receipt.TransactionIndex)
	// }
	// 使用交易哈希查询，调用 TransactionReceipt 方法
	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0xe864ac24bf3de1d93f3df94941b2df14163d6de87f8ff3621af36c17b4cb8f13"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("receipt.Status: ", receipt.Status) // 1
	log.Println("receipt.Logs:", receipt.Logs)      //
}
