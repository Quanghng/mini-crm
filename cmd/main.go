package main

import (
	"bufio"
	"flag"
	"fmt"
	handler "mini-crm/internal"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Mode flag
	var addContact bool
	var nom string
	var email string

	// Définition des flags
	flag.BoolVar(&addContact, "addContact", false, "Ajouter un contact")
	flag.StringVar(&nom, "nom", "", "Nom du contact")
	flag.StringVar(&email, "email", "", "Email du contact")
	flag.Parse()

	if addContact {
		if nom == "" || email == "" {
			fmt.Println("Erreur : --nom et --email sont obligatoires avec --addContact")
			return
		}
		newContact := &handler.Contact{Id: handler.CurrentId(), Nom: nom, Email: email}
		newContact.AddContact(nom, email)
		return
	}

	// Mode Menu
	scanner := bufio.NewReader(os.Stdin)

	for {
		// Afficher le menu principal
		fmt.Println("----- Mini CRM -----")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister tous les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter l'application")
		fmt.Print("Sélectionnez votre option: ")

		// Lire entrée utilisateur
		optString, _ := scanner.ReadString('\n')
		optString = strings.TrimSpace(optString)

		opt, err := strconv.Atoi(optString)
		if err != nil {
			fmt.Println("Veuillez entrer un nombre valide.")
			continue
		}

		switch opt {
		case 1:
			// Ajouter un contact
			fmt.Print("Nom: ")
			nom, _ := scanner.ReadString('\n')
			nom = strings.TrimSpace(nom)

			fmt.Print("Email: ")
			email, _ := scanner.ReadString('\n')
			email = strings.TrimSpace(email)

			// Pointeur vers le nouveau struct
			newContact := &handler.Contact{Id: handler.CurrentId(), Nom: nom, Email: email}
			newContact.AddContact(nom, email)
		case 2:
			handler.GetContacts()
		case 3:
			// Supprimer un contact
			fmt.Print("ID du contact à supprimer: ")
			idStr, _ := scanner.ReadString('\n')
			idStr = strings.TrimSpace(idStr)

			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("ID invalide.")
				continue
			}

			if c := handler.GetContactById(id); c != nil {
				c.DeleteContact(id)
			} else {
				fmt.Println("Aucun contact trouvé avec cet ID.")
			}
		case 4:
			// Mettre à jour un contact
			fmt.Print("ID du contact à mettre à jour: ")
			idStr, _ := scanner.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("ID invalide.")
				continue
			}

			fmt.Print("Nouveau nom: ")
			nom, _ := scanner.ReadString('\n')
			nom = strings.TrimSpace(nom)

			fmt.Print("Nouvel email: ")
			email, _ := scanner.ReadString('\n')
			email = strings.TrimSpace(email)

			if c := handler.GetContactById(id); c != nil {
				c.UpdateContact(id, nom, email)
			} else {
				fmt.Println("Aucun contact trouvé avec cet ID.")
			}
		case 5:
			fmt.Println("Fermeture du programme. Au revoir !")
			return
		default:
			fmt.Println("Option invalide, réessayez.")
		}
	}

}
