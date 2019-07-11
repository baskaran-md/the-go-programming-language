/*
1.1: Echo Program to print the command that invoked it
*/
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	// os.Args[0] prints the command name
	fmt.Printf("Program: %v\n", os.Args[0])
	fmt.Fprintf(out, "All Args: %v\n", strings.Join(os.Args[1:], ","))
}
