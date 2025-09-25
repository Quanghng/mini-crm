package storage

import (
	"fmt"
	"mini-crm/internal/models"
)

// Storer est un CONTRAT de stockage
// Il définit un ensemble de comportements (méthodes) que tout type
// de stockage doit respecter. On ne se soucie par du comment c'est fait
// (en mémoire, fichier, BDD...) seulement de ce qui peut être fait
type Storer interface {
	Add(contact *models.Contact) error
	GetAll() ([]*models.Contact, error)
	GetByID(id int) (*models.Contact, error)
	Update(id int, newName, newEmail string) error
	Delete(id int) error
}

// Creer une erreur personalise
var ErrContactNotFound = func(id int) error {
	return fmt.Errorf("Contact avec l'ID %d non trouve", id)
}