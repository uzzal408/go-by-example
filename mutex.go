package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var balance int
var wg sync.WaitGroup

func deposite(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Deposite amount %d to account with balance %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()

}

func callMutex() {
	balance = 100
	wg.Add(2)
	deposite(700, &wg)
	withdraw(500, &wg)
	wg.Wait()

	fmt.Printf("New Balance is: %d\n", balance)
}
