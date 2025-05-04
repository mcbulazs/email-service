package handlers

import (
	"encoding/json"
	"net/http"

	"mcbulazs/email-service/internal/interfaces"
)

type VerifyRequest struct {
	Domain string `json:"domain"`
	APIKey string `json:"api_key"`
}

type Controller struct {
	Repo interfaces.DB
}

func (c *Controller) InitVerifyHandler(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Domain == "" || req.APIKey == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// TODO> validate API key

	// TODO: validate Domain (owned by user)

	// TODO: store info in mongo
}
