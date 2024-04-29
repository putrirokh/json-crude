package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func addItemToJSONFile(item Item, filename string) error {
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
	var items []Item
	if err := json.Unmarshal(existingData, &items); err != nil {
		return err
	}

	// Assign ID to the new item
	items = append(items, item)
	for i := range items {
		items[i].ID = i + 1
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

func readItemsFromJSONFile(filename string) ([]Item, error) {
	// Read existing data from the file
	existingData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal existing data
	var items []Item
	if err := json.Unmarshal(existingData, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func updateItem(filename string, itemID int, newName string, newPrice float64, newQuantity int) error {
	// Read existing items from JSON file
	items, err := readItemsFromJSONFile(filename)
	if err != nil {
		return err
	}

	// Find the item by ID
	found := false
	for i, item := range items {
		if item.ID == itemID {
			// Update item details
			items[i].Name = newName
			items[i].Price = newPrice
			items[i].Quantity = newQuantity
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
			var newItem Item
			fmt.Println("Enter item details:")
			fmt.Print("ID: ")
			fmt.Scanln(&newItem.ID)
			fmt.Print("Name: ")
			fmt.Scanln(&newItem.Name)
			fmt.Print("Price: ")
			fmt.Scanln(&newItem.Price)
			fmt.Print("Quantity: ")
			fmt.Scanln(&newItem.Quantity)

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
