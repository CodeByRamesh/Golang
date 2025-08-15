package main

import "fmt"

// Communication Between Go Routines Using Channel
func main() { // Main Go Routine 01
	// Allocating Unbuffered Channel

	// Go Channels
	// If goroutines are the activities of a concurrent Go program, 
	// channels are the connections between them. A channel is a 
	// communication mechanism that lets one goroutine send values
	// to another goroutine. Each channel is a conduit for values of 
	// a particular type, called the channel’s element type. 
	// The type of a channel whose elements have type int is written
	//		chan int.

	// To create a channel, we use the built-in make function:
	//		ch := make(chan int) // ch has type 'chan int'

	// As with maps, a channel is a reference to the data structure 
	// created by make. When we copy a channel or pass one as an argument 
	// to a function, we are copying a reference, so caller and callee 
	// refer to the same data structure. 
	// As with other reference types, the zero value of a channel is nil.


	//Two channels of the same type may be compared using ==. 
	// The comparison is true if both are references to the same channel 
	// data structure. A channel may also be compared to nil.

	//	Channels support a third operation, close, which sets a flag indicating 
	/// that no more values will ever be sent on this channel; 
	// subsequent attempts to send will panic. 
	// Receive operations on a  closed channel yield the values that have 
	// been sent until no more values are left; any receive operations 
	// thereafter complete immediately and yield the zero value of the 
	// channel’s element type.

	// To close a channel, we call the built-in close function:
	// 			close(ch)
	
	// A channel created with a simple call to make is called an 
	// unbuffered channel, but make accepts an optional second 
	// argument, an integer called the channel’s capacity. 
	// If the capacity is non zero, make creates a buffered channel.
			// ch = make(chan int)   	// unbuffered channel
			// ch = make(chan int, 0)  // unbuffered channel
			// ch = make(chan int, 3)  // buffered channel with capacity 3

	// A send operation on an unbuffered channel blocks the sending goroutine 
	// until another goroutine executes a corresponding receive on the 
	// same channel, at which point the value is transmitted and both 
	// goroutines may continue. Conversely, if the receive operation was
	//attempted first, the receiving goroutine is blocked until another 
	// goroutine performs a send on the same channel.

	// Communication over an unbuffered channel causes the sending and 
	// receiving goroutines to synchronize. Because of this, unbuffered 
	// channels are sometimes called synchronous channels. When a value is sent 
	// on an unbuffered channel, the receipt of the value happens before the
	// reawakening of the sending goroutine.

	messageChannel := make( chan string )

	sendMessageClosure := func() {
		// Writing/Sending "Ping" Message To messageChannel
		messageChannel <- "Ping"
		messageChannel <- "Pong"
		messageChannel <- "Ting"
		messageChannel <- "Tong"
	}

	// Clsoure Go Routine 02 : Invoking Clousre As Go Routine 
	go sendMessageClosure()

	var message string
	//		   Reading/Receiving From messageChannel
	// Storing Message Read In Variable message
	message = 	<- messageChannel
	fmt.Println( message )
	
	message = 	<- messageChannel
	fmt.Println( message )

	message = 	<- messageChannel
	fmt.Println( message )

	message = 	<- messageChannel
	fmt.Println( message )
}
