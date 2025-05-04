package handlers

import (
	"encoding/json"
	"net/http"
)

type VerifyRequest struct {
	Domain string `json:"domain"`
	APIKey string `json:"api_key"`
}

type VerifyServiceInterface interface {
	VerifyDomain(domain, apiKey string) error
}

type Controller struct {
	Service VerifyServiceInterface
}

func (c *Controller) InitVerifyHandler(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Domain == "" || req.APIKey == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	err = c.Service.VerifyDomain(req.Domain, req.APIKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
}
