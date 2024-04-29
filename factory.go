package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ItemFactory struct responsible for creating items
type ItemFactory struct {
	LastID int
}

// Item represents a product
type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// NewItemFactory creates a new instance of ItemFactory
func NewItemFactory() *ItemFactory {
	return &ItemFactory{}
}

// CreateItem creates a new item
func (factory *ItemFactory) CreateItem(id int, name string, price float64, quantity int) *Item {
	return &Item{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

// addItemToJSONFile adds an item to the JSON file
func addItemToJSONFile(item *Item, filename string) error {
	// Read existing data from the file
	existingData, err := ioutil.ReadFile(filename)
	if err != nil {
		// If the file doesn't exist, create an empty array
		if os.IsNotExist(err) {
			existingData = []byte("[]")
		} else {
			return err
		}
	}

	// Unmarshal existing data
	var items []*Item
	if err := json.Unmarshal(existingData, &items); err != nil {
		return err
	}

	// Append new item
	items = append(items, item)

	// Marshal updated data
	newData, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		return err
	}

	// Write back to the file
	if err := ioutil.WriteFile(filename, newData, 0644); err != nil {
		return err
	}

	return nil
}

// readItemsFromJSONFile reads items from a JSON file
func readItemsFromJSONFile(filename string) ([]*Item, error) {
	// Read existing data from the file
	existingData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal existing data
	var items []*Item
	if err := json.Unmarshal(existingData, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// updateItem updates an item in the JSON file
func updateItem(filename string, itemID int, newName string, newPrice float64, newQuantity int) error {
	// Read existing items from JSON file
	items, err := readItemsFromJSONFile(filename)
	if err != nil {
		return err
	}

	// Find the item by ID
	found := false
	for _, item := range items {
		if item.ID == itemID {
			// Update item details
			item.Name = newName
			item.Price = newPrice
			item.Quantity = newQuantity
			found = true
			break
		}
	}

	// If item not found, return error
	if !found {
		return fmt.Errorf("item with ID '%d' not found", itemID)
	}

	// Marshal updated data
	newData, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		return err
	}

	// Write back to the file
	if err := ioutil.WriteFile(filename, newData, 0644); err != nil {
		return err
	}

	return nil
}

// deleteItem deletes an item from the JSON file
func deleteItem(filename string, itemID int) error {
	// Read existing items from JSON file
	items, err := readItemsFromJSONFile(filename)
	if err != nil {
		return err
	}

	// Find and remove the item by ID
	found := false
	for i, item := range items {
		if item.ID == itemID {
			items = append(items[:i], items[i+1:]...)
			found = true
			break
		}
	}

	// If item not found, return error
	if !found {
		return fmt.Errorf("item with ID '%d' not found", itemID)
	}

	// Marshal updated data
	newData, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		return err
	}

	// Write back to the file
	if err := ioutil.WriteFile(filename, newData, 0644); err != nil {
		return err
	}

	return nil
}

func main() {
	itemFactory := NewItemFactory()

	for {
		var choice string
		fmt.Println("Choose an action:")
		fmt.Println("1. Add item")
		fmt.Println("2. Read items")
		fmt.Println("3. Update item")
		fmt.Println("4. Delete item")
		fmt.Println("5. Exit")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			// Get item details from user
			var id int
			var name string
			var price float64
			var quantity int
			fmt.Println("Enter item details:")
			fmt.Print("ID: ")
			fmt.Scanln(&id)
			fmt.Print("Name: ")
			fmt.Scanln(&name)
			fmt.Print("Price: ")
			fmt.Scanln(&price)
			fmt.Print("Quantity: ")
			fmt.Scanln(&quantity)

			// Create new item
			newItem := itemFactory.CreateItem(id, name, price, quantity)

			// Add item to JSON file
			if err := addItemToJSONFile(newItem, "items.json"); err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Item added successfully.")
		case "2":
			// Read items from JSON file
			items, err := readItemsFromJSONFile("items.json")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			// Display items
			fmt.Println("Existing items:")
			for _, item := range items {
				fmt.Printf("ID: %d, Name: %s, Price: %.2f, Quantity: %d\n", item.ID, item.Name, item.Price, item.Quantity)
			}
		case "3":
			// Get item details from user
			var itemID int
			var newName string
			var newPrice float64
			var newQuantity int
			fmt.Println("Enter item ID to update:")
			fmt.Scanln(&itemID)
			fmt.Println("Enter new name:")
			fmt.Scanln(&newName)
			fmt.Println("Enter new price:")
			fmt.Scanln(&newPrice)
			fmt.Println("Enter new quantity:")
			fmt.Scanln(&newQuantity)

			// Update item details in JSON file
			if err := updateItem("items.json", itemID, newName, newPrice, newQuantity); err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Item updated successfully.")
		case "4":
			// Get item ID from user
			var itemID int
			fmt.Println("Enter item ID to delete:")
			fmt.Scanln(&itemID)

			// Delete item from JSON file
			if err := deleteItem("items.json", itemID); err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Item deleted successfully.")
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
