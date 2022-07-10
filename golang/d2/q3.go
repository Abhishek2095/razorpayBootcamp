package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct {
	mu      sync.Mutex
	balance int
}

func (acc *Account) deposit(depositAmt int) {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	acc.balance += depositAmt

	fmt.Println("Depositing Amount : ", depositAmt)
	fmt.Println("Current Account Balance ", acc.getBalance())
}

func (acc *Account) withdraw(withdrawAmt int) {

	acc.mu.Lock()
	defer acc.mu.Unlock()

	if acc.balance < withdrawAmt {
		panic("Not sufficient amount to withdraw")
	} else {
		acc.balance -= withdrawAmt
	}
	fmt.Println("Withdrawing Amount : ", withdrawAmt)
	fmt.Println("Current Account Balance ", acc.getBalance())
}

func (acc *Account) getBalance() int {
	return acc.balance
}

func main() {
	rand.Seed(time.Now().UnixNano())

	acc := Account{
		balance: 5000,
	}
	var wg sync.WaitGroup

	iterations := 2
	for i := 0; i < iterations; i++ {
		wg.Add(1)

		if i%2 == 0 {
			depositAmt := rand.Intn(1000)
			go func() {
				defer wg.Done()
				acc.deposit(depositAmt)
			}()
		} else {
			withdrawAmt := rand.Intn(1000)
			go func() {
				defer wg.Done()
				acc.withdraw(withdrawAmt)
			}()
		}

	}

	wg.Wait()
	fmt.Println("Current Account Balance ", acc.getBalance())
}
