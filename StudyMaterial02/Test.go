package main

import (
	"fmt"
	"math"
)

func sum(x, y int32) {
	fmt.Println(math.MaxInt32, math.MinInt32)
	if (y > 0 && x > math.MaxInt32-y) || (y < 0 && x < math.MinInt32-y) {
		fmt.Printf("%d\t%d\n", x, y)
		fmt.Println("Can't calculate given sum: integer overflow")
		return
	}
	fmt.Println(x + y)
}

func main() {
	sum(2147483640, 10) // Will cause overflow for int32
	sum(100, 200)       // Safe
}
