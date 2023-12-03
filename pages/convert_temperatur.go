// temperatur converter.go
package main

import "fmt"

func FarhenheitToCelcius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func celciusToFarhenheit(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	farhenheitTemp := 32.0
	celsiusTemp := 0.0

	// konveri farhenheit ke celcius
	celsiusTemp = FarhenheitToCelcius(farhenheitTemp)
	fmt.Println("Temperature in Celsius: ", farhenheitTemp, celsiusTemp, "\n")

	// konversi celcius ke farhenheit
	farhenheitTemp = celciusToFarhenheit(celsiusTemp)
	fmt.Printf("%.2f C is %2.f F\n", celsiusTemp, farhenheitTemp)
}
