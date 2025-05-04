package models

import "time"

type Domain struct {
	Domain       string    `json:"domain" bson:"domain"`
	Verification string    `json:"verification" bson:"verification"`
	VerifiedAt   time.Time `json:"verified_at" bson:"verified_at"`
	RemovedAt    time.Time `json:"removed_at" bson:"removed_at"`
}

type APIKey struct {
	APIKey    string    `json:"key" bson:"key"`
	Domains   []Domain  `json:"domains" bson:"domains"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	RemovedAt time.Time `json:"removed_at" bson:"removed_at"`
}
