package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strings"
	"time"
	"uptime-api/m/v2/internal/core/util"
)

type Response struct {
	Headers   map[string]string `json:"headers,omitempty"`
	Body      string            `json:"body,omitempty"`
	Error     string            `json:"error,omitempty"`
	Region    string            `json:"region"`
	Latency   int64             `json:"latency"`
	Timestamp time.Time         `json:"timestamp"`
	Status    int               `json:"status,omitempty"`
	Timing    util.Timing       `json:"timing"`
}

type HttpCheckerRequest struct {
	Headers []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"headers,omitempty"`
	URL           string            `json:"url"`
	MonitorID     int64             `json:"monitorId"`
	Method        string            `json:"method"`
	Status        string            `json:"status"`
	Body          string            `json:"body"`
	Trigger       string            `json:"trigger,omitempty"`
	RawAssertions []json.RawMessage `json:"assertions,omitempty"`
	CronTimestamp int64             `json:"cronTimestamp"`
	Timeout       int64             `json:"timeout"`
	DegradedAfter int64             `json:"degradedAfter,omitempty"`
}

func performHttp(ctx context.Context, client *http.Client, inputData HttpCheckerRequest) (Response, error) {
	b := []byte(inputData.Body)
	if inputData.Method == http.MethodPost {
		for _, header := range inputData.Headers {
			if header.Key == "Content-Type" && header.Value == "application/octet-stream" {
				//  split the body by comma and convert it to bytes it's data url base64
				data := strings.Split(inputData.Body, ",")
				if len(data) == 2 {
					decoded, err := base64.StdEncoding.DecodeString(data[1])
					if err != nil {
						return Response{}, fmt.Errorf("error while decoding base64: %w", err)
					}

					b = decoded

				}
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, inputData.Method, inputData.URL, bytes.NewReader(b))
	if err != nil {
		return Response{}, fmt.Errorf("unable to create req: %w", err)
	}
	req.Header.Set("User-Agent", "Entries/1.0")
	for _, header := range inputData.Headers {
		if header.Key != "" {
			req.Header.Set(header.Key, header.Value)
		}
	}

	if inputData.Method != http.MethodGet {
		head := req.Header
		_, ok := head["Content-Type"]
		if !ok {
			// by default we set the content type to application/json if it's a POST request
			req.Header.Set("Content-Type", "application/json")
		}
	}

	timing := util.Timing{}

	trace := &httptrace.ClientTrace{
		DNSStart:          func(_ httptrace.DNSStartInfo) { timing.DnsStart = time.Now().UTC().UnixMilli() },
		DNSDone:           func(_ httptrace.DNSDoneInfo) { timing.DnsDone = time.Now().UTC().UnixMilli() },
		ConnectStart:      func(_, _ string) { timing.ConnectStart = time.Now().UTC().UnixMilli() },
		ConnectDone:       func(_, _ string, _ error) { timing.ConnectDone = time.Now().UTC().UnixMilli() },
		TLSHandshakeStart: func() { timing.TlsHandshakeStart = time.Now().UTC().UnixMilli() },
		TLSHandshakeDone:  func(_ tls.ConnectionState, _ error) { timing.TlsHandshakeDone = time.Now().UTC().UnixMilli() },
		GotConn: func(_ httptrace.GotConnInfo) {
			timing.FirstByteStart = time.Now().UTC().UnixMilli()
		},
		GotFirstResponseByte: func() {
			timing.FirstByteDone = time.Now().UTC().UnixMilli()
			timing.TransferStart = time.Now().UTC().UnixMilli()
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	start := time.Now()

	response, err := client.Do(req)
	timing.TransferDone = time.Now().UTC().UnixMilli()
	latency := time.Since(start).Milliseconds()
	timing.Latency = latency
	if err != nil {

		var urlErr *url.Error
		if errors.As(err, &urlErr) && urlErr.Timeout() {
			return Response{
				Latency:   latency,
				Timing:    timing,
				Timestamp: start.UTC(),
				Error:     fmt.Sprintf("Timeout after %d ms", latency),
			}, nil
		}

		return Response{}, fmt.Errorf("error with monitorURL %s: %w", inputData.URL, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Response{
			Latency:   latency,
			Timing:    timing,
			Timestamp: start.UTC(),
			Error:     fmt.Sprintf("Cannot read response body: %s", err.Error()),
		}, fmt.Errorf("error with monitorURL %s: %w", inputData.URL, err)
	}

	headers := make(map[string]string)
	for key := range response.Header {
		headers[key] = response.Header.Get(key)
	}

	return Response{
		Timestamp: start.UTC(),
		Status:    response.StatusCode,
		Headers:   headers,
		Timing:    timing,
		Latency:   latency,
		Body:      string(body),
	}, nil
}
