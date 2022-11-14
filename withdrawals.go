package rls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// WithdrawalDetail is a portion of Withdrawal object
type WithdrawalDetail struct {
	Network  string `json:"network"`
	Invoice  string `json:"destination"`
	FeeLimit int64  `json:"fee_limit"`
}

// Withdrawal contains the result of a call that returns a withdrawal
type Withdrawal struct {
	Amount   int64            `json:"amount"`
	Currency string           `json:"currency"`
	Details  WithdrawalDetail `json:"withdrawal_details"`
	State    string           `json:"state,omitempty"`
	ID       string           `json:"id,omitempty"`
}

// Invoice returns Withdrawal Invoice string
func (wd *Withdrawal) Invoice() string {
	return wd.Details.Invoice
}

// Network returns Deposit Invoice Network
func (wd *Withdrawal) Network() string {
	return wd.Details.Network
}

// FeeLimit returns Withdrawal Detail's Fee Limit
func (wd *Withdrawal) FeeLimit() int64 {
	return wd.Details.FeeLimit
}

const (
	// LN is the default and only valid network
	LN string = "LN"
	// BTC is the default and only currency
	BTC string = "BTC"
	// DefaultFeeLimit is the default fee limit of this client
	DefaultFeeLimit int64 = 300
)

func (rls *RLSClient) handleWithdrawal(req *http.Request, err error) (*Withdrawal, error) {
	if err != nil {
		return nil, err
	}

	var withdrawal Withdrawal
	err = rls.sendRequest(req, &withdrawal)
	if err != nil {
		return nil, err
	}
	return &withdrawal, nil
}

// NewWithdrawal returns a Withdrawal object to be passed to SubmitWithdrawal
func NewWithdrawal(amount int64, invoice string) *Withdrawal {
	return &Withdrawal{
		Amount:   amount,
		Currency: BTC,
		Details: WithdrawalDetail{
			Invoice:  invoice,
			FeeLimit: DefaultFeeLimit,
			Network:  LN,
		},
	}
}

// NewWithdrawalWithFeeLimit returns a Withdrawal object with a defined fee_limit to be passed to SubmitWithdrawal
func NewWithdrawalWithFeeLimit(amount int64, invoice string, feeLimit int64) *Withdrawal {
	return &Withdrawal{
		Amount:   amount,
		Currency: BTC,
		Details: WithdrawalDetail{
			Invoice:  invoice,
			FeeLimit: feeLimit,
			Network:  LN,
		},
	}
}

// NewWithdrawal initiates a withdrawal from RLS API by paying a specific invoice
func (rls *RLSClient) NewWithdrawal(withdrawal *Withdrawal) (*Withdrawal, error) {
	body, err := json.Marshal(withdrawal)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/accounts/%s/withdrawals", rls.BaseURL(), rls.AccountID()), bytes.NewBuffer(body))

	return rls.handleWithdrawal(req, err)
}

// GetWithdrawal returns a withdrawal based on the passed withdrawal_id
func (rls *RLSClient) GetWithdrawal(withdrawalID string) (*Withdrawal, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/accounts/%s/withdrawals/%s",
			rls.BaseURL(),
			rls.AccountID(),
			withdrawalID),
		nil,
	)
	return rls.handleWithdrawal(req, err)
}
