
package main 

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var n = flag.Bool("n", false, "Omit Trailing Newline...")
	var sep = flag.String("s", " ", "Used As Seperator")

	flag.Parse()

	fmt.Print( strings.Join( flag.Args(), *sep ) )
	if !*n {
		fmt.Println()
	}
}

// GoCommandLineArguments:
//   -n	Omit Trailing Newline...
//   -s string
//     	seperator (default " ")
