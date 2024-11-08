package provider

import (
	"os"
	"strings"
	"uptime-api/m/v2/internal/core/domain"
)

type EmailNotificationProvider struct {
	ResendApiKey string
	Recipients   []string
}

func (e *EmailNotificationProvider) SendAlerting(m *domain.Monitor, statusCode int, message string) error {
	return nil
}

func (e *EmailNotificationProvider) SendDegraded(m *domain.Monitor, statusCode int, message string) error {
	return nil
}

func (e *EmailNotificationProvider) SendRecovered(m *domain.Monitor, statusCode int, message string) error {
	return nil
}

func NewEmailNotificationProvider(data map[string]interface{}) *EmailNotificationProvider {
	return &EmailNotificationProvider{
		ResendApiKey: os.Getenv("RESEND_API_KEY"),
		Recipients:   strings.Split(data["recipients"].(string), ","),
	}
}
