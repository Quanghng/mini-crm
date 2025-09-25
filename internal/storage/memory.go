package storage

import (
	"mini-crm/internal/models"
)

// MemoryStore est un stockage éphémère en mémoire
type MemoryStore struct {
	contacts map[int]*models.Contact
	nextID   uint
}

// NewMemoryStore crée un nouveau MemoryStore vide
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*models.Contact),
		nextID:   1,
	}
}

// Add ajoute un contact
func (ms *MemoryStore) Add(contact *models.Contact) error {
	contact.ID = ms.nextID
	ms.contacts[int(contact.ID)] = contact
	ms.nextID++
	return nil
}

// GetAll retourne tous les contacts
func (ms *MemoryStore) GetAll() ([]*models.Contact, error) {
	list := make([]*models.Contact, 0, len(ms.contacts))
	for _, c := range ms.contacts {
		list = append(list, c)
	}
	return list, nil
}

func (ms *MemoryStore) GetByID(id int) (*models.Contact, error) {
	c, ok := ms.contacts[id]
	if !ok {
		return nil, ErrContactNotFound(id)
	}
	return c, nil
}

func (ms *MemoryStore) Update(id int, newName, newEmail string) error {
	c, ok := ms.contacts[id]
	if !ok {
		return ErrContactNotFound(id)
	}
	if newName != "" {
		c.Name = newName
	}
	if newEmail != "" {
		c.Email = newEmail
	}
	return nil
}

func (ms *MemoryStore) Delete(id int) error {
	if _, ok := ms.contacts[id]; !ok {
		return ErrContactNotFound(id)
	}
	delete(ms.contacts, id)
	return nil
}
