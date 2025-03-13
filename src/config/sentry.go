package config

import (
	"log"

	"github.com/getsentry/sentry-go"
)

func InitSentry() {
	config := GetConfig()

	if config.Sentry.Environment != "local" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:              config.Sentry.Dsn,
			TracesSampleRate: config.Sentry.TracesSampleRate,
			Release:          config.Sentry.Release,
			Environment:      config.Sentry.Environment,
		})

		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		}
	}
}
