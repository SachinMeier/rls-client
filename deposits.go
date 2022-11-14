package rls

import (
	"fmt"
	"net/http"
)

// Deposit contains a single Deposit object
type Deposit struct {
	ID        string        `json:"id"`
	Invoice   Invoice       `json:"deposit_intent"`
	Amount    int64         `json:"amount"`
	Detail    DepositDetail `json:"deposit_details"`
	State     string        `json:"state"`
	Timestamp int64         `json:"timestamp"`
}

// DepositDetail forms a part of a Deposit
type DepositDetail struct {
	Network string `json:"network"`
	Proof   string `json:"proof"`
}

// DepositList is a single page of paginated results from RLS API's GetDeposits call
type DepositList struct {
	Deposits      []Deposit `json:"deposits"`
	NextTimestamp int64     `json:"next_timestamp"`
}

// Count returns the number of deposits in a DepositList
func (dl *DepositList) Count() int {
	return len(dl.Deposits)
}

func (rls *RLSClient) GetDeposit(depositID string) (*Deposit, error) {
	url := fmt.Sprintf("%s/accounts/%s/deposits/%s", rls.BaseURL(), rls.AccountID(), depositID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get deposit : %w", err)
	}

	var deposit Deposit
	err = rls.sendRequest(req, &deposit)
	if err != nil {
		return nil, fmt.Errorf("failed to get deposit : %w", err)
	}
	return &deposit, nil
}

// GetDeposits returns a list of deposits (settled invoices) to RLS
func (rls *RLSClient) GetDeposits(limit, nextTimestamp int64) (*DepositList, error) {
	url := fmt.Sprintf("%s/accounts/%s/deposits", rls.BaseURL(), rls.AccountID())
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Add query params
	query := req.URL.Query()
	query.Add("limit", fmt.Sprint(limit))
	if nextTimestamp != 0 {
		query.Add("next_timestamp", fmt.Sprint(nextTimestamp))
	}
	req.URL.RawQuery = query.Encode()

	var deposits DepositList
	err = rls.sendRequest(req, &deposits)
	if err != nil {
		return nil, err
	}
	return &deposits, nil
}
