package inputs

type CreateNotificationInput struct {
	Title        string                 `json:"title" validate:"required"`
	Description  string                 `json:"description"`
	Provider     string                 `json:"provider" validate:"required"`
	ProviderData map[string]interface{} `json:"provider_data" validate:"required"`
	MonitorIds   []int64                `json:"monitor_ids"`
}
