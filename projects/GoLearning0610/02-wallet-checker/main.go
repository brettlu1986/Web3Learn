package main

//单元测试 go test ./02-wallet-checker/wallet  或 go test ./...  执行素有单元测试
// go test -v ./02-wallet-checker/wallet 显示单元测试所有的过程
//terminal 测试： go run ./02-wallet-checker 0x1234567890123456789012345678901234567890
import (
	walllet "GoLearning0610/02-wallet-checker/wallet"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run ./02-wallet-checker <ethereum-address>")
		return
	}

	address := os.Args[1]

	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])

	if err := walllet.ValidateAddress(address); err != nil {
		fmt.Println("invalid address:", err)
		return
	}

	fmt.Println("valid ethereum address", address)
}
