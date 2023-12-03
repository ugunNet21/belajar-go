// src/contact_manager.go

package main

import "fmt"

// Struct untuk merepresentasikan kontak
type Contact struct {
	Name  string
	Phone string
}

// Slice untuk menyimpan daftar kontak
var contacts []Contact

// Menambahkan kontak ke dalam daftar
func addContact(name, phone string) {
	newContact := Contact{Name: name, Phone: phone}
	contacts = append(contacts, newContact)
	fmt.Println("Contact added successfully.")
}

// Menampilkan semua kontak dalam daftar
func listContacts() {
	fmt.Println("Contacts:")
	for _, contact := range contacts {
		fmt.Printf("Name: %s, Phone: %s\n", contact.Name, contact.Phone)
	}
}

func main() {
	// Menambahkan beberapa kontak
	addContact("John Doe", "123-456-789")
	addContact("Jane Doe", "987-654-321")

	// Menampilkan daftar kontak
	listContacts()
}
