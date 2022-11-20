package rls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Webhook API endpoints

// Webhook contains the response from a webhook call
type Webhook struct {
	URL     string `json:"url"`
	Secret  string `json:"secret" default:""`
	Enabled bool   `json:"enabled"`
}

// SubscribeToWebhook subscribes to a webhook
func (rls *RLSClient) SubscribeToWebhook(callbackURL string) (*Webhook, error) {
	data := map[string]string{
		"url": callbackURL,
	}
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/accounts/%s/webhooks", rls.BaseURL(), rls.AccountID())
	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(body),
	)
	return rls.handleWebhookRequest(req, err)
}

// GetSubscribedWebhook queries subscribed webhook
func (rls *RLSClient) GetSubscribedWebhook() (*Webhook, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/accounts/%s/webhooks", rls.BaseURL(), rls.AccountID()),
		nil,
	)
	return rls.handleWebhookRequest(req, err)
}

// DeleteWebhook deletes the existing webhook
func (rls *RLSClient) DeleteWebhook() error {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/accounts/%s/webhooks", rls.BaseURL(), rls.AccountID()),
		nil,
	)
	if err != nil {
		return err
	}

	return rls.sendRequest(req, nil)
}

func (rls *RLSClient) handleWebhookRequest(req *http.Request, err error) (*Webhook, error) {
	if err != nil {
		return nil, err
	}

	var webhook Webhook
	err = rls.sendRequest(req, &webhook)
	if err != nil {
		return nil, err
	}
	return &webhook, nil
}

// Webhook Events

const (
	WebhookTypeDeposit    string = "DEPOSIT"
	WebhookTypeWithdrawal string = "WITHDRAWAL"

	WebhookStateSuccess string = "SUCCESS"
	WebhookStatePending string = "PENDING"
	WebhookStateFail    string = "FAIL"
)

// RLSWebhookMessage represents a webhook event sent from RLS
type RLSWebhookMessage struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	State string `json:"state"`
}
