package database

import (
	"log"
	"mini-crm/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("contacts.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Echec de la connexion a la base de donnees : %v", err)
	}

	log.Printf("Connexion a la base de donnees SQLite reussie !")

	err = DB.AutoMigrate(&models.Contact{})
	if err != nil {
		log.Fatalf("Echec de la migration de la base de donnees : %v", err)
	}
	log.Printf("Migration de la base de donnees reussie !")
}