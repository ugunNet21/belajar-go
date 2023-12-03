// src/slice_search_sort.go

package main

import (
	"fmt"
	"sort"
)

func searchInSlice(slice []int, target int) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

func sortSlice(slice []int) {
	sort.Ints(slice)
}

func main() {
	// Contoh slice
	numbers := []int{4, 2, 7, 1, 9, 5, 8}

	// Mencari angka 7 dalam slice
	targetNumber := 7
	found := searchInSlice(numbers, targetNumber)

	if found {
		fmt.Printf("%d found in the slice.\n", targetNumber)
	} else {
		fmt.Printf("%d not found in the slice.\n", targetNumber)
	}

	// Mengurutkan slice
	fmt.Println("Original Slice:", numbers)
	sortSlice(numbers)
	fmt.Println("Sorted Slice:", numbers)
}
