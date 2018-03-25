/*
1.3: Echo Program that measures runtime
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Printf("Program: %v\n", os.Args[0])

	// Method#1
	// Prints each argument along with its index (one per line)
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("Arg#%d: %v\n", i, os.Args[i])
	}
	sec := time.Since(start).Seconds()
	fmt.Printf("Time taken for method#1: %fs\n", sec)

	// Method#2
	// Joins all the argument using strings.Join function and prints them
	start = time.Now()
	fmt.Printf("All Args: %v\n", strings.Join(os.Args[1:], ","))
	sec = time.Since(start).Seconds()
	fmt.Printf("Time taken for method#2: %fs\n", sec)

	// Method#3
	// Loops directly on the Args, and print them with index
	start = time.Now()
	for i, val := range os.Args[1:] {
		fmt.Printf("Arg#%d: %v\n", i, val)
	}
	sec = time.Since(start).Seconds()
	fmt.Printf("Time taken for method#3: %fs\n", sec)
}
