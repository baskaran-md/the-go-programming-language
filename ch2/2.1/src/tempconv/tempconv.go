package tempconv

import "fmt"

// Celsius represents the float64 value of temperature in Celsius.
type Celsius float64

// Fahrenheit represents the float64 value of temperature in Fahrenheit.
type Fahrenheit float64

// Kelvin represents the float64 value of temperature in Kelvin.
type Kelvin float64

// Constants representing different temperatures in Celsius.
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g*C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g*F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g*K", k)
}
