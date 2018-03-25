/*
1.1 Echo Program to print the command that invoked it
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// os.Args[0] prints the command name
	fmt.Printf("Program: %v\n", os.Args[0])
	fmt.Printf("All Args: %v\n", strings.Join(os.Args[1:], ","))
}
