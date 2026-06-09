# Web3 后端 + 智能合约短期复习行动手册

> 目标：用最短时间把之前学习过的 Web3 后端 + 智能合约内容重新串起来，恢复到能做项目、能看懂岗位 JD、能准备面试的状态。  
> 推荐方向：**Go 后端 + 合约辅助**，而不是纯智能合约开发。

---

## 1. 复习总目标

这次复习不追求“把所有资料重新学一遍”，而是以项目驱动的方式，把核心能力快速恢复：

- 理解区块链基础原理
- 恢复 Go 后端开发手感
- 掌握 Gin / GORM 常用后端开发能力
- 使用 go-ethereum / geth SDK 与链交互
- 恢复 Solidity 合约开发能力
- 掌握 Foundry / Hardhat 基础脚手架
- 最终整合出一个可写进简历的 Web3 后端项目

---

## 2. 学习路线总览

整体路线可以压缩成 6 个模块：

```text
区块链基础
  ↓
Go 基础
  ↓
Gin / GORM 后端开发
  ↓
geth SDK / go-ethereum 链上交互
  ↓
Solidity 智能合约
  ↓
Foundry / Hardhat 项目实战
```

对应能力目标：

| 模块 | 目标 |
|---|---|
| 区块链基础 | 能讲清 BTC、Ethereum、账户模型、交易、Gas、共识、L1/L2 |
| Go 基础 | 能写结构体、接口、goroutine、channel、错误处理、项目结构 |
| Gin / GORM | 能写 REST API，连数据库，完成 CRUD |
| geth SDK | 能用 Go 查余额、查区块、查交易、监听事件、调用合约 |
| Solidity | 能写 ERC20、简单 NFT、质押、投票、权限控制等基础合约 |
| Foundry / Hardhat | 能测试、部署、调试合约，并和后端联动 |

---

## 3. 14 天极速复习计划

### 第 1-2 天：区块链基础复活

#### 目标

能用自己的话讲清楚 Ethereum 是怎么运行起来的。

#### 必须掌握

- 区块、交易、区块哈希
- Merkle Tree
- 公钥、私钥、地址、签名
- BTC 的 UTXO 模型
- Ethereum 的账户模型
- EOA / Contract Account
- Gas、Nonce、交易生命周期
- PoS、Validator、区块确认
- L1 / L2、Rollup、跨链桥的基本概念

#### 产出物

写一篇 Markdown：

```text
notes/01-blockchain-basic.md
```

标题：

```md
# 以太坊交易流程复习
```

内容模板：

```md
## 1. 用户发起交易

## 2. 钱包签名

## 3. 交易广播到节点

## 4. 节点验证交易

## 5. 交易进入区块

## 6. EVM 执行交易

## 7. 状态更新

## 8. 区块确认

## 9. 后端如何查询这笔交易
```

#### 验收问题

你要能回答：

> 用户点击“转账 1 ETH”后，从钱包签名到链上确认，中间发生了什么？

---

### 第 3-4 天：Go 基础快速唤醒

#### 目标

恢复 Go 的基础语法和工程手感。

#### 必须复习

```go
struct
interface
error
defer
go func()
channel
select
context
go.mod
testing
```

#### 小练习

写一个命令行小工具：

> 输入钱包地址，校验格式，并打印查询参数。

#### 验收命令

```bash
go version
go test ./...
go run main.go
```

只要这三个能顺畅跑，就说明 Go 环境和基础手感恢复了。

---

### 第 5-6 天：Gin + GORM 后端恢复

#### 目标

恢复传统后端能力，这是 Web3 后端的兜底能力。

#### 小项目

项目名：

```text
Wallet Notes API
```

功能：钱包地址备注系统。

#### 接口设计

| 接口 | 功能 |
|---|---|
| POST /wallets | 添加钱包地址 |
| GET /wallets | 查询钱包列表 |
| GET /wallets/:address | 查询单个钱包 |
| PUT /wallets/:address | 修改备注 |
| DELETE /wallets/:address | 删除钱包 |

