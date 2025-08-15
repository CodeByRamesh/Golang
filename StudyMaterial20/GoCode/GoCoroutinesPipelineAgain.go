
package main

import "fmt"

func generateNumbers( outChannel chan<- int ) {
	for number := 0 ; number <= 100 ; number++ {
		outChannel <- number
	}
	close( outChannel )
}

func squareNumbers(outChannel chan<- int, inChannel <-chan int) {
	for number := range inChannel {
		outChannel <- number * number
	}
	close( outChannel )	
}

func printer( inChannel <-chan int ) {
	for number := range inChannel {
		fmt.Println( number )
	}
}

func main() {
	numbersChannel 		:= make( chan int )
	squaresChannel 		:= make( chan int )

	go generateNumbers( numbersChannel )
	go squareNumbers( squaresChannel, numbersChannel )
	printer( squaresChannel )
}

