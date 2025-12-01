package tokentransfer

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	sha3 "golang.org/x/crypto/sha3"
)

func Run() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		log.Fatal(err)
	}

	// 通过私钥hash获取私钥对象
	privateKey, err := crypto.HexToECDSA("账户私钥")
	if err != nil {
		log.Fatal(err)
	}
	// 通过私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type of publicKey to *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 账户的nonce，这里是随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("账户的nonce：%d", nonce) // 第一次nonce 4  第二次nonce5
	// 计算方法 ID
	// 这里是MyToken合约的地址
	tokenAddress := common.HexToAddress("0x20a556Bd66C552Dfff497F3D1523F0Bd5EEAF408")
	// transfer(address,uint256)函数的签名
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	log.Println(hexutil.Encode(methodID)) //0xa9059cbb

	// 给我们发送代币的地址左填充到 32 字节
	toAddress := common.HexToAddress("0xb500295aa4ae9e36f420601603634f11ff21a6c9")
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	log.Println("左填充后的地址：", hexutil.Encode(paddedAddress)) // 0x000000000000000000000000b500295aa4ae9e36f420601603634f11ff21a6c9
	// 发送多少个代币
	amount := new(big.Int)
	amount.SetString("1000000000000000", 10) // 转的是tokens 不是ether
	// 代币量也需要左填充到 32 个字节
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	log.Println("左填充后的代币量：", hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
	// 组装数据字段字节片，只需将方法 ID，填充后的地址和填后的转账量拼接起来即可
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	value := big.NewInt(0)
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	gasPrice := big.NewInt(1000000000)    // 调大gasPrice
	log.Println("建议的gasPrice：", gasPrice) // 第一次1000017，第二次 1026310
	// gaslimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	To: &toAddress,
	// 	// GasPrice: big.NewInt(1000000000),
	// 	// Value:    big.NewInt(0),
	// 	Data: data,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	gaslimit := uint64(100000)            // 增大gaslimit
	log.Println("估算的gasLimit：", gaslimit) // 第一次22915，第二次 22885

	// 构建交易
	tx := types.NewTransaction(nonce, tokenAddress, value, gaslimit, gasPrice, data)
	// 签名交易
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("tx sent: %s", signedTx.Hash().Hex())
}
