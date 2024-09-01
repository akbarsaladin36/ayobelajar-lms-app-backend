package main

import (
	"ayobelajar-app-backend/config"
	"ayobelajar-app-backend/routes"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	config.ConnectDB()
	routes.ConnectRoutes()
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed to load .env file")
	}
}
