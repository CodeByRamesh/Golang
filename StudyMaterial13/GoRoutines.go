
package main 

import (
	"fmt"
	"time"
)

//_____________________________________________________________________

func spinner( delay time.Duration ) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep( delay )
		}
	}
}

func fib( x int ) int {
	if x < 2 {
		return x 
	}
	return fib( x - 1 ) + fib( x - 2 )
}

func playWithFibonacci() {
	// spinner( 100 * time.Millisecond )
	// Making It As Go Routine/Coroutine
	//		It Runs Concurrently
	go spinner( 100 * time.Millisecond ) 

	const n = 40
	fibN := fib( n ) // Slow
	fmt.Printf("\nFibonnacci (%d) = %d\n", n, fibN)
}

//_____________________________________________________________________


//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {
	fmt.Println("\nFunction : playWithFibonacci")
	playWithFibonacci()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}

