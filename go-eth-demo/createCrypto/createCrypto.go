package createcrypto

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func Run() {
	// 生成一个新的钱包，导入 go-ethereum crypto 包，该包提供用于生成随机私钥的 GenerateKey 方法。
	// privateKey, err := crypto.GenerateKey()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// 打印私钥
	// log.Printf("私钥: %x", privateKey)
	// privateKeyBytes := crypto.FromECDSA(privateKey)
	// 打印私钥的十六进制编码, 去掉前缀 0x ,
	// 用于签署交易的私钥，将被视为密码，永远不应该被共享给别人，因为谁拥有它可以访问你的所有资产。
	// log.Println("私钥十六进制编码:", hexutil.Encode(privateKeyBytes)[2:]) // ec686999e0ddfcbe6d85c4a6f3f9481639cd1f1517f8c785a70d772571576e8e
	// 通过私钥的 Hex 字符串，使用HexToECDSA 方法恢复为私钥对象
	privateKey, err := crypto.HexToECDSA("ec686999e0ddfcbe6d85c4a6f3f9481639cd1f1517f8c785a70d772571576e8e")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("私钥的 Hex 字符串恢复的私钥: %x", privateKey)
	// 公钥是从私钥派生的，因此 go-ethereum 的加密私钥具有一个返回公钥的 Public 方法
	publicKey := privateKey.Public()
	// 断言公钥是否为 ECDSA 类型,断言是接口类型转为具体类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// 打印公钥的十六进制编码, 去掉前缀 0x 和 04(是ECDSA的标识符)
	log.Println("公钥十六进制编码:", hexutil.Encode(publicKeyBytes)[4:])
	// 公钥的哈希值，用于标识公钥，通常作为地址使用。
	// 我们拥有公钥，就可以轻松生成你经常看到的公共地址。
	// go-ethereum 加密包有一个 PubkeyToAddress 方法，它接受一个 ECDSA 公钥，并返回公共地址。
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex() //
	log.Printf("公钥哈希值地址格式化是bytes: %x", address)
	log.Printf("公钥哈希值地址格式化是string: %s", address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	// 完整的公钥哈希值，包括标识符和公钥，通常用于标识公钥，但不适合作为地址。
	log.Println("full完整公钥哈希值:", hexutil.Encode(hash.Sum(nil)[:]))
	address1 := hexutil.Encode(hash.Sum(nil)[12:]) // 共32字节，截取公钥哈希值后20字节作为地址
	log.Printf("公钥哈希值地址: %s", address1)
}
