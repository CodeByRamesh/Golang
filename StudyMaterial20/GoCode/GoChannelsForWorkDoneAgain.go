
package main 

import (
	"fmt"
	"time"
)

func chefOneCookRecipe( cookingDone chan bool ) {
	fmt.Println("Chef 01:Prepare Ingrediants")
	fmt.Println("Chef 01:Prepare Receipe...")
	time.Sleep( time.Second * 20 ) // Cooking Taking Time...

	// Using Channel To Communicate Completion Of Work
	cookingDone <- true
}

func chefTwoCookRecipe( cookingDone chan bool ) {
	fmt.Println("Chef 02:Prepare Ingrediants")
	fmt.Println("Chef 02:Prepare Receipe...")
	time.Sleep( time.Second * 5 ) // Cooking Taking Time...

	// Using Channel To Communicate Completion Of Work
	cookingDone <- true
}

func chefThreeCookRecipe( cookingDone chan bool ) {
	fmt.Println("Chef 03:Prepare Ingrediants")
	fmt.Println("Chef 03:Prepare Receipe...")
	time.Sleep( time.Second * 2 ) // Cooking Taking Time...

	// Using Channel To Communicate Completion Of Work
	cookingDone <- true
}

// Using Done Pattern To Communicate Work Completion
func main() { // Main Go Routine 01
	cookingStatusOne 	:= make( chan bool, 1 )
	cookingStatusTwo 	:= make( chan bool, 1 )
	cookingStatusThree 	:= make( chan bool, 1 )

	// Go Routine 02
	go chefOneCookRecipe( cookingStatusOne )
	// Go Routine 03
	go chefTwoCookRecipe( cookingStatusTwo )
	// Go Routine 04
	go chefThreeCookRecipe( cookingStatusThree )

	// //	Reading Work Status From Channel
	// var cooked bool

	// // Reading From Empty Channel Will Block Go Routine
	// // Block Till Channel Is Empty
	// cooked = <- cookingStatusOne 
	// if cooked {
	// 	fmt.Println("Chef 01: Food Is Ready...")
	// }

	// cooked = <- cookingStatusTwo
	// if cooked {
	// 	fmt.Println("Chef 02: Food Is Ready...")
	// }

	// cooked = <- cookingStatusThree 
	// if cooked {
	// 	fmt.Println("Chef 03: Food Is Ready...")
	// }

	// select Statement Waits On Multiple Channels Simultaneosly
	//		Combining Go Routines and Channels With select

	for i := 0 ; i < 3 ; i++ {
		select {
		case status := <- cookingStatusOne:
			fmt.Println("Chef 01: Food Is Ready...", status)			
		case status := <- cookingStatusTwo:
			fmt.Println("Chef 02: Food Is Ready...", status)			
		case status := <- cookingStatusThree:
			fmt.Println("Chef 03: Food Is Ready...", status)			
		}
	}
}

