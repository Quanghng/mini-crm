package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load contacts from file into memory
func (ms *MemoryStore) LoadFromFile() error {
	data, err := os.ReadFile(ms.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // file doesn't exist yet, that's fine
		}
		return fmt.Errorf("error reading file: %w", err)
	}
	if len(data) == 0 {
		return nil
	}

	var raw struct {
		Contacts map[int]*Contact `json:"contacts"`
		NextID   int              `json:"next_id"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}
	ms.contacts = raw.Contacts
	ms.nextID = raw.NextID
	return nil
}

// Save memory state into file
func (ms *MemoryStore) saveToFile() error {
	raw := struct {
		Contacts map[int]*Contact `json:"contacts"`
		NextID   int              `json:"next_id"`
	}{
		Contacts: ms.contacts,
		NextID:   ms.nextID,
	}

	data, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}
	return os.WriteFile(ms.filePath, data, 0644)
}