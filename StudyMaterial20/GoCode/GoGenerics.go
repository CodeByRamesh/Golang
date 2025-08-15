package main

import "fmt"

//____________________________________________________________
//____________________________________________________________

func SumInts( numbers map[string]int64 ) int64 {
	var sum int64
	for _, number := range numbers {
		sum = sum + number
	}

	return sum
}

func SumFloats( numbers map[string]float64 ) float64 {
	var sum float64
	for _, number := range numbers {
		sum = sum + number
	}

	return sum
}

//____________________________________________________________

func playWithSumInts() {
	intNumbers := map[string]int64 {
		"first" : 100,
		"second": 200,
		"third" : 22,
	}

	result := SumInts( intNumbers )
	fmt.Println("\nResult : ", result)

	floatNumbers := map[string]float64 {
		"first" : 100.1,
		"second": 200.2,
		"third" : 22.3,
	}

	resultAgain := SumFloats( floatNumbers )
	fmt.Println("\nResult : ", resultAgain)

}

//____________________________________________________________


func SumIntsOrFloats [K comparable, V int64 | float64]( numbers map[K]V ) V {
	var sum V
	for _, number := range numbers {
		sum = sum + number
	}

	return sum
}

func playWithSumIntsOrFloats() {
	intNumbers := map[string]int64 {
		"first" : 100,
		"second": 200,
		"third" : 22,
	}

	result := SumIntsOrFloats( intNumbers )
	fmt.Println("\nResult : ", result)

	floatNumbers := map[string]float64 {
		"first" : 100.1,
		"second": 200.2,
		"third" : 22.3,
	}

	resultAgain := SumIntsOrFloats( floatNumbers )
	fmt.Println("\nResult : ", resultAgain)

}

//____________________________________________________________

func SumNumbers [K comparable, V Nummber ]( numbers map[K]V ) V {
	var sum V
	for _, number := range numbers {
		sum = sum + number
	}

	return sum
}

func playWithSumNumbers() {
	intNumbers := map[string]int64 {
		"first" : 100,
		"second": 200,
		"third" : 22,
	}

	result := SumNumbers( intNumbers )
	fmt.Println("\nResult : ", result)

	floatNumbers := map[string]float64 {
		"first" : 100.1,
		"second": 200.2,
		"third" : 22.3,
	}

	resultAgain := SumNumbers( floatNumbers )
	fmt.Println("\nResult : ", resultAgain)

}


type Nummber interface {
	int64 | float64
}


//____________________________________________________________
//____________________________________________________________
//____________________________________________________________
//____________________________________________________________

func main() {
	fmt.Println("\nFunction : playWithSumInts")
	playWithSumInts()

	fmt.Println("\nFunction : playWithSumIntsOrFloats")
	playWithSumIntsOrFloats()

	fmt.Println("\nFunction : playWithSumNumbers")
	playWithSumNumbers()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}

