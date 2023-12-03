// src/concurrent_operations.go

package main

import (
	"fmt"
	"sync"
)

func calculateSum(numbers []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	ch <- sum
}

func main() {
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int{6, 7, 8, 9, 10}

	// Membuat channel untuk komunikasi antar goroutines
	ch := make(chan int)

	// Membuat WaitGroup untuk menunggu selesainya semua goroutines
	var wg sync.WaitGroup

	// Menambahkan 2 ke WaitGroup karena kita akan menjalankan 2 goroutines
	wg.Add(2)

	// Menjalankan goroutine untuk menghitung sum dari numbers1
	go calculateSum(numbers1, ch, &wg)

	// Menjalankan goroutine untuk menghitung sum dari numbers2
	go calculateSum(numbers2, ch, &wg)

	// Menjalankan goroutine untuk menunggu selesainya semua goroutines
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Menerima hasil dari channel dan menjumlahkannya
	total := 0
	for sum := range ch {
		total += sum
	}

	fmt.Printf("Total sum: %d\n", total)
}
