package storage

import (
	"fmt"
)

type MemoryStore struct {
	contacts map[int]*Contact
	nextID   int
	filePath string
}

// Constructor
func NewMemoryStore(path string) (*MemoryStore, error) {
	ms := &MemoryStore{
		contacts: make(map[int]*Contact),
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

func (ms *MemoryStore) Add(contact *Contact) error {
	contact.ID = ms.nextID
	ms.contacts[contact.ID] = contact
	ms.nextID++
	return ms.saveToFile()
}

func (ms *MemoryStore) GetAll() ([]*Contact, error) {
	contacts := make([]*Contact, 0, len(ms.contacts))
	for _, c := range ms.contacts {
		contacts = append(contacts, c)
	}
	return contacts, nil
}

func (ms *MemoryStore) GetByID(id int) (*Contact, error) {
	c, ok := ms.contacts[id]
	if !ok {
		return nil, fmt.Errorf("contact %d not found", id)
	}
	return c, nil
}

func (ms *MemoryStore) Update(id int, newName, newEmail string) error {
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

func (ms *MemoryStore) Delete(id int) error {
	if _, ok := ms.contacts[id]; !ok {
		return fmt.Errorf("contact %d not found", id)
	}
	delete(ms.contacts, id)
	return ms.saveToFile()
}
