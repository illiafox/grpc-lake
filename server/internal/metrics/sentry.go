package metrics

import (
	"errors"

	"github.com/getsentry/sentry-go"
	"server/internal/config"
)

// SetupSentry setups sentry and returns a function to Flush buffered events
func SetupSentry(cfg config.Sentry) (func() error, error) {

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
