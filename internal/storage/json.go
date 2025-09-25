package storage

import (
	"encoding/json"
	"fmt"
	"mini-crm/internal/models"
	"os"
)

type JsonStore struct {
	contacts map[int]*models.Contact
	nextID   uint
	filePath string
}

// Constructor
func NewJsonStore(path string) (*JsonStore, error) {
	ms := &JsonStore{
		contacts: make(map[int]*models.Contact),
		nextID:   1,
		filePath: path,
	}
	// Load existing data from file if it exists
	if err := ms.LoadFromFile(); err != nil {
		return nil, err
	}
	return ms, nil
}


// CRUD operations

func (ms *JsonStore) Add(contact *models.Contact) error {
	contact.ID = ms.nextID
	ms.contacts[int(contact.ID)] = contact
	ms.nextID++
	return ms.saveToFile()
}

func (ms *JsonStore) GetAll() ([]*models.Contact, error) {
	contacts := make([]*models.Contact, 0, len(ms.contacts))
	for _, c := range ms.contacts {
		contacts = append(contacts, c)
	}
	return contacts, nil
}

func (ms *JsonStore) GetByID(id int) (*models.Contact, error) {
	c, ok := ms.contacts[id]
	if !ok {
		return nil, fmt.Errorf("contact %d not found", id)
	}
	return c, nil
}

func (ms *JsonStore) Update(id int, newName, newEmail string) error {
	c, ok := ms.contacts[id]
	if !ok {
		return fmt.Errorf("contact %d not found", id)
	}
	if newName != "" {
		c.Name = newName
	}
	if newEmail != "" {
		c.Email = newEmail
	}
	return ms.saveToFile()
}

func (ms *JsonStore) Delete(id int) error {
	if _, ok := ms.contacts[id]; !ok {
		return fmt.Errorf("contact %d not found", id)
	}
	delete(ms.contacts, id)
	return ms.saveToFile()
}

// Load contacts from file into memory
func (ms *JsonStore) LoadFromFile() error {
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
		Contacts map[int]*models.Contact `json:"contacts"`
		NextID   uint              `json:"next_id"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}
	ms.contacts = raw.Contacts
	ms.nextID = raw.NextID
	return nil
}

// Save memory state into file
func (ms *JsonStore) saveToFile() error {
	raw := struct {
		Contacts map[int]*models.Contact `json:"contacts"`
		NextID   uint              `json:"next_id"`
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