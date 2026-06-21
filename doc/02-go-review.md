# Go 基础练习总表

|编号	|程序	|重点知识点
|---|---|---|
|01	|命令行计算器	|变量、函数、参数、错误处理
|02	|钱包地址校验器	|string、error、测试
|03	|学生成绩统计	|slice、map、struct
|04	|钱包资产管理器	|struct、method、指针
|05	|JSON 配置读取	|struct tag、JSON、文件读取
|06	|日志解析器	|文件读取、字符串处理、map 统计
|07	|并发余额查询模拟器	|goroutine、channel、WaitGroup
|08	|带超时的任务执行器	|context、select、timeout
|09	|接口版余额查询服务	|interface、依赖注入、mock
|10	|小型 CLI 项目结构	|package、go.mod、工程分层、testing

## 练习 01：命令行计算器
### 目标

熟悉：
```
变量
函数
命令行参数
strconv
error
switch
```

### 功能

运行：
`go run . add 3 5`

输出：
`result: 8`

支持：
```
add
sub
mul
div
```

###示例代码

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("usage: go run . <add|sub|mul|div> <a> <b>")
		return
	}

	op := os.Args[1]

	a, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("invalid number a:", err)
		return
	}

	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("invalid number b:", err)
		return
	}

	result, err := Calculate(op, a, b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("result:", result)
}

func Calculate(op string, a int, b int) (int, error) {
	switch op {
	case "add":
		return a + b, nil
	case "sub":
		return a - b, nil
	case "mul":
		return a * b, nil
	case "div":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", op)
	}
}
```

### 你要掌握

这段是 Go 最经典的错误处理模式：

```go
result, err := Calculate(op, a, b)
if err != nil {
    return
}
```

## 练习 02：钱包地址校验器
这个和 Web3 直接相关

### 目标

熟悉：
```
string
len
strings.HasPrefix
error
testing
```

### 功能

校验以太坊地址：
```
必须以 0x 开头
长度必须是 42
0x 后面应该是 40 个十六进制字符
```

### 项目结构

```
wallet-checker/
  go.mod
  main.go
  wallet/
    validator.go
    validator_test.go
```

### validator.go

```go
package wallet

import (
	"fmt"
	"strings"
)

func ValidateAddress(address string) error {
	if address == "" {
		return fmt.Errorf("address is empty")
	}

	if !strings.HasPrefix(address, "0x") {
		return fmt.Errorf("address must start with 0x")
	}

	if len(address) != 42 {
		return fmt.Errorf("address length must be 42")
	}

	for _, ch := range address[2:] {
		if !isHexChar(ch) {
			return fmt.Errorf("address contains non-hex character: %c", ch)
		}
	}

	return nil
}

func isHexChar(ch rune) bool {
	return ch >= '0' && ch <= '9' ||
		ch >= 'a' && ch <= 'f' ||
		ch >= 'A' && ch <= 'F'
}
```

### validator_test.go

```go
package wallet

import "testing"

func TestValidateAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		wantErr bool
	}{
		{
			name:    "valid address",
			address: "0x1234567890123456789012345678901234567890",
			wantErr: false,
		},
		{
			name:    "empty address",
			address: "",
			wantErr: true,
		},
		{
			name:    "missing 0x",
			address: "1234567890123456789012345678901234567890",
			wantErr: true,
		},
		{
			name:    "wrong length",
			address: "0x123",
			wantErr: true,
		},
		{
			name:    "invalid hex char",
			address: "0x123456789012345678901234567890123456789Z",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAddress(tt.address)

			if tt.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("expected nil, got %v", err)
			}
		})
	}
}
```

### 你要掌握

这里有几个重点：

```go
for _, ch := range address[2:]
```

ch 是 rune。

address[2:] 是字符串切片，表示去掉前面的 0x。


## 练习 03：学生成绩统计

这个练习不是 Web3，但非常适合练 `slice + map + struct`

### 目标

熟悉：
```
struct
slice
map
for range
函数拆分
功能
```

统计学生成绩：
```
平均分
最高分
最低分
按学生名查询成绩
```

### 示例代码

```go
package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func main() {
	students := []Student{
		{Name: "Alice", Score: 90},
		{Name: "Bob", Score: 78},
		{Name: "Cindy", Score: 95},
		{Name: "David", Score: 66},
	}

	avg := AverageScore(students)
	max := MaxScore(students)
	min := MinScore(students)

	fmt.Println("average:", avg)
	fmt.Println("max:", max)
	fmt.Println("min:", min)

	scoreMap := BuildScoreMap(students)
	fmt.Println("Bob score:", scoreMap["Bob"])
}

