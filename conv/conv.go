package conv

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(fahr float64) float64 {
	var cel = (fahr - 32) * 5 / 9
	return cel
}

// Konverterer Celsius til Farhenheit
func CelsiusToFahrenheit(cel float64) float64 {
	var farh = cel*9/5 + 32
	return farh
}

// Konverterer Kelvin til Farhenheit
func KelvinToFarhenheit(kel float64) float64 {
	var fahr = (kel-273.15)*9/5 + 32
	return fahr
}

// Konverterer Farhenheit til Kelvin
func FahrenheitToKelvin(fahr float64) float64 {
	var kel = (fahr-32)*5/9 + 273.15
	return kel
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(cel float64) float64 {
	var kel = cel + 273.15
	return kel
}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(kel float64) float64 {
	var cel = kel - 273.15
	return cel
}
