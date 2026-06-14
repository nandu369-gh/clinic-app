package main

import (
	"clinic-app/config"
	"clinic-app/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file contents into system environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using code fallback defaults")
	}
	// 1.Initialize pgx connection pool
	config.ConnectDatabase()
	defer config.DB.Close() // Safely close connections on stop
	// 2. Setup Routes
	r := routes.SetupRouter()
	// 3. Start Server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
