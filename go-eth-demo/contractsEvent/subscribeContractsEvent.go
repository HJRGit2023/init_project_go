package contractsevent

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var StoreABI2 = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

func Run1() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		panic(err)
	}
	// 合约地址
	contractAddress := common.HexToAddress("0x8129D1357e64b3ef220503bb8B7f6074C58160AA")
	// 构造 FilterQuery，指定要从哪个块开始过滤，哪个块结束过滤，导入 FilterQuery 结构体并用过滤选项初始化它
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(9729380),
		// ToBlock:   big.NewInt(9729380),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	// 声明一个 channel，用来接收日志数据
	logs := make(chan types.Log)
	// 调用 ethclient 的 FilterLogs，它接收我们的查询并将返回所有的匹配事件日志
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	// 解码日志数据，abi.JSON 接收 ABI 字符串并返回一个 abi.ABI 对象，即解析过的ABI接口
	contractAbi, err := abi.JSON(strings.NewReader(StoreABI2))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vlog := <-logs:
			log.Printf("Log BlockNumber: %d", vlog.BlockNumber)
			log.Printf("Log TxHash: %s", vlog.TxHash.Hex())
			log.Printf("Log Address: %s", vlog.Address.Hex())
			log.Printf("Log Data: %s", vlog.Data)
			// event := new(StoreEvent)
			event := struct { // 匿名结构体event
				Key   [32]byte
				Value [32]byte
			}{}
			// 解析后的 ABI 接口(contractAbi)的 Unpack 函数,三个参数：
			// 第一个参数&event，第二个参数是要尝试解码的事件名称ItemSet，第三个参数是编码的日志数据vlog.Data
			// 解码日志数据，UnpackIntoInterface 接收一个结构体指针和事件签名，并将日志数据解码到结构体中
			err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vlog.Data)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Log Key: %s", common.Bytes2Hex(event.Key[:]))
			log.Printf("Log Value: %s", common.Bytes2Hex(event.Value[:]))
			var topics []string
			for i := range vlog.Topics {
				topics = append(topics, vlog.Topics[i].Hex())
			}
			log.Printf("Log Topics[0]: %s", topics[0])
			if len(topics) > 1 {
				log.Println("indexed topics:", topics[1:])
			}
		}
	}
}
