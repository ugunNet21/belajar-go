// src/sqlite_database.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Item adalah struktur data untuk merepresentasikan item
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./items.db")
	if err != nil {
		log.Fatal(err)
	}

	// Membuat tabel jika belum ada
	createTable := `
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			price INTEGER NOT NULL
		);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllItems() ([]Item, error) {
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func addItem(name string, price int) (int64, error) {
	result, err := db.Exec("INSERT INTO items (name, price) VALUES (?, ?)", name, price)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func main() {
	initDB()

	// Menambahkan beberapa item awal
	_, err := addItem("Item A", 15)
	if err != nil {
		log.Fatal(err)
	}
	_, err = addItem("Item B", 25)
	if err != nil {
		log.Fatal(err)
	}

	// Membaca semua item dari database
	items, err := getAllItems()
	if err != nil {
		log.Fatal(err)
	}

	// Menampilkan hasil
	fmt.Println("Items in Database:")
	for _, item := range items {
		fmt.Printf("ID: %d, Name: %s, Price: %d\n", item.ID, item.Name, item.Price)
	}

	// Menutup koneksi database pada akhir program
	defer db.Close()

	// Menghapus database file setelah program dijalankan
	err = os.Remove("./items.db")
	if err != nil {
		log.Fatal(err)
	}
}
