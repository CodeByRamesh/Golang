package main

import (
	"fmt"
)

//_____________________________________________________________________

func playWithArrays() {

	// Creating Array of 3 ints
	// Index Starts From 0 Till len - 1 i.e. [0, len)
	var a [3]int

	fmt.Println("\nArray a Length: ", len(a))
	for index, value := range a {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])

	// Compilation Error
	// invalid argument: array index 3 out of bounds [0:3]
	// fmt.Println( a[ len( a ) ] )

	fmt.Println("\nArray a Length: ", len(a))
	for _, value := range a {
		fmt.Printf("Value: %d\n", value)
	}

	var q [5]int = [5]int{10, 20, 30, 40, 50}
	var r [5]int = [5]int{10, 20, 30}

	fmt.Println("\nArray q Length: ", len(q))
	for index, value := range q {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	fmt.Println("\nArray r Length: ", len(r))
	for index, value := range r {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	s := [...]int{10, 20, 100, 111}
	fmt.Println("\nArray s Length: ", len(s))
	fmt.Printf("\nArray s Data Type: %T\n", s)
	for index, value := range s {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	some := [3]int{100, 200, 300}
	fmt.Println("\nArray s Length: ", len(some))
	fmt.Printf("\nArray s Data Type: %T\n", some)

	// NOTE: Array Size Is Part of Array Type
	// Following Are Three Different Arrays Allocated
	//		In Three Different Memory Locations
	aa := [2]int{10, 20} // Type [2]int Array Of int Of Size 2
	// ... Means Array Size Computed From Number Of Member Elements
	bb := [...]int{10, 20} // Type [2]int
	cc := [2]int{10, 30}   // Type [2]int

	fmt.Printf("Array s Data Type: %T\n", aa)
	fmt.Printf("Array s Data Type: %T\n", bb)
	fmt.Printf("Array s Data Type: %T\n", cc)

	// In Go
	// 		Using == Operator
	//		Array Values Are Compared Of Arrays of Same Type

	// In Java
	// 		Using == Operator
	// 		Array References Are Compared Of Arrays of Same Type
	fmt.Println(aa == bb, aa == cc, bb == cc)
	// true false false

	dd := [3]int{10, 30} // Type [3]int
	fmt.Printf("Array s Data Type: %T\n", dd)
	fmt.Println(dd)

	//Compiler Error
	// 	invalid operation: aa == dd (mismatched types [2]int and [3]int)
	// Two Values Are Equal Only When Types Are Same
	// fmt.Println( aa == dd )

	var someAgain1 [10]int
	fmt.Println("\nArray s Length: ", len(someAgain1))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain1)
	for index, value := range someAgain1 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\nArray s Length: ", len(someAgain2))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain2)
	for index, value := range someAgain2 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\nArray s Length: ", len(someAgain3))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain3)
	for index, value := range someAgain3 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain4 := [...]int{5: 50, 7: 70}
	fmt.Println("\nArray s Length: ", len(someAgain4))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain4)
	for index, value := range someAgain4 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain5 := [...]int{5: 50, 7: 70, 12: 222}
	fmt.Println("\nArray s Length: ", len(someAgain5))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain5)
	for index, value := range someAgain5 {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}

	someAgain := [...]int{99: -1}
	fmt.Println("\nArray s Length: ", len(someAgain))
	fmt.Printf("\nArray s Data Type: %T\n", someAgain)
	for index, value := range someAgain {
		fmt.Printf("At Index: %d Value: %d\n", index, value)
	}
}

//_____________________________________________________________________

func zeroValue(ival int) {
	ival = 0
}

// Pointer Variable
//		Is A Variable Used To Address Of Memory Location

func zeroPointer(iptr *int) {
	*iptr = 0
}

func playWithPointers() {
	var something = 111

	fmt.Println("Something Value: ", something)
	zeroValue(something)
	fmt.Println("Something Value: ", something)

	zeroPointer(&something)
	fmt.Println("Something Value: ", something)
}

// Function : playWithPointers
// Something Value:  111
// Something Value:  111
// Something Value:  0

//_____________________________________________________________________

type Flags uint

const ( // Definining Constants For Status BITs
	FlagUp           Flags = 1 << iota // Is Up  //  1 << 0
	FlagBroadcast                      //  1 << 1
	FlagLoopback                       //  1 << 2
	FlagPointToPoint                   //  1 << 3
	FlagMulticast                      //  1 << 4
)

func IsUp(v Flags) bool { return v&FlagUp == FlagUp } // & Is Bitwise AND
func TurnDown(v *Flags) { *v &^= FlagUp }             // &^ AND NOT Operator
// This Is A Bit Clear Operator
func SetBroadcast(v *Flags) { *v |= FlagBroadcast } // |  Bitwise OR Operator
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func playWithFlags() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))

	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))

	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}

// Function : playWithFlags
// 10001 true
// 10000 false
// 10010 false
// 10010 true

func main() {
	fmt.Println("\nFunction : playWithArrays")
	playWithArrays()

	fmt.Println("\nFunction : playWithPointers")
	playWithPointers()

	fmt.Println("\nFunction : playWithFlags")
	playWithFlags()
}
