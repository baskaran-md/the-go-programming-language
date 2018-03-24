/*
1.2 Echo Program to print the index and value of each args
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Program: %v\n", os.Args[0])

    // Method#1
    // Prints each argument along with its index (one per line)
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("Arg#%d: %s\n", i, os.Args[i])
	}

    // Method#2
    // Loops directly on the Args, and print them with index
    for i, val := range os.Args[1:] {
        fmt.Printf("Arg#%d: %v\n", i, val)
    }
}
