

package main 

import "fmt"

// func ping( pingsChannel chan string, message string ) {

// pingsChannel Is Write/Send Only Channel
func ping( pingsChannel chan<- string, message string ) {
	// Writing To Channel
	pingsChannel <- message

	// invalid operation: cannot receive from send-only channel 
	// pingsChannel (variable of type chan<- string)
	// msg := <- pingsChannel
}

// func pong( pingsChannel chan string, pongsChannel chan string ) {

// pingsChannel Is Read/Receive Only Channel
// pongsChannel Is Read/Receive Only Channel
func pong( pingsChannel <-chan string, pongsChannel chan<- string ) {
	// Reading From Channel
	message := <- pingsChannel
	pongsChannel <- "Hello, " + message
}

func main() {
	// Allocating Two Buffered Channels of Size 01
	pingsChannel := make( chan string, 1 )
	pongsChannel := make( chan string, 1 )

	ping( pingsChannel, "Good Morning!")
	pong( pingsChannel, pongsChannel )
	fmt.Println( <- pongsChannel )

	ping( pingsChannel, "Good Afternoon!")
	pong( pingsChannel, pongsChannel )
	fmt.Println( <- pongsChannel )

	ping( pingsChannel, "Good Evening!")
	pong( pingsChannel, pongsChannel )
	fmt.Println( <- pongsChannel )
}

