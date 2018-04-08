package tempconv

//CToF Converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//FToC Converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//CToK Converts Celsius to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(237.15 + c)
}

//KToC Converts Kelvin to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k - 237.15)
}

//KToF Converts Kelvin to Fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(CToF(KToC(k)))
}

//FToK Converts Fahrenheit to Kelvin
func FToK(f Fahrenheit) Kelvin {
	return Kelvin(CToK(FToC(f)))
}
