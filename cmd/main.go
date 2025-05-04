package main

import (
	"net/http"

	"mcbulazs/email-service/internal/config"
	"mcbulazs/email-service/internal/handlers"
	"mcbulazs/email-service/internal/logging"
	"mcbulazs/email-service/internal/mongo"
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
	defer db.Disconnect()

	// controllers
	controller := handlers.Controller{
		Repo: db,
	}
	http.HandleFunc("POST /verify", controller.InitVerifyHandler)

	// start server
	logging.InfoLogger.Println("Server listening on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		logging.ErrorLogger.Fatalf("Could not start server: %v", err)
	}
}
