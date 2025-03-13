package models

type SentryConfig struct {
	Dsn              string
	TracesSampleRate float64
	Release          string
	Environment      string
}

type Config struct {
	Sentry SentryConfig
}