package main

import (
	"fmt"
	"math"
)

//_____________________________________________________________________

// In Go float32 And float64 Type
//		float32 Is Same As float in C/C++/Java
//		float64 Is Same As double in C/C++/Java

func playWithTypes() {
	// Inferred Type From RHS Is float64 i.e. double In C
	// 		float64 Binded With LHS
	//		Hence pi Type Is float64
	pi := 3.1415

	// Inferred Type From RHS Is float64 i.e. double In C
	// 		float64 Binded With LHS
	//		Hence piValue Type Is float64
	var piValue = 3.1415
	var piAgain float64 = 3.1415

	// All Above Declarations Are Equivalent
	fmt.Println(pi, piValue, piAgain)

	// To Make float32 Explicitly Annotate It!
	//		float32 Is Same As float in C/C++/Java
	var piAgain1 float32 = 3.1415
	fmt.Println(piAgain1)

	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)
	test := 123456.9999999
	fmt.Printf("%3.2f\n", test)
}

//_____________________________________________________________________

func playWithFloatingPointTypes() {
	// Type Is float64
	piggyBank := 0.1
	piggyBank = piggyBank + 0.2

	fmt.Println(piggyBank == 0.3)

	// Function : playWithFloatingPointTypes
	// false

	// BAD CODE : Buggy Logic
	//		Serious Security Threat :  Validation Can Be Bypassed
	if piggyBank == 0.3 {
		fmt.Println("piggyBank Have Dollers : ", piggyBank)
	} else {
		fmt.Println("piggyBank Is Dead...")
	}

	// GOOD CODE
	fmt.Println("math.Abs(piggyBank)::", math.Abs(piggyBank))
	fmt.Println("math.Abs(piggyBank-0.3)::", math.Abs(piggyBank-0.3))
	if math.Abs(piggyBank-0.3) < 0.000001 {
		fmt.Println("piggyBank Have Dollers : ", piggyBank)
	} else {
		fmt.Println("piggyBank Is Dead...")
	}
}

// DESIGN PRINCIPLE AND PRACTICE
//		Never Compare Floating Point Values With Equality/Not Equality Operator ==, !=
//		Correct Mathematics For Comparing Floating Points As Follows
// 			| floatingValue1 - floatingValue2 | < ToleranceValue

type X struct {
	data []int
}

func modify(x X) {
	x.data[0] = 100
}
func main() {
	fmt.Println("\nFunction : playWithTypes")
	playWithTypes()

	fmt.Println("\nFunction : playWithFloatingPointTypes")
	playWithFloatingPointTypes()

	obj := X{data: []int{1, 2, 3}}
	modify(obj)
	fmt.Println(obj.data[0])
	fmt.Println(obj)

	evenNumbers := []int{2, 4}
	oddNumbers := []int{1, 3}

	// add elements of oddNumbers to evenNumbers
	evenNumbers = append(evenNumbers, oddNumbers)
	fmt.Println(evenNumbers)

}
