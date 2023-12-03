// src/file_operations.go

package main

import (
	"fmt"
	"io/ioutil"
)

func writeFile(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	// Menulis ke file
	filename := "example.txt"
	content := "Hello, this is a sample content."
	err := writeFile(filename, content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Content written to", filename)

	// Membaca dari file
	readContent, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println("Content read from", filename, ":\n", readContent)
}
