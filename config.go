package rls

// ClienntConfig contains the configurable values used by an RLS
type Config struct {
	APIVersion string
	BaseURL    string
	AccountID  string
	FeeLimit   int64
	credential string
}

// NewConfig creates a new Config
func NewConfig(apiVersion string, baseURL string, apiKey string, accountID string, feeLimit int64) *Config {
	return &Config{
		APIVersion: apiVersion,
		BaseURL:    baseURL,
		credential: createCredential(apiKey),
		AccountID:  accountID,
		FeeLimit:   feeLimit,
	}
}