func AverageScore(students []Student) float64 {
	if len(students) == 0 {
		return 0
	}

	total := 0
	for _, s := range students {
		total += s.Score
	}

	return float64(total) / float64(len(students))
}

func MaxScore(students []Student) int {
	if len(students) == 0 {
		return 0
	}

	max := students[0].Score
	for _, s := range students {
		if s.Score > max {
			max = s.Score
		}
	}

	return max
}

func MinScore(students []Student) int {
	if len(students) == 0 {
		return 0
	}

	min := students[0].Score
	for _, s := range students {
		if s.Score < min {
			min = s.Score
		}
	}

	return min
}

func BuildScoreMap(students []Student) map[string]int {
	result := make(map[string]int)

	for _, s := range students {
		result[s.Name] = s.Score
	}

	return result
}
```

### 你要掌握
```go
students := []Student{}
```
这是 slice。

```go
map[string]int
```

表示 key 是 string，value 是 int。

后面 Web3 里会经常用：
```go
map[string]*big.Int
map[string]Wallet
```

## 练习 04：钱包资产管理器

这个开始贴近 Web3。

### 目标

熟悉：
```
struct
method
指针接收者
*big.Int
```

### 功能

管理一个钱包资产：
```
创建钱包
充值
扣款
查询余额
```

### 示例代码
```go
package main

import (
	"fmt"
	"math/big"
)

type Wallet struct {
	Address string
	Balance *big.Int
}

func NewWallet(address string) *Wallet {
	return &Wallet{
		Address: address,
		Balance: big.NewInt(0),
	}
}

func (w *Wallet) Deposit(amount *big.Int) {
	w.Balance = new(big.Int).Add(w.Balance, amount)
}

func (w *Wallet) Withdraw(amount *big.Int) error {
	if w.Balance.Cmp(amount) < 0 {
		return fmt.Errorf("insufficient balance")
	}

	w.Balance = new(big.Int).Sub(w.Balance, amount)
	return nil
}

func (w *Wallet) GetBalance() *big.Int {
	return new(big.Int).Set(w.Balance)
}

