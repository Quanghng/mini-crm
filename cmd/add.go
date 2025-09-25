package cmd

import (
	"bufio"
	"fmt"
	"mini-crm/internal/models"
	"mini-crm/internal/utils"
	"os"

	"github.com/spf13/cobra"
)

var (
	newName 	string
	newEmail string
)

func init() {
	addCmd.Flags().StringVarP(&newName, "name", "n", "", "Nom du contact")
	addCmd.Flags().StringVarP(&newEmail, "email", "e", "", "Email du contact")
}

var addCmd = &cobra.Command {
	Use: "add",
	Short: "Ajouter un contact",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		// Si le flag n'est pas fourni, demander à l'utilisateur
		if newName == "" {
			fmt.Print("Nom: ")
			newName = utils.ReadLine(reader)
		}

		if newEmail == "" {
			fmt.Print("Email: ")
			newEmail = utils.ReadLine(reader)
		}

		contact := &models.Contact{
			Name:  newName,
			Email: newEmail,
		}

		if err := store.Add(contact); err != nil {
			fmt.Println("Erreur:", err)
			return
		}

		fmt.Printf("Contact '%s' ajouté avec ID %d\n", contact.Name, contact.ID)
	},
}

