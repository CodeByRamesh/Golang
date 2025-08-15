
package main 

import "fmt"

func main() {
	// Buffered Channel Having Size 02
	queue := make( chan string, 2 )

	// Writing Data To The Channel
	queue <- "One"
	queue <- "Two"
	
	//fatal error: all goroutines are asleep - deadlock!
	//queue <- "Three"

	// Closing The Channel
	close( queue )
	
	// panic: send on closed channel
	// queue <- "Three"

	// Looping Over The Channel
	// Reading Channel Using For Loop
	for message := range queue {
		fmt.Println( message )
	}
}


