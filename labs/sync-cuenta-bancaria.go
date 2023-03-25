package labs

import (
	"fmt"
	"sync"
)

var (
	balance int = 0
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()

	lock.Lock()
	b := balance
	balance = b + amount
	fmt.Println("Deposit:         $", amount)
	lock.Unlock()
}

func Withdrawal(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()

	lock.Lock()
	b := balance

	if balance >= amount {
		balance = b - amount
		fmt.Printf("Withdrawal:      $ -%d\n", amount)
	} else {
		fmt.Println("Saldo insuficiente")
	}

	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()

	return b
}

func StartSyncCuentaBancaria() {
	var wg sync.WaitGroup
	var lock sync.RWMutex

	fmt.Println("Saldo al inicio: $", Balance(&lock))

	wg.Add(1)
	go Deposit(100, &wg, &lock)

	wg.Add(1)
	go Deposit(300, &wg, &lock)

	wg.Add(1)
	go Deposit(200, &wg, &lock)

	wg.Add(1)
	go Withdrawal(1000, &wg, &lock)

	wg.Add(1)
	go Withdrawal(10, &wg, &lock)

	wg.Wait()
	println("Saldo actual:    $", Balance(&lock))
}
