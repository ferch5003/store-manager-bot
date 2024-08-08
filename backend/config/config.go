package config

import (
	"os"
)

type EnvVars struct {
	// Server Environment.
	Port string

	// Database Environment.
	DSN string

	// Vertex Model environments.
	ProjectTunedID  string
	RegionTuned     string
	EndpointTunedID string

	ProjectFlashID string
	RegionFlash    string
	ModelFlashID   string

	SAPrivateKeyID      string
	SAPrivateKey        string
	SAClientEmail       string
	SAClientID          string
	SAClientX509CertURL string
}

func NewConfigurations() (*EnvVars, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dsn := os.Getenv("DSN")

	projectTunedID := os.Getenv("PROJECT_TUNED_ID")
	regionTuned := os.Getenv("REGION_TUNED")
	endpointTunedID := os.Getenv("ENDPOINT_TUNED_ID")

	projectFlashID := os.Getenv("PROJECT_FLASH_ID")
	regionFlash := os.Getenv("REGION_FLASH")
	modelFlashID := os.Getenv("MODEL_FLASH_ID")

	saPrivateKeyID := os.Getenv("SA_PRIVATE_KEY_ID")
	saPrivateKey := os.Getenv("SA_PRIVATE_KEY")
	saClientEmail := os.Getenv("SA_CLIENT_EMAIL")
	saClientID := os.Getenv("SA_CLIENT_ID")
	saClientX509CertURL := os.Getenv("SA_CLIENT_X509_CERT_URL")

	environment := &EnvVars{
		Port: port,
		DSN:  dsn,

		ProjectTunedID:  projectTunedID,
		RegionTuned:     regionTuned,
		EndpointTunedID: endpointTunedID,

		ProjectFlashID: projectFlashID,
		RegionFlash:    regionFlash,
		ModelFlashID:   modelFlashID,

		SAPrivateKeyID:      saPrivateKeyID,
		SAPrivateKey:        saPrivateKey,
		SAClientEmail:       saClientEmail,
		SAClientID:          saClientID,
		SAClientX509CertURL: saClientX509CertURL,
	}

	return environment, nil
}
