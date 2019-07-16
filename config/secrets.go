package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Secrets ... get environment variable from config
var Secrets SecretParameters

type SecretParameters struct {
	NewrelicTeamID        string `envconfig:"NEWRELIC_TEAM_ID"`
	NewrelicLoginEmail    string `envconfig:"NEWRELIC_LOGIN_EMAIL"`
	NewrelicLoginPassword string `envconfig:"NEWRELIC_LOGIN_PASS"`

	RollbarReadAccessToken  string `envconfig:"ROLLBAR_READ_ACCESS_TOKEN"`
	RollbarWriteAccessToken string `envconfig:"ROLLBAR_WRITE_ACCESS_TOKEN"`

	DatadogAPIKey string `envconfig:"DATADOG_API_KEY"`
	DatadogAppKey string `envconfig:"DATADOT_APP_KEY"`
}

func init() {
	setEnvConfig()
}

// setEnvConfig ... set environment config
func setEnvConfig() {
	if err := envconfig.Process("", &Secrets); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}
}
