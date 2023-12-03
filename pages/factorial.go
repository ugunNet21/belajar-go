// factorial
package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	num := 5
	result := factorial(num)
	fmt.Printf("The factorial of %d is: %d", num, result)
}