#### 数据库表设计

```text
wallets
- id
- address
- nickname
- chain
- created_at
- updated_at
```

#### 重点恢复能力

- 路由
- 中间件
- 参数绑定
- 错误返回
- 数据库模型
- CRUD
- 项目分层

---

### 第 7-8 天：geth SDK / go-ethereum 核心能力

#### 目标

让 Go 后端具备和 Ethereum 链交互的能力。

#### 必须掌握

- 连接 RPC
- 查询区块高度
- 查询账户余额
- 查询交易详情
- 查询 ERC20 余额
- 监听合约事件
- 构造并签名交易
- 使用 abigen 生成合约绑定代码

#### 小项目升级

在 `Wallet Notes API` 中增加接口：

```http
GET /wallets/:address/balance
```

返回示例：

```json
{
  "address": "0x...",
  "eth_balance": "1.23",
  "block_number": 12345678
}
```

#### 验收标准

完成后，你的项目就从普通 Go 后端升级为 Web3 后端。

---

### 第 9-10 天：Solidity 基础恢复

#### 目标

能写、能测、能部署基础合约。

#### 必须掌握

```solidity
contract
state variable
mapping
struct
modifier
event
error
require / revert
constructor
visibility
payable
receive / fallback
```

#### 第一份合约

合约名：

```text
WalletRegistry
```

功能：

- 绑定钱包地址
- 设置昵称
- 修改昵称
- 发出事件
- 只有 owner 可以管理某些配置

#### 验收问题

你要能讲清楚：

- storage / memory / calldata 的区别
- public / external / internal / private 的区别
- event 为什么重要
- modifier 用来干嘛
- msg.sender 是谁
- tx.origin 为什么危险

---

### 第 11-12 天：Foundry / Hardhat 脚手架

#### 建议

主用 **Foundry**，Hardhat 了解即可。

Foundry 更适合快速写合约、跑测试、做命令行交互，对工程师比较友好。

#### 必须掌握命令

```bash
forge init
forge build
forge test
forge script
cast call
cast send
anvil
```

#### 小项目继续升级

- 用 Foundry 写 `WalletRegistry` 合约
- 写测试
- 本地 `anvil` 部署
- 用 Go 后端调用这个合约
- 后端监听合约 event，写入数据库

这一步完成后，就形成了完整闭环：

```text
Solidity 合约
  ↓
Foundry 测试部署
  ↓
Go 后端调用合约
  ↓
监听事件
  ↓
写入数据库
  ↓
REST API 对外提供服务
```

---

### 第 13-14 天：整合成一个作品

#### 最终项目名

```text
Web3 Wallet Tracker
```

#### 功能清单

- 添加钱包地址
- 查询钱包 ETH 余额
- 查询 ERC20 余额
- 给钱包添加备注
- 部署一个 `WalletRegistry` 合约
- Go 后端读取链上合约数据
- 监听合约事件
- 提供 REST API
- 编写 README
- 整理成简历项目

#### 技术栈

```text
Go
Gin
GORM
Ethereum RPC
go-ethereum
Solidity
Foundry
MySQL / SQLite
REST API
```

#### 简历描述示例

> 基于 Go + Gin + GORM + go-ethereum + Solidity + Foundry 实现 Web3 钱包追踪服务，支持钱包地址管理、链上余额查询、合约事件监听、ERC20 资产读取与本地数据同步。负责后端接口设计、链上 RPC 交互、智能合约开发与测试部署。

---

## 4. 推荐项目目录

建议建立一个总仓库：

```text
web3-return/
  notes/
    01-blockchain-basic.md
    02-go-review.md
    03-gin-gorm.md
    04-geth-sdk.md
    05-solidity.md
    06-foundry.md
  backend/
    cmd/
    internal/
    pkg/
    go.mod
    README.md
  contracts/
    src/
    test/
    script/
    foundry.toml
    README.md
  README.md
```

