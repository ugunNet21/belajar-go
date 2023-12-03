// membuat aplikasi calculator sederhana
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func substract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}
func divide(x, y int) float64 {
	if y == 0 {
		return 0.0
	}
	return float64(x) / float64(y)
}

func main() {
	num1 := 10
	num2 := 5

	fmt.Printf("Addition: %d\n", add(num1, num2))
	fmt.Printf("Substraction: %d\n", substract(num1, num2))
	fmt.Printf("Multiplication: %d\n", multiply(num1, num2))
	fmt.Printf("Division: %.2f\n", divide(num1, num2))
}
