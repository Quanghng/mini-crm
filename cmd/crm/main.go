package main

import (
	"mini-crm/internal/app"
	"mini-crm/internal/database"
	"mini-crm/internal/repository"
)


func main() {
	database.ConnectDB()

	store := repository.NewGORMStore()
	app.Run(store)
}
