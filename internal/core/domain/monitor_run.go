package domain

import "time"

type MonitorRun struct {
	Id                    int64         `json:"id" db:"id"`
	MonitorId             int64         `json:"monitor_id" db:"id"`
	Status                MonitorStatus `json:"status" db:"status"`
	DnsStartedAt          time.Time     `json:"dns_started_at" db:"dns_started_at"`
	DnsEndedAt            time.Time     `json:"dns_ended_at" db:"dns_ended_at"`
	ConnectStartedAt      time.Time     `json:"connect_started_at" db:"connect_started_at"`
	ConnectEndedAt        time.Time     `json:"connect_ended_at" db:"connect_ended_at"`
	TlsHandshakeStartedAt time.Time     `json:"tls_handshake_started_at" db:"tls_handshake_started_at"`
	TlsHandshakeEndedAt   time.Time     `json:"tls_handshake_ended_at" db:"tls_handshake_ended_at"`
	Latency               int64         `json:"latency" db:"latency"`
	CreatedAt             time.Time     `json:"created_at" db:"created_at"`
	RanAt                 time.Time     `json:"ran_at" db:"ran_at"`
}
