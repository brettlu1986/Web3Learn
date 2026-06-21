package main

import "fmt"

type BalanceReader interface { //定义的是一种能力
	GetBalance(address string) (int, error)
}

type EthereumClient struct{} //真实rpc

func (c *EthereumClient) GetBalance(address string) (int, error) { //隐式实现interface，不需要implement
	fmt.Println("query balance from ethereum rpc:", address)
	return 1000, nil
}

type MockClient struct{} //测试假数据

func (m *MockClient) GetBalance(address string) (int, error) {
	fmt.Println("query from mock:", address)
	return 123, nil
}

type BalanceService struct {
	reader BalanceReader //凡是 实现了 GetBalance接口的 都可以传进来
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
	fmt.Println("balance : ", balance)
	return nil
}

func main() {
	reader := &EthereumClient{}
	service := NewBalanceService(reader) // or NewBalanceService(&MockClient)

	err := service.PrintBalance("0x123")
	if err != nil {
		fmt.Println("error:", err)
	}
}
