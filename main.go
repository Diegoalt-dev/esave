package main

import (
	"esave/database"
	"esave/router"
)

func main() {
	database.InitializeDatabase()
	router := router.SetupRouter()
	router.Run(":8080")
}
