package cmd

import (
	"fmt"
	"log"
	"mini-crm/internal/config"
	"mini-crm/internal/database"
	"mini-crm/internal/repository"
	"mini-crm/internal/storage"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var store storage.Storer

var rootCmd = &cobra.Command{
	Use:   "minicrm",
	Short: "Mini CRM - Gestion simple de contacts",
	Long:  "Mini CRM est une application CLI permettant d'ajouter, lister, mettre à jour et supprimer des contacts avec plusieurs backends de stockage.",
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig, initStore)

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
}

func initStore() {
	storageType := viper.GetString("storage.type")
	switch storageType {
	case "gorm":
		database.ConnectDB()
		store = repository.NewGORMStore()
	case "json":
		var err error
		store, err = storage.NewJsonStore("./data/contacts.json")
		if err != nil {
			fmt.Printf("Error initializing JSON store: %v\n", err)
			os.Exit(1)
		}
	case "memory":
		store = storage.NewMemoryStore()
	default:
		log.Fatalf("Type de stockage inconnu: %s", storageType)
	} 
}