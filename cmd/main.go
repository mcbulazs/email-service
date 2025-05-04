package main

import (
	"context"
	"net/http"

	"mcbulazs/email-service/internal/config"
	"mcbulazs/email-service/internal/handlers"
	"mcbulazs/email-service/internal/logging"
	"mcbulazs/email-service/internal/mongo"
	"mcbulazs/email-service/internal/repositories"
	"mcbulazs/email-service/internal/services"
)

func main() {
	// config
	config.Load()
	// logger
	logging.Init()

	// mongo
	db, err := mongo.Connect()
	if err != nil {
		logging.ErrorLogger.Fatalf("Could not connect to MongoDB: %v", err)
		return
	}
	defer func() {
		err := db.Disconnect(context.Background())
		if err != nil {
			logging.ErrorLogger.Printf("Could not disconnect from MongoDB: %v", err)
		}
	}()

	// repositories
	verifyRepo := repositories.NewVerifyRepository(db)

	// services
	verifyService := services.NewVerifyService(verifyRepo)

	// controllers
	controller := handlers.Controller{
		Service: verifyService,
	}
	http.HandleFunc("POST /verify", controller.InitVerifyHandler)

	// start server
	logging.InfoLogger.Println("Server listening on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		logging.ErrorLogger.Fatalf("Could not start server: %v", err)
	}
}
