package converter

// CToF convert celsius to fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit((9/5)*c + 32)
}

// FToC convert fahrenheit to celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK convert celsius to kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c + Celsius(KelvinConst))
}

// KToC convert kelvin to celsius
func KToC(k Kelvin) Celsius {
	return Celsius((k - KelvinConst))
}

// FToK convert fahrenheit to kelvin
func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

// KToF convert kelvin to fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

// MToF convert meters to feets
func MToF(m Meter) Feet {
	return Feet(m) / FeetConst
}

// FToM convert feets to meters
func FToM(f Feet) Meter {
	return Meter(f * FeetConst)
}

// KgToP convert kilograms to pounds
func KgToP(k Kilogram) Pound {
	return Pound(k) * PoundConst
}

// PToKg convert pounds to kilograms
func PToKg(p Pound) Kilogram {
	return Kilogram(p / PoundConst)
}
