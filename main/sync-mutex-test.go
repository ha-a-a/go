package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享变量的并发
// 竞争条件：程序在多个goroutine交叉执行时，没有给出正确的结果
// 避免竞争： 1不要对共享变量进行写操作；
//			2避免从多个goroutine，即不要使用变量通信，使用通信来共享变量；
//			3多个goroutine访问变量时，同一时刻最多只有一个goroutine在访问变量，sync.Mutex互斥锁,sync.RWMutex读写锁

// sync.Mutex 互斥锁
var (
	mu      = sync.Mutex{}
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)

}
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	b := balance
	fmt.Printf("balance now：%d \n", b)
	return b
}

// 同一个goroutine中不能sync.Mutex不能重入,导致死锁
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}
func deposit(amount int) {
	balance = balance + amount
	fmt.Printf("after deposit：%d \n", balance)
}
func main() {
	//go Deposit(100)
	//go Balance()
	isSuc := Withdraw(100)
	fmt.Printf("balance is enough, %t \n", isSuc)
	time.Sleep(10 * time.Second)
}
