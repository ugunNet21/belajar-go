// membuat star piramid
package main

import "fmt"

func PrintStarPiramid(rows int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= (2*rows - 1); j++ {
			for j := 0; j < rows-1-1; j++ {
				fmt.Print(" ")
			}
			for k := 0; k < 2*i+1; k++ {
				fmt.Printf("*")
			}
			fmt.Println()
		}
	}
}
func main() {
	rows := 5
	PrintStarPiramid(rows)
}
