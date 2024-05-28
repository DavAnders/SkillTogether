package main

import (
	"log"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}



	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer database.Close()
}
