package metrics

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"io"
	"net/http"
	"os"
	"server/internal/config"
)

// HealthCheck checks if sentry is available
// https://docs.sentry.io/product/relay/monitoring/#health-checks
func HealthCheck(url string) error {

	type Response struct {
		IsHealthy bool `json:"is_healthy"`
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http.Get: %w", err)
	}

	var response map[string]any
	if err = json.NewDecoder(io.TeeReader(resp.Body, os.Stdout)).Decode(&response); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	//if !response.IsHealthy {
	//	return errors.New("health check failed")
	//}

	fmt.Println(response)
	return fmt.Errorf("health check failed")
	return nil
}

// https://docs.sentry.io/product/relay/monitoring/#health-checks
const (
	SentryHealthCheckLive  = "/api/relay/healthcheck/live/"
	SentryHealthCheckReady = "/api/relay/healthcheck/ready/"
)

// SetupSentry setups sentry and returns a function to Flush buffered events
func SetupSentry(cfg config.Sentry) (func() error, error) {

	//// Health Checks
	//if false {
	//	if err := HealthCheck(cfg.DSN + SentryHealthCheckLive); err != nil {
	//		return nil, fmt.Errorf("health check live: %w", err)
	//	}
	//
	//	if err := HealthCheck(cfg.DSN + SentryHealthCheckReady); err != nil {
	//		return nil, fmt.Errorf("health check ready: %w", err)
	//	}
	//}

	// Init Sentry

	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn:   cfg.DSN,
		Debug: cfg.Debug,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: cfg.TracesSampleRate,

		// Suggest that all captures are errors
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			event.Level = sentry.LevelError
			return event
		},
	})

	if err != nil {
		return nil, err
	}
	// //

	flush := func() error {
		if !sentry.Flush(cfg.FlushTimeout) {
			return errors.New("sentry flush timeout")
		}
		return nil
	}

	return flush, nil
}
