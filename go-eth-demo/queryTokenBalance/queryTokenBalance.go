package querytokenbalance

import (
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	token "github.com/local/go-eth-demo/erc20"
)

func Run() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d9049ebd315048ab81699e12e1d1fac1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Token contract address
	tokenAddress := common.HexToAddress("0x20a556Bd66C552Dfff497F3D1523F0Bd5EEAF408")

	// Get the token contract instance
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Address of the account to query the token balance
	accountAddress := common.HexToAddress("0xd865a5913887b5790137b589ab81ba3fef6bcf0a")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, accountAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Token balance of account %s is %d", accountAddress.Hex(), bal)

	// Token symbol
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Token symbol is %s", symbol)
	// Token name
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Token name is %s", name)
	// Token decimals
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Token decimals is %d", decimals)
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	log.Printf("Token balance of account %s is %s", accountAddress.Hex(), value.String())
}
