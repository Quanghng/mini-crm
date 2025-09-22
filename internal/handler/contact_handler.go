package handler

import (
	"TP/internal/domain"
	"fmt"
)

// Variables globales
var contacts = make(map[int]domain.Contact)
var currentId int = 1

func AddContact(nom, email string) {
	newContact := domain.Contact{
		Id:		 currentId,
		Nom:   nom,
		Email: email,
	}
	contacts[currentId] = newContact
	currentId++
	fmt.Printf("Contact ID: %d | Nom: %s | Email: %s ajouté avec succès.\n", newContact.Id, newContact.Nom, newContact.Email)
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

func DeleteContact(id int) {
	if _, ok := contacts[id]; ok {
		delete(contacts, id)
		fmt.Println("Contact supprimé avec succès.")
	} else {
		fmt.Println("Aucun contact trouvé avec cet ID.")
	}
}

func UpdateContact(id int, nom, email string) {
	contact, ok := contacts[id]
	if !ok {
		fmt.Println("Aucun contact trouvé avec cet ID.")
		return
	}

	if nom != "" {
		contact.Nom = nom
	}
	if email != "" {
		contact.Email = email
	}

	contacts[id] = contact
	fmt.Println("Contact mis à jour avec succès.")
}
