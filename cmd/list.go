package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister tous les contacts",
	Run: func(cmd *cobra.Command, args []string) {
		contacts, err := store.GetAll()
		if err != nil {
			fmt.Println("Erreur:", err)
			return
		}

		if len(contacts) == 0 {
			fmt.Println("Aucun contact trouvé.")
			return
		}

		fmt.Println("--- Liste des contacts ---")
		for _, c := range contacts {
			fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Name, c.Email)
		}
	},
}