---

## 5. 每日固定学习循环

每天建议按下面节奏推进：

```text
30 分钟：复习概念
90 分钟：写代码
30 分钟：整理笔记
30 分钟：复盘 + 面试表达
```

重点提醒：

> 不要一天 6 小时都看资料。Web3 学习很容易掉进资料黑洞。  
> 最有效的方法是：一边复习，一边写代码，一边整理成自己的表达。

---

## 6. 当前第一步任务

今天只做 3 件事。

### 任务 1：环境恢复

确认这些命令：

```bash
go version
git --version
node -v
npm -v
```

如果准备使用 Foundry，继续确认：

```bash
forge --version
cast --version
anvil --version
```

---

### 任务 2：建立复习项目目录

执行：

```bash
mkdir web3-return
cd web3-return

mkdir notes backend contracts
touch README.md

touch notes/01-blockchain-basic.md
touch notes/02-go-review.md
touch notes/03-gin-gorm.md
touch notes/04-geth-sdk.md
touch notes/05-solidity.md
touch notes/06-foundry.md
```

Windows PowerShell 可以用：

```powershell
mkdir web3-return
cd web3-return

mkdir notes, backend, contracts
New-Item README.md
New-Item notes/01-blockchain-basic.md
New-Item notes/02-go-review.md
New-Item notes/03-gin-gorm.md
New-Item notes/04-geth-sdk.md
New-Item notes/05-solidity.md
New-Item notes/06-foundry.md
```

---

### 任务 3：写第一篇笔记

文件：

```text
notes/01-blockchain-basic.md
```

标题：

```md
# 以太坊交易流程复习
```

先按下面模板写粗稿：

```md
## 1. 用户发起交易

## 2. 钱包签名

## 3. 交易广播到节点

## 4. 节点验证交易

## 5. 交易进入区块

## 6. EVM 执行交易

## 7. 状态更新

## 8. 区块确认

## 9. 后端如何查询这笔交易
```

---

## 7. 后续推进方式

后续可以按“关卡制”推进：

```text
第 1 关：区块链基础复活
第 2 关：Go 后端小项目
第 3 关：geth SDK 接链
第 4 关：Solidity 合约
第 5 关：Foundry 测试部署
第 6 关：整合成简历项目
```

每一关都必须有：

- 一份笔记
- 一段代码
- 一个验收问题
- 一段面试表达

---

## 8. 核心原则

### 原则 1：不重新从零学

你之前已经学过，现在目标是恢复，而不是重开。

### 原则 2：项目优先

只看资料会忘，写项目才会重新长回来。

### 原则 3：后端兜底，合约加分

你的主线是：

```text
Go 后端工程师
+
Web3 链上交互能力
+
Solidity 合约辅助开发能力
```

这比单纯卷 Solidity 更稳。

### 原则 4：每天都要形成可沉淀内容

每天至少留下：

```text
1 篇笔记
或
1 个接口
或
1 个合约
或
1 个测试
或
1 段面试表达
```

---

## 9. 最终目标状态

完成这轮短期复习后，你应该能做到：

- 看懂 Web3 后端 JD
- 写一个 Go 后端服务
- 使用 Gin / GORM 做基础业务
- 用 go-ethereum 查询链上数据
- 调用 ERC20 合约
- 写基础 Solidity 合约
- 用 Foundry 测试和部署
- 解释链上交易流程
- 把项目写进简历
- 进入 Web3 后端面试准备阶段

---

## 10. 下一步

建议下一步从 `第 1 关：区块链基础复活` 开始。

优先完成：

```text
notes/01-blockchain-basic.md
```

核心问题：

> 用户点击“转账 1 ETH”后，从钱包签名到链上确认，中间发生了什么？

只要这个问题能讲清楚，Web3 的主干就重新立起来了。
