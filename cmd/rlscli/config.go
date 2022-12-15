package main

import (
	"context"
	"fmt"
	"os"

	"github.com/SachinMeier/rls-client"
)

const msgFailedToLoadConfig string = "failed to load config : %s"

func LoadRLSConfig() (*rls.Config, error) {
	env := os.Getenv("RLS_ENV")
	baseURL := os.Getenv(fmt.Sprintf("%s_URL", env))
	if baseURL == "" {
		return nil, fmt.Errorf("%s_URL not set", env)
	}
	accountID := os.Getenv(fmt.Sprintf("%s_RIVER_ACCOUNT_ID", env))
	if accountID == "" {
		return nil, fmt.Errorf("%s_RIVER_ACCOUNT_ID not set", env)
	}
	apiKey := os.Getenv(fmt.Sprintf("%s_RIVER_API_SECRET", env))
	if apiKey == "" {
		return nil, fmt.Errorf("%s_RIVER_API_SECRET not set", env)
	}
	// optionals
	webhookSecret := os.Getenv(fmt.Sprintf("%s_WEBHOOK_SECRET", env))
	return rls.NewConfig(baseURL, apiKey, accountID, webhookSecret), nil
}

func NewRLSClient(ctx context.Context) (*rls.RLSClient, error) {
	cfg, err := LoadRLSConfig()
	if err != nil {
		return nil, fmt.Errorf(msgFailedToLoadConfig, err)
	}
	return rls.NewRLSClient(ctx, *cfg), nil
}
