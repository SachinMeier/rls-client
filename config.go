package rls

// ClienntConfig contains the configurable values used by an RLS
type Config struct {
	BaseURL    string
	AccountID  string
	FeeLimit   int64
	credential string
}

// NewConfig creates a new Config
func NewConfig(baseURL string, apiKey string, accountID string, feeLimit int64) Config {
	return Config{
		BaseURL:    baseURL,
		credential: createCredential(apiKey),
		AccountID:  accountID,
		FeeLimit:   feeLimit,
	}
}
