package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/SachinMeier/rls-client"
	cli "github.com/urfave/cli"
)

const (
	rlsEnvKey           = "RLS_ENV"
	rlsTLSPathKey       = "RLS_TLSPATH"
	rlsURLKey           = "_URL"
	rlsAccountIDKey     = "_RIVER_ACCOUNT_ID"
	rlsAPISecretKey     = "_RIVER_API_SECRET"
	rlsWebhookSecretKey = "_WEBHOOK_SECRET"
	rlsHeadersKey       = "_HEADERS"
)

const msgFailedToLoadConfig string = "failed to load config : %s"

func LoadRLSConfig() (*rls.Config, error) {
	env := os.Getenv(rlsEnvKey)
	baseURL := os.Getenv(env + rlsURLKey)
	if baseURL == "" {
		return nil, fmt.Errorf("%s not set", env+rlsURLKey)
	}
	accountID := os.Getenv(env + rlsAccountIDKey)
	if accountID == "" {
		return nil, fmt.Errorf("%s not set", env+rlsAccountIDKey)
	}
	apiKey := os.Getenv(env + rlsAPISecretKey)
	if apiKey == "" {
		return nil, fmt.Errorf("%s not set", env+rlsAPISecretKey)
	}
	// optionals
	webhookSecret := os.Getenv(env + rlsWebhookSecretKey)
	extraHeaders := parseExtraHeaders(make(map[string]string), os.Getenv(env+rlsHeadersKey))
	return rls.NewConfig(baseURL, apiKey, accountID, webhookSecret, extraHeaders), nil
}

func NewRLSClient(ctx context.Context, cliCtx *cli.Context) (*rls.RLSClient, error) {
	cfg, err := LoadRLSConfig()
	if err != nil {
		return nil, fmt.Errorf(msgFailedToLoadConfig, err)
	}
	if cliCtx.GlobalIsSet(flagHeaders) {
		cfg.ExtraHeaders = parseExtraHeaders(cfg.ExtraHeaders, cliCtx.GlobalString(flagHeaders))
	}
	httpClient := loadTLS(cliCtx)
	return rls.NewRLSClient(ctx, *cfg, httpClient), nil
}

func parseExtraHeaders(headerMap map[string]string, headerStr string) map[string]string {
	// in case of outOfBounds panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("error reading RLS headers: %v", r)
		}
	}()
	headers := strings.Split(headerStr, ",")
	for _, header := range headers {
		kv := strings.Split(header, ":")
		if len(kv) == 2 {
			headerMap[kv[0]] = kv[1]
		}
	}
	return headerMap
}
