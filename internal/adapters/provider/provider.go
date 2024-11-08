package provider

import "uptime-api/m/v2/internal/core/ports"

var providerToFunction = map[string]func(data map[string]interface{}) ports.NotificationProvider{
	"discord": func(data map[string]interface{}) ports.NotificationProvider {
		return NewDiscordNotificationProvider(data)
	},
	//"email": func(data map[string]interface{}) ports.NotificationProvider {
	//	return NewEmailNotificationProvider(data)
	//},
}

func NewNotificationProvider(provider string, data map[string]interface{}) (ports.NotificationProvider, error) {
	p := providerToFunction[provider](data)
	return p, nil
}
