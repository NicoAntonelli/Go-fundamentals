package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ContactInfo struct {
	name string
	email string
	zipCode int
}

type PurchaseInfo struct {
	name string
	item string
	price float64
	amount int
}

// Work with files
func files() {
	fmt.Println("Files")
	var contacts []ContactInfo = loadJSON[ContactInfo]("./contactInfo.json")
	var purchases []ContactInfo = loadJSON[ContactInfo]("./contactInfo.json")
	fmt.Printf("Contacts: %+v\n", contacts)
	fmt.Printf("Purchases: %+v\n", purchases)
	fmt.Println()
}

// Read a JSON file
func loadJSON[T ContactInfo | PurchaseInfo](path string) []T {
	// Read file and validate if exists
	var data []T
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Parse file and validate if succeed
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return data
}
