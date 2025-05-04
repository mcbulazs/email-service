package services

import (
	"fmt"
	"net"
	"slices"

	"mcbulazs/email-service/internal/config"
)

// errors
var (
	ErrDNSRecordNotFound = fmt.Errorf("DNS record not found")
)

type VerifyRepoInterface interface {
	GetVerifiactionCode(domain, apiKey string) (string, error)
	UpdateDomainVerifiedAt(domain, apiKey string) error
}

type VerifyService struct {
	Repo VerifyRepoInterface
}

func NewVerifyService(repo VerifyRepoInterface) *VerifyService {
	return &VerifyService{
		Repo: repo,
	}
}

func (s *VerifyService) VerifyDomain(domain, apiKey string) error {
	verifiaction, err := s.Repo.GetVerifiactionCode(domain, apiKey)
	if err != nil {
		return err
	}
	err = verifyDNSRecord(domain, verifiaction)
	if err != nil {
		return err
	}

	err = s.Repo.UpdateDomainVerifiedAt(domain, apiKey)
	if err != nil {
		return err
	}

	return nil
}

func verifyDNSRecord(domain string, verification string) error {
	records, err := net.LookupTXT(domain)
	if err != nil {
		return ErrDNSRecordNotFound
	}

	if slices.Contains(records, fmt.Sprintf("%s_email-service=%s", config.AppConfig.DNSPrefix, verification)) {
		return nil // verified
	}

	return ErrDNSRecordNotFound // not verified
}
