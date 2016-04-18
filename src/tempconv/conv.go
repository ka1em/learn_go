package tempconv

// CToF ...
func CToF(C Celsius) Fahrenheit {
	return Fahrenheit(C*9/5 + 32)
}

// FToC ...
func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32) * 5 / 9)
}