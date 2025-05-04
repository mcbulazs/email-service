package repositories

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"

	"mcbulazs/email-service/internal/config"
	"mcbulazs/email-service/internal/logging"
	"mcbulazs/email-service/internal/models"
)

type VerifyRepository struct {
	DB *mongo.Client
}

func NewVerifyRepository(db *mongo.Client) *VerifyRepository {
	return &VerifyRepository{
		DB: db,
	}
}

var (
	ErrAPIKeyNotFound = fmt.Errorf("API key not found")
	ErrDomainNotFound = fmt.Errorf("domain not found")
)

func (db *VerifyRepository) GetVerifiactionCode(domain, apiKey string) (string, error) {
	var apiKeyData models.APIKey
	filter := map[string]string{
		"api_key": apiKey,
	}

	// Perform the MongoDB query to retrieve the APIKey document
	err := db.DB.Database(config.AppConfig.Mongo.DATABASE).
		Collection("api_keys").
		FindOne(context.Background(), filter).Decode(&apiKeyData)
	if err != nil {
		logging.ErrorLogger.Printf("failed to get API key data: %v", err)
		return "", ErrAPIKeyNotFound
	}

	for _, d := range apiKeyData.Domains {
		if d.Domain == domain {
			if d.RemovedAt.IsZero() {
				return d.Verification, nil
			}
		}
	}
	return "", ErrDomainNotFound
}

func (db *VerifyRepository) UpdateDomainVerifiedAt(domain, apiKey string) error {
	// Get the current date and time
	currentTime := time.Now()

	// Define the filter to find the correct API key and domain
	filter := map[string]string{
		"api_key":        apiKey,
		"domains.domain": domain, // Match the domain inside the domains array
	}

	// Define the update operation
	update := map[string]any{
		"$set": map[string]any{
			"domains.$.verified_at": currentTime, // Update the verified_at field of the domain
		},
	}

	// Get the MongoDB collection for "api_keys"
	collection := db.DB.Database(config.AppConfig.Mongo.DATABASE).Collection("api_keys")

	// Perform the update operation on the collection
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logging.ErrorLogger.Printf("failed to update verified_at for domain %s: %v", domain, err)
		return err
	}

	return nil
}
