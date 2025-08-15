package main

import "fmt"

var deposits = make( chan int )
var balances = make( chan int )

var balance int

func Deposit( amount int ) { balance = balance + amount }
func Balance() int { return balance }

func main() {
	fmt.Println("Main Function Called...")

	thakurDeposit := func() {
		Deposit( 200 )
		fmt.Println("Balance =", Balance() )
	}

 	gabbarDeposit := func() {
		Deposit( 100 )
	}

	go thakurDeposit()
	go gabbarDeposit()

	if got, want := Balance(), 300 ; got != want {
		fmt.Printf("\nBalance = %d, Want = %d", got, want)
	}
}

// Race Conditions : Data Race
