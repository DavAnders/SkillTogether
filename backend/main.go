package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/DavAnders/SkillTogether/backend/internal/auth"
	"github.com/DavAnders/SkillTogether/backend/internal/handler"
	"github.com/gin-gonic/gin"
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

	// Test query to check if database is connected and accessible
	rows, err := database.Query("SELECT * FROM skills LIMIT 1")
	if err != nil {
    	log.Fatalf("Error querying skills table: %v", err)
	}
	defer rows.Close()
	log.Println("Accessed skills table successfully")


	defer database.Close()

	router := gin.Default()
	dbQueries := db.New(database)
	handler := handler.NewHandler(dbQueries)
	authHandler := auth.NewAuthHandler(dbQueries)

	// Skill routes
	router.GET("/skills", handler.GetAllSkills)
	router.GET("/skills/:id", handler.GetSkill)
	router.POST("/skills", handler.AddSkill)
	router.PUT("/skills/:id", handler.UpdateSkill)
	router.DELETE("/skills/:id", handler.DeleteSkill)

	// User routes
	router.PUT("/users/:discord_id", handler.UpdateUser)
	router.GET("/users/:discord_id", handler.GetUser)
	router.DELETE("/users/:discord_id", handler.DeleteUser)

	// Discord OAuth2 routes
	router.GET("/callback", authHandler.DiscordCallbackHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
        // Service connections
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    // Wait for interrupt signal to shut down the server with a timeout of 5 seconds
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("Shutting down server...")

    // The context is used to inform the server it has 5 seconds to finish the request it is currently handling
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }

}
