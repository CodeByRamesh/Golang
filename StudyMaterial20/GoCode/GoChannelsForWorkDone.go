
package main 

import (
	"fmt"
	"time"
)

func chefCookRecipe( cookingDone chan bool ) {
	fmt.Println("Prepare Ingrediants")
	fmt.Println("Prepare Receipe...")
	time.Sleep( time.Second * 20 ) // Cooking Taking Time...

	// Using Channel To Communicate Completion Of Work
	cookingDone <- true
}

// Using Done Pattern To Communicate Work Completion
func main() { // Main Go Routine 01
	cookingStatus := make( chan bool, 1 )

	// chefCookRecipe Go Routine 02
	go chefCookRecipe( cookingStatus )

	//	Reading Work Status From Channel
	cooked := <- cookingStatus 
	if cooked {
		fmt.Println("Food Is Ready...")
	}
}

