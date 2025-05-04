package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"mcbulazs/email-service/internal/config"
	"mcbulazs/email-service/internal/logging"
)

type mongoDB struct {
	Client *mongo.Client
}

// Connect to MongoDB
func Connect() (*mongoDB, error) {
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
	return &mongoDB{Client: client}, nil
}

// Disconnect from MongoDB
func (client *mongoDB) Disconnect() {
	err := client.Client.Disconnect(context.Background())
	if err != nil {
		logging.ErrorLogger.Fatalf("Could not disconnect from MongoDB: %v", err)
	}
	logging.InfoLogger.Println("Disconnected from MongoDB!")
}

// InsertData inserts a document into a collection
func (client *mongoDB) InsertData(collectionName string, data any) error {
	collection := client.Client.Database(config.AppConfig.Mongo.DATABASE).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}

func (client *mongoDB) GetData(collectionName string, filter any) any {
	collection := client.Client.Database(config.AppConfig.Mongo.DATABASE).Collection(collectionName)
	result := collection.FindOne(context.Background(), filter)
	return result
}
