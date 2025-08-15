package main

import (
	"fmt"
	"sync"
)


var deposits = make( chan int )
var balances = make( chan int )

func Deposit( amount int ) { deposits <- amount }
func Balance() int { return <- balances }

func bankTeller() {
	var balance int // balance is Comfined To One Goroutine
	
	for {
		select {
		case amount := <- deposits:
			// Balance Getting Updated At Only One Place
			balance = balance + amount
			// Balance Reading At Only One Place
		case balances <- balance:
		}
	}
}

func init() {
	fmt.Println("Init Function Called...")

	// Monitor Go Routine 
	//		Monitors Status Of balance
	go bankTeller() 
}

var wg sync.WaitGroup

func depositAmount(amount int ) {
	Deposit( amount )
	// fmt.Println("Balance =", Balance() )
	wg.Done()
}

func main() {
	fmt.Println("Main Function Called...")

	for amount := 10 ; amount <= 20 ; amount++ {
		wg.Add(1)
		go depositAmount( amount )		
	}

	wg.Wait()

	// Checking Balance Only After Confirmation Of Work Completion
	if got, want := Balance(), 300 ; got != want {
		fmt.Printf("\nBalance = %d, Want = %d", got, want)
	}

	fmt.Println("Main Function Exiting...")
}

//_______________________________________________________

// var (
// 	mutex      sync.Mutex // guards balance
// 	balance int
// )

// func Deposit(amount int) {
// 	mutex.Lock()
// 	balance = balance + amount
// 	mutex.Unlock()
// }

// func Balance() int {
// 	mutex.Lock()
// 	b := balance
// 	mutex.Unlock()
// 	return b
// }

//_______________________________________________________

