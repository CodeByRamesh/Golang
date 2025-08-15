package main

import (
	"fmt"
)

func playWithArrays() {

	// Array Of string Type Data
	planets := [...]string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune",
	}

	planets1 := [8]string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune",
	}

	fmt.Println(planets)
	fmt.Println(planets1)

	livablePlanets := planets[2:4]
	icePlanets := planets[6:8]

	fmt.Println(livablePlanets, icePlanets)

	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]
	fmt.Println(terrestrial, gasGiants, iceGiants)
}

func playWithMaps() {
	// Map Is Collection Of Tuples
	//		Each Tuple Is (Key, Value) Pair

	// Key Type Is string and Value Type int
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	fmt.Println(temperature)
	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %vº C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)

}

func main() {

	fmt.Println("\nFunction :playWithArrays")
	playWithArrays()

	fmt.Println("\nFunction : playWithMaps")
	playWithMaps()

}
