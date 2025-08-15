
package main

import "fmt"

func main() {
	numbersChannel 		:= make( chan int )
	squaresChannel 		:= make( chan int )

	generateNumbers := func() {
		for number := 0 ;  ; number++ {
			numbersChannel <- number
		}
	}

	squareNumbers := func() {
		for {
			number := 	   <- numbersChannel
			squaresChannel <- number * number
		}
	}

	go generateNumbers( )
	go squareNumbers( )

	for {
		fmt.Println( <-squaresChannel )
	}
}

