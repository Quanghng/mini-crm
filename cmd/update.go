package cmd

import (
	"bufio"
	"fmt"
	"mini-crm/internal/utils"
	"os"

	"github.com/spf13/cobra"
)

var (
	updateID   int
	updateName string
	updateEmail string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Mettre à jour un contact",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		if updateID == 0 {
			fmt.Print("ID du contact à mettre à jour: ")
			updateID = utils.ReadInt(reader)
		}

		contact, err := store.GetByID(updateID)
		if err != nil {
			fmt.Println("Erreur:", err)
			return
		}

		if updateName == "" {
			fmt.Printf("Nouveau nom (%s) : ", contact.Name)
			updateName = utils.ReadLine(reader)
		}

		if updateEmail == "" {
			fmt.Printf("Nouvel email (%s) : ", contact.Email)
			updateEmail = utils.ReadLine(reader)
		}

		if err := store.Update(updateID, updateName, updateEmail); err != nil {
			fmt.Println("Erreur:", err)
			return
		}

		fmt.Println("Contact mis à jour avec succès.")
	},
}

func init() {
	updateCmd.Flags().IntVarP(&updateID, "id", "i", 0, "ID du contact")
	updateCmd.Flags().StringVarP(&updateName, "name", "n", "", "Nouveau nom")
	updateCmd.Flags().StringVarP(&updateEmail, "email", "e", "", "Nouvel email")
}

