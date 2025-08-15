
package main

import "fmt"

var deposits = make( chan int )
var balances = make( chan int )

func Deposit( amount int ) { deposits <- amount }
func Balance() int { return <-balances }

func bankTeller() {
	var balance int 
	
	for {
		select {
		case amount := <- deposits :
				balance = balance + amount
		case balances <- balance:
		}
	}
}

func init() { // To Initialse Package Level Things
	fmt.Println("Init Function Called...")
	// var something = struct{}{}
	// fmt.Println( something )
	go bankTeller()
}

func main() {
	fmt.Println("Main Function Called...")
	done := make( chan struct{} ) 

	thakurDeposit := func() {
		Deposit( 200 )
		fmt.Println("Balance =", Balance() )
		done <- struct{}{}
	}

 	gabbarDeposit := func() {
		Deposit( 100 )
		done <- struct{}{}
	}

	go thakurDeposit()
	go gabbarDeposit()

	<- done
	<- done

	if got, want := Balance(), 300 ; got != want {
		fmt.Println("Balance = %d, Want = %d", got, want)
	}
}

// Race Conditions : Data Race

