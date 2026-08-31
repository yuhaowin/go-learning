package bank_test

import (
	"fmt"
	"testing"

	bank "github.com/yuhaowin/go-learning/ch09/bank/version1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
		// struct{}{} 拆开看是两部分：
		// - struct{} —— 一个空结构体类型（没有任何字段的匿名结构体）
		// - {} —— 用这个类型构造出的一个值（零值实例）
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
