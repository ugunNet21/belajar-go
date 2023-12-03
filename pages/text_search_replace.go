// src/text_search_replace.go

package main

import (
	"fmt"
	"strings"
)

func searchAndReplace(text, search, replace string) string {
	result := strings.Replace(text, search, replace, -1)
	return result
}

func main() {
	originalText := "Ini adalah contoh teks. Teks ini digunakan untuk mencoba pencarian dan penggantian kata."

	searchWord := "teks"
	replaceWord := "string"

	updatedText := searchAndReplace(originalText, searchWord, replaceWord)

	fmt.Println("Original Text:")
	fmt.Println(originalText)
	fmt.Println("\nUpdated Text:")
	fmt.Println(updatedText)
}
