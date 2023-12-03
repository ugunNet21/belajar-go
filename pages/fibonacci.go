// fibonacci.go

package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func sumFibonacci(limit int) int {
	sum := 0
	for i := 0; fibonacci(i) <= limit; i++ {
		sum += fibonacci(i)
	}
	return sum
}

func main() {
	limit := 50
	result := sumFibonacci(limit)
	fmt.Printf("Sum of Fibonacci numbers less than or equel to %d is : %d\n", limit, result)
}
