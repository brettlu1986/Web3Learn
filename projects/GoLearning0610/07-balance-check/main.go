package main

// 1. 创建 resultCh
// 2. 创建 WaitGroup
// 3. 遍历 addresses
// 4. 每个地址 wg.Add(1)
// 5. 每个地址启动一个 goroutine 查询余额
// 6. 额外启动一个 goroutine，等待所有任务完成后关闭 resultCh
// 7. main goroutine 进入 for range resultCh，等待接收结果
// 8. 某个查询 goroutine 查完，发送 BalanceResult 到 resultCh
// 9. main 收到结果并打印
// 10. 每个查询 goroutine 执行 defer wg.Done()
// 11. 所有查询完成后，wg.Wait() 结束
// 12. 关闭 resultCh
// 13. for range resultCh 结束
// 14. main 函数结束

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BalanceResult struct {
	Address string
	Balance int
}

//我有多个钱包地址，不想一个个慢慢查余额，而是同时发起多个查询，谁先查完谁先把结果交回来。

func main() {
	addresses := []string{
		"0x111",
		"0x222",
		"0x333",
		"0x444",
	}

	resultCh := make(chan BalanceResult) //创建一个goroutine 之间传递 BalanceResult数据的通道
	var wg sync.WaitGroup                //等待一组 goroutine 全部执行完成。

	for _, address := range addresses {
		wg.Add(1) //启动goroutine前先+1， 代表增加一个要等待的任务

		go func(addr string) {
			defer wg.Done() //一个任务完成

			balance := QueryBalance(addr)
			resultCh <- BalanceResult{ //往通道里送数据
				Address: addr,
				Balance: balance,
			}
		}(address)
	}

	//开一个 goroutine，专门等待所有查询任务完成；等全部完成后，关闭 resultCh。
	go func() {
		wg.Wait() //阻塞等待 知道所有任务完成
		close(resultCh)
	}()

	//一直等channel里的数据
	for result := range resultCh {
		fmt.Println(result.Address, result.Balance)
	}

}

func QueryBalance(address string) int {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return rand.Intn(10000)
}
