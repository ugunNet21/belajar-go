// src/simple_api.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Item adalah struktur data untuk merepresentasikan item
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items []Item

func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Mendapatkan ID dari URL
	id := r.URL.Query().Get("id")

	// Mengonversi ID menjadi integer
	itemID := 0
	fmt.Sscanf(id, "%d", &itemID)

	// Mencari item dengan ID yang sesuai
	var foundItem Item
	for _, item := range items {
		if item.ID == itemID {
			foundItem = item
			break
		}
	}

	// Mengirimkan item atau pesan error jika tidak ditemukan
	if foundItem.ID != 0 {
		json.NewEncoder(w).Encode(foundItem)
	} else {
		http.Error(w, "Item not found", http.StatusNotFound)
	}
}

func addItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Menambahkan item baru ke dalam daftar
	items = append(items, newItem)
	json.NewEncoder(w).Encode(newItem)
}

func main() {
	// Menambahkan beberapa item awal
	items = []Item{
		{ID: 1, Name: "Item 1", Price: 10},
		{ID: 2, Name: "Item 2", Price: 20},
		{ID: 3, Name: "Item 3", Price: 30},
	}

	// Menetapkan route ke handler
	http.HandleFunc("/api/items", getAllItems)
	http.HandleFunc("/api/item", getItemByID)
	http.HandleFunc("/api/add", addItem)

	// Menjalankan server pada port 8080
	fmt.Println("API Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
