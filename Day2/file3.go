package main

import (
	"fmt"
	"math/rand"
	"sync"
)
// Money Deposit
func deposit(bal *int, add int) {
	*bal = *bal + add
}
// Money Withdrawl
func Withdrawl(bal *int, sub int) {
	if *bal < 0 {
		fmt.Println("Less that 0 account balance")
	} else {
		*bal = *bal - sub
	}
}
// Use of mutex
func main() {
	fmt.Println("Question 3")
	bal := 500
	var vs sync.Mutex
	for i := 0; i < 10; i++ {
		sub := rand.Intn(100)
		add := rand.Intn(100)
		vs.Lock()
		go deposit(&bal, add)
		vs.Unlock()
		vs.Lock()
		go Withdrawl(&bal, sub)
		vs.Unlock()
		fmt.Println("Remaining Balance", bal)
	}
	fmt.Println("Final Balance", bal)
}

