package domain

type UptimeEntry struct {
	Key    string  `json:"key" db:"key"`
	Ok     int64   `json:"ok" db:"ok"`
	Total  int64   `json:"total" db:"total"`
	Uptime float64 `json:"uptime" db:"uptime"`
}
