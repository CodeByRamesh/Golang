package main

import "fmt"

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

func main() {
	fmt.Println("Main Function Called...")
	done := make( chan bool ) 

	thakurDeposit := func() {
		Deposit( 200 )
		fmt.Println("Balance =", Balance() )
		done <- true // Marking Work Completion
	}

 	gabbarDeposit := func() {
		Deposit( 100 )
		done <- true // Marking Work Completion
	}

	go thakurDeposit()
	go gabbarDeposit()

	<- done
	<- done 

	// Checking Balance Only After Confirmation Of Work Completion
	if got, want := Balance(), 300 ; got != want {
		fmt.Printf("\nBalance = %d, Want = %d", got, want)
	}

	fmt.Println("Main Function Exiting...")
}

// Race Conditions : Data Race