func main() {
	wallet := NewWallet("0x1234567890123456789012345678901234567890")

	wallet.Deposit(big.NewInt(1000))
	wallet.Deposit(big.NewInt(500))

	err := wallet.Withdraw(big.NewInt(300))
	if err != nil {
		fmt.Println("withdraw failed:", err)
		return
	}

	fmt.Println("address:", wallet.Address)
	fmt.Println("balance:", wallet.GetBalance())
}
```

### 你要掌握

为什么这里用：

```go
func (w *Wallet) Deposit(amount *big.Int)
```

而不是：
```go
func (w Wallet) Deposit(amount *big.Int)
```

因为我们要修改 wallet 的余额，所以需要指针接收者。

另外：
```go
new(big.Int).Add(w.Balance, amount)
```

不要直接乱改原对象。big.Int 是可变对象，复制时要小心。

## 练习 05：JSON 配置读取

以后你做后端项目，经常要读配置。

### 目标

熟悉：
```go
struct tag
encoding/json
os.ReadFile
```

### 配置结构

```go
config.json
{
  "rpc_url": "https://eth-mainnet.g.alchemy.com/v2/your-key",
  "database_url": "sqlite://wallet.db",
  "port": 8080
}
```

### 示例代码

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	RPCURL      string `json:"rpc_url"`
	DatabaseURL string `json:"database_url"`
	Port       int    `json:"port"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func main() {
	cfg, err := LoadConfig("config.json")
	if err != nil {
		fmt.Println("load config failed:", err)
		return
	}

	fmt.Println("rpc:", cfg.RPCURL)
	fmt.Println("db:", cfg.DatabaseURL)
	fmt.Println("port:", cfg.Port)
}
``` 

### 你要掌握

这段：
```go
RPCURL string `json:"rpc_url"`
```

是 struct tag。

它告诉 JSON 解析器：
```
JSON 里的 rpc_url 对应 Go 里的 RPCURL 字段
```


## 练习 06：日志解析器

### 目标

熟悉：
```
文件读取
strings
map 统计
错误处理
```

### 假设有 access.log
```
INFO wallet created
ERROR rpc timeout
INFO balance queried
ERROR invalid address
ERROR rpc timeout
```

### 目标输出
```
INFO: 2
ERROR: 3
rpc timeout: 2
```

### 示例代码
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stats, err := ParseLog("access.log")
	if err != nil {
		fmt.Println("parse log failed:", err)
		return
	}

	for key, count := range stats {
		fmt.Printf("%s: %d\n", key, count)
	}
}

func ParseLog(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "INFO") {
			stats["INFO"]++
		}

		if strings.HasPrefix(line, "ERROR") {
			stats["ERROR"]++
		}

		if strings.Contains(line, "rpc timeout") {
			stats["rpc timeout"]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
```

你要掌握
```go
defer file.Close()
```

表示函数结束前关闭文件。

`defer` 在资源释放里很常见：
```go
defer rows.Close()
defer cancel()
defer file.Close()
```


## 练习 07：并发余额查询模拟器

这个很重要，Go 的并发从这里开始。

### 目标

熟悉：
```
goroutine
channel
sync.WaitGroup
并发任务
```

### 功能

模拟同时查询多个钱包余额。

### 示例代码
```go
package main

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

func main() {
	addresses := []string{
		"0x111",
		"0x222",
		"0x333",
		"0x444",
	}

	resultCh := make(chan BalanceResult)

	var wg sync.WaitGroup

	for _, address := range addresses {
		wg.Add(1)

		go func(addr string) {
			defer wg.Done()

			balance := QueryBalance(addr)
			resultCh <- BalanceResult{
				Address: addr,
				Balance: balance,
			}
		}(address)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		fmt.Println(result.Address, result.Balance)
	}
}

func QueryBalance(address string) int {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return rand.Intn(10000)
}
```

你要掌握

这个结构很经典：

```go
var wg sync.WaitGroup

wg.Add(1)
go func() {
    defer wg.Done()
}()

wg.Wait()
```

还有：
```go
for result := range resultCh
```
表示不断从 channel 读取，直到 channel 被关闭。

这个模型以后可以用在：
```
批量查询钱包余额
批量查交易 receipt
批量请求外部 RPC
```

## 练习 08：带超时的任务执行器

### 目标

熟悉：
```
context
timeout
select
ctx.Done()
```

### 功能

模拟 RPC 调用，超过 2 秒就取消。

### 示例代码
```go
package main

import (
	"context"
	"fmt"
	"time"
)

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

func SlowRPC(ctx context.Context) (string, error) {
	resultCh := make(chan string, 1)

	go func() {
		time.Sleep(3 * time.Second)
		resultCh <- "rpc result"
	}()

	select {
	case result := <-resultCh:
		return result, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
```

### 你要掌握

这段是核心：
```go
select {
case result := <-resultCh:
    return result, nil
case <-ctx.Done():
    return "", ctx.Err()
}
```

Web3 后端调用 RPC 时，必须考虑：
```
网络卡住
RPC 服务商超时
节点不响应
请求被取消
```
所以 `context` 很重要。


