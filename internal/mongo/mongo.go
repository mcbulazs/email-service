package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"mcbulazs/email-service/internal/config"
	"mcbulazs/email-service/internal/logging"
)

// Connect to MongoDB
func Connect() (*mongo.Client, error) {
	// Set client options
	var clientOptions *options.ClientOptions
	if config.AppConfig.Mongo.USER != "" {
		clientOptions = options.Client().ApplyURI(fmt.Sprintf(
			"mongodb://%s:%s@%s",
			config.AppConfig.Mongo.USER,
			config.AppConfig.Mongo.PASS,
			config.AppConfig.Mongo.URL))
	} else {
		clientOptions = options.Client().ApplyURI(fmt.Sprintf(
			"mongodb://%s",
			config.AppConfig.Mongo.URL))
	}

	// Connect to MongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		logging.ErrorLogger.Fatalf("Could not connect to MongoDB: %v", err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logging.ErrorLogger.Fatalf("Could not ping MongoDB: %v", err)
		return nil, err
	}

	logging.InfoLogger.Println("Successfully connected to MongoDB!")
	return client, nil
}
