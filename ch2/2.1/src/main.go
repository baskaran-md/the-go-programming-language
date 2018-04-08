package main

import (
	"fmt"
	"tempconv"
)

func main() {
	var k tempconv.Kelvin
	fmt.Printf("%g*K = %s\n", k, tempconv.KToC(k))
	fmt.Printf("%g*K = %s\n", k, tempconv.KToF(k))

	var c tempconv.Celsius
	fmt.Printf("%g*C = %s\n", c, tempconv.CToK(c))
	fmt.Printf("%g*C = %s\n", c, tempconv.CToF(c))

	var f tempconv.Fahrenheit
	fmt.Printf("%g*F = %s\n", f, tempconv.FToC(f))
	fmt.Printf("%g*F = %s\n", f, tempconv.FToK(f))
}