## 练习 09：接口版余额查询服务

这是工程化关键。

### 目标

熟悉：
```
interface
依赖注入
mock
解耦
```

### 示例代码
```go
package main

import "fmt"

type BalanceReader interface {
	GetBalance(address string) (int, error)
}

type EthereumClient struct{}

func (c *EthereumClient) GetBalance(address string) (int, error) {
	return 1000, nil
}

type MockBalanceReader struct{}

func (m *MockBalanceReader) GetBalance(address string) (int, error) {
	return 123, nil
}

type BalanceService struct {
	reader BalanceReader
}

func NewBalanceService(reader BalanceReader) *BalanceService {
	return &BalanceService{
		reader: reader,
	}
}

func (s *BalanceService) PrintBalance(address string) error {
	balance, err := s.reader.GetBalance(address)
	if err != nil {
		return err
	}

	fmt.Println("balance:", balance)
	return nil
}

func main() {
	reader := &EthereumClient{}
	service := NewBalanceService(reader)

	err := service.PrintBalance("0x123")
	if err != nil {
		fmt.Println("error:", err)
	}
}
```

### 你要掌握

`BalanceService` 不关心你到底是：
```
真实 Ethereum RPC
还是测试 Mock
还是数据库缓存
```

它只关心你有没有实现：
```go
GetBalance(address string) (int, error)
```
这就是 interface 的价值。

## 练习 10：小型 CLI 项目结构

最后把前面的东西整理成一个小项目。

### 项目名
go-review-cli
### 项目结构
```
go-review-cli/
  go.mod
  main.go
  internal/
    wallet/
      validator.go
      validator_test.go
    balance/
      service.go
      service_test.go
    config/
      config.go
```
### 功能

支持命令：
```Bash
go run . validate 0x1234567890123456789012345678901234567890
go run . balance 0x1234567890123456789012345678901234567890
``` 
### main.go 简化版
```go
package main

import (
	"fmt"
	"os"

	"go-review-cli/internal/wallet"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: go run . <validate|balance> <address>")
		return
	}

	command := os.Args[1]
	address := os.Args[2]

	switch command {
	case "validate":
		if err := wallet.ValidateAddress(address); err != nil {
			fmt.Println("invalid address:", err)
			return
		}
		fmt.Println("valid address:", address)

	case "balance":
		fmt.Println("mock balance:", address, "1000 wei")

	default:
		fmt.Println("unknown command:", command)
	}
}
```

这个项目做完，你 Go 的基本工程感觉就回来了。


## 小结
### 第 1 组：语法恢复
```
01 命令行计算器
02 钱包地址校验器
03 学生成绩统计
```
目标：恢复函数、slice、map、error、testing。

### 第 2 组：Web3 相关基础
```
04 钱包资产管理器
05 JSON 配置读取
06 日志解析器
```
目标：恢复 struct、method、指针、big.Int、文件处理。

### 第 3 组：Go 后端核心
```
07 并发余额查询模拟器
08 带超时任务执行器
09 接口版余额查询服务
```

目标：恢复 goroutine、channel、context、interface。

### 第 4 组：工程整合
```
10 小型 CLI 项目结构
```

目标：恢复 package、go.mod、internal、测试和工程组织。

## 今天验收题

做完后，你回答我这几个：

### Q1

`func Calculate(op string, a int, b int) (int, error)` 为什么要返回两个值？

### Q2

为什么地址校验函数返回 `error`，而不是直接返回 `bool`？

### Q3

`func (w *Wallet) Deposit(amount *big.Int)` 里为什么接收者是 *Wallet？

### Q4

`new(big.Int).Add(a, b)` 这句话是什么意思？

### Q5

`b := a` 和 `b := new(big.Int).Set(a)` 在 `*big.Int` 里有什么区别？

### Q6

为什么 `Go` 项目里要把代码拆成不同 `package`，而不是全写在 `main.go`