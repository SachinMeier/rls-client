package rls

// Config contains the configurable values used by an RLS
type Config struct {
	BaseURL       string
	credential    string
	AccountID     string
	WebhookSecret string
}

// NewConfig creates a new Config
func NewConfig(baseURL string, apiKey string, accountID string, webhookSecret string) *Config {
	return &Config{
		BaseURL:       baseURL,
		credential:    createCredential(apiKey),
		AccountID:     accountID,
		WebhookSecret: webhookSecret,
	}
}
