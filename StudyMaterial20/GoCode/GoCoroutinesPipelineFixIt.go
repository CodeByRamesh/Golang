
package main

import "fmt"

func main() {
	numbersChannel 		:= make( chan int )
	evenNumbersChannel 	:= make( chan int )
	squaresChannel 		:= make( chan int )

	generateNumbers := func() {
		for number := 0 ;  ; number++ {
			numbersChannel <- number
		}
	}

	// filterNumbers := func(even bool, numbersChannel chan int ) {
	// 	number := 	  <- numbersChannel
	// 	// if even == true && number % 2 == 0 {
	// 	// 	squaresChannel <- number * number			
	// 	// } else {
	// 	// 	squaresChannel <- number * number
	// 	// }

	// 	if number % 2 == 0 {
	// 		evenNumbersChannel <- number
	// 	}
	// }

	squareNumbers := func( evenNumbersChannel chan int ) {
		for {
			number := 	   <- evenNumbersChannel
			squaresChannel <- number * number
		}
	}

	go generateNumbers( )
	// go filterNumbers( true, numbersChannel )	
	go squareNumbers( evenNumbersChannel )

	for {
		fmt.Println( <-squaresChannel )
	}
}

