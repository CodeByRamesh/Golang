package main

import "fmt"

var deposits = make( chan int )
var balances = make( chan int )

var balance int

func Deposit( amount int ) { balance = balance + amount }
func Balance() int { return balance }

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
