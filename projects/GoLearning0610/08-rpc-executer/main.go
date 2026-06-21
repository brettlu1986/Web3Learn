package main

import (
	"context"
	"fmt"
	"time"
)

// 1. context.Context：上下文，用来传递取消信号和超时时间
// 2. context.WithTimeout：创建一个带超时的 context
// 3. select：同时等待多个 channel，谁先来就执行谁
// 4. ctx.Done()：context 被取消或超时时会关闭的 channel
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	result, err := SlowRPC(ctx)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("result:", result)
}

// 执行过程
// 0 秒：main 创建 2 秒超时的 ctx
// 0 秒：调用 SlowRPC(ctx)
// 0 秒：SlowRPC 创建 resultCh
// 0 秒：SlowRPC 启动 goroutine，开始模拟 3 秒 RPC
// 0 秒：SlowRPC 进入 select
// 0-2 秒：select 同时等待 resultCh 和 ctx.Done()
// 2 秒：ctx 超时，ctx.Done() 触发
// 2 秒：select 进入 ctx.Done() 分支
// 2 秒：返回 "", context deadline exceeded
// 2 秒：main 打印 error
// 3 秒：模拟 RPC 的 goroutine 醒来，尝试往 resultCh 发送结果
// 3 秒：因为 resultCh 有 1 个缓冲位，所以发送成功，然后 goroutine 结束
func SlowRPC(ctx context.Context) (string, error) {
	resultCh := make(chan string, 1) //代表一个有缓冲，可以暂存一个结果的channel， 类型是string
	//模拟一个慢的rpc任务，context只等2秒，这个任务3秒肯定超时，同时为了保证select能够监听到
	go func() {
		time.Sleep(3 * time.Second)
		resultCh <- "rpc result"
	}()

	//同时等待多个channel，哪个先返回就执行哪个
	select {
	case result := <-resultCh:
		return result, nil
	case <-ctx.Done(): //代表context先超时了 ctx.Done() 是一个 channel，context 超时或取消时会触发。
		return "", ctx.Err()
	}
}
