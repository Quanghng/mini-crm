package handler

import "fmt"

type Contact struct {
	Id    int
	Nom   string
	Email string
}

// Variables globales
var contacts = make(map[int]*Contact)
var currentId int = 1

func (c *Contact) AddContact(nom, email string) {
	contacts[currentId] = c
	currentId++
	fmt.Printf("Contact ID: %d | Nom: %s | Email: %s ajouté avec succès.\n", c.Id, c.Nom, c.Email)
}

func GetContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact disponible.")
		return
	}

	fmt.Println("Liste des contacts :")
	for _, c := range contacts {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", c.Id, c.Nom, c.Email)
	}
}

func (c *Contact) DeleteContact(id int) {
	delete(contacts, c.Id)
	fmt.Println("Contact supprimé avec succès.")
}

func (c *Contact) UpdateContact(id int, nom, email string) {
	if nom != "" {
		c.Nom = nom
	}
	if email != "" {
		c.Email = email
	}
	fmt.Println("Contact mis à jour avec succès.")
}

func GetContactById(id int) *Contact {
	if c, ok := contacts[id]; ok {
		return c
	}
	return nil
}

func CurrentId() int {
	return currentId
}