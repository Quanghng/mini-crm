package repository

import (
	"fmt"
	"mini-crm/internal/database"
	"mini-crm/internal/models"
	"mini-crm/internal/storage"

	"gorm.io/gorm"
)

type GORMStore struct{}

// Constructor
func NewGORMStore() *GORMStore {
	return &GORMStore{}
}

// Add a new contact
func (gs *GORMStore) Add(contact *models.Contact) error {
	// Create a new record in DB
	model := models.Contact{
		Name:  contact.Name,
		Email: contact.Email,
	}
	if err := database.DB.Create(&model).Error; err != nil {
		return fmt.Errorf("erreur lors de l'insertion: %w", err)
	}

	contact.ID = model.ID 
	return nil
}

// Get all contacts
func (gs *GORMStore) GetAll() ([]*models.Contact, error) {
	var modelsContacts []models.Contact
	if err := database.DB.Find(&modelsContacts).Error; err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération: %w", err)
	}

	contacts := make([]*models.Contact, len(modelsContacts))
	for i, m := range modelsContacts {
		contacts[i] = &models.Contact{
			Model: gorm.Model{ID: m.ID},
			Name:  m.Name,
			Email: m.Email,
		}
	}
	return contacts, nil
}

// Get contact by ID
func (gs *GORMStore) GetByID(id int) (*models.Contact, error) {
	var m models.Contact
	if err := database.DB.First(&m, id).Error; err != nil {
		return nil, storage.ErrContactNotFound(id)
	}
	return &models.Contact{
		Model: gorm.Model{ID: m.ID},
		Name:  m.Name,
		Email: m.Email,
	}, nil
}

// Update a contact
func (gs *GORMStore) Update(id int, newName, newEmail string) error {
	var m models.Contact
	if err := database.DB.First(&m, id).Error; err != nil {
		return storage.ErrContactNotFound(id)
	}

	if newName != "" {
		m.Name = newName
	}
	if newEmail != "" {
		m.Email = newEmail
	}

	if err := database.DB.Save(&m).Error; err != nil {
		return fmt.Errorf("erreur lors de la mise à jour: %w", err)
	}
	return nil
}

// Delete a contact
func (gs *GORMStore) Delete(id int) error {
	if err := database.DB.Delete(&models.Contact{}, id).Error; err != nil {
		return storage.ErrContactNotFound(id)
	}
	return nil
}