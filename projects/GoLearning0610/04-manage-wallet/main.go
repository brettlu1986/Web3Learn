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
		Balance: big.NewInt(0), //new(big.Int).set(0)
	}
}

// 存
func (w *Wallet) Deposit(amount *big.Int) {
	w.Balance = new(big.Int).Add(w.Balance, amount)
}

// 取
func (w *Wallet) Withdraw(amount *big.Int) error {
	if w.Balance.Cmp(amount) < 0 {
		return fmt.Errorf("insufficient ballance")
	}

	w.Balance = new(big.Int).Sub(w.Balance, amount)
	return nil
}

// 余额
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
	}

	fmt.Println("address:", wallet.Address)
	fmt.Println("balance:", wallet.GetBalance())
}
