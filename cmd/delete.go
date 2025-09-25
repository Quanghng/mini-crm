package cmd

import (
	"bufio"
	"fmt"
	"mini-crm/internal/utils"
	"os"

	"github.com/spf13/cobra"
)

var deleteID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprimer un contact",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		if deleteID == 0 {
			fmt.Print("ID du contact à supprimer: ")
			deleteID = utils.ReadInt(reader)
		}

		if err := store.Delete(deleteID); err != nil {
			fmt.Println("Erreur:", err)
			return
		}

		fmt.Println("Contact supprimé avec succès.")
	},
}

func init() {
	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "ID du contact à supprimer")
}

