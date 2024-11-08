package provider

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"uptime-api/m/v2/internal/core/domain"
	"uptime-api/m/v2/internal/core/util"
)

type Embed struct {
	Color string `json:"color,omitempty"`
}

type MessageData struct {
	Content string `json:"content,omitempty"`
}

const (
	Alerting  string = "#fd5c63"
	Recovered string = "#32de84"
	Degraded  string = "#f7b731"
)

type DiscordNotificationProvider struct {
	WebhookUrl string
	ApiKey     string
}

func (d *DiscordNotificationProvider) SendAlerting(m *domain.Monitor, statusCode int, message string) error {
	log.Printf("Sending alerting message to discord for monitor with id: %d\n", m.Id)
	md := MessageData{
		Content: fmt.Sprintf(`**🚨 Alert [%s]**

Status Code: %d
Message: %s`, m.Title, statusCode, message),
	}

	return d.Send(md)
}

func (d *DiscordNotificationProvider) SendDegraded(m *domain.Monitor, statusCode int, timing util.Timing) error {
	log.Printf("Sending degraded message to discord for monitor with id: %d\n", m.Id)
	md := MessageData{
		Content: fmt.Sprintf(`**⚠️ Degraded [%s]**

We received a successful status code: %d, but the latency degraded.

Expected Response Time: Less than %d ms
Actual Response Time: %d ms`, m.Title, statusCode, m.DegradedAfter, timing.Latency),
	}

	return d.Send(md)
}

func (d *DiscordNotificationProvider) SendRecovered(m *domain.Monitor, statusCode int, timing util.Timing) error {
	log.Printf("Sending recovered message to discord for monitor with id: %d\n", m.Id)
	md := MessageData{
		Content: fmt.Sprintf(`**✅ Recovered [%s]**

The latency is back to normal. We received a successful status code: %d

Expected Response Time: Less than %d ms
Received Response Time: %d ms
`, m.Title, statusCode, m.DegradedAfter, timing.Latency),
	}

	return d.Send(md)
}

func NewDiscordNotificationProvider(data map[string]interface{}) *DiscordNotificationProvider {
	wh := data["webhook_url"].(string)
	return &DiscordNotificationProvider{
		WebhookUrl: wh,
	}
}

func (d *DiscordNotificationProvider) Send(messageData MessageData) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   45 * time.Second,
		Transport: tr,
	}

	b, err := json.Marshal(messageData)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, d.WebhookUrl, bytes.NewReader(b))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)

	if err != nil {
		log.Println("error sending discord message: ", err)
		return err
	}

	body, _ := io.ReadAll(response.Body)

	log.Printf("Discord response status: %v\n", response.Status)
	log.Printf("Discord response body: %v\n", string(body))

	return nil
}
