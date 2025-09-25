package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	// Des valeurs par defaut
	viper.SetDefault("storage.type", "gorm")
	viper.SetDefault("database.name", "contacts.db")
	viper.SetDefault("database.dsn", "./data/contacts.db")
	viper.SetDefault("app.environment", "development")

	fmt.Println("Valeurs par defaut definies.")


	viper.SetConfigName("config")

	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./cmd/config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Fichier de configuration non trouve")
		} else {
			log.Fatalf("Erreur lors de la lecture du fichier de configuration: %v", err)
		}
	} else {
		// log.Println("Fichier de configuration charge avec succes.")
	}

	fmt.Printf("Environmen: %s\n", viper.GetString("app.environment"))
	fmt.Printf("Storage: %s\n", viper.GetString("storage.type"))

	var dbName, dbDSN string
	switch viper.GetString("storage.type") {
		case "gorm":
			dbName = viper.GetString("database.gorm.name")
			dbDSN = viper.GetString("database.gorm.dsn")
		case "json":
			dbName = viper.GetString("database.json.name")
			dbDSN = viper.GetString("database.json.dsn")
		default:
			log.Fatalf("Type de stockage inconnu: %s", viper.GetString("storage.type"))
	}
	
	fmt.Printf("Database: %s\n", dbName)
	fmt.Printf("Data source: %s\n", dbDSN)
}