package main

import (
	"fmt"
	"mini-crm/internal/app"
	"mini-crm/internal/storage"
)


func main() {
	store, err := storage.NewMemoryStore("contacts.json")
	if err != nil {
		fmt.Printf("Error initializing store: %v\n", err)
		return
	}

	app.Run(store)
}
