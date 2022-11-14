package rls

import (
	"fmt"
	"net/http"
)

type CurrencyBalance struct {
	Currency     string `json:"currency"`
	Amount       int64  `json:"amount"`
	AmountOnHold int64  `json:"amount_on_hold"`
}

// Account contains the balances of an account
type Account struct {
	ID               string             `json:"id"`
	Balance          int64              `json:"balance,omitempty"`
	AvailableBalance int64              `json:"available_balance,omitempty"`
	CurrencyBalances []*CurrencyBalance `json:"currency_balances,omitempty"`
}

// GetReservedBalance returns the reserved balance of an account, calculated as Balance - AvailableBalance
func (as *Account) GetReservedBalance() int64 {
	return as.Balance - as.AvailableBalance
}

// GetAccount returns a  of the account's balance and available balance
func (rls *RLSClient) GetAccount() (Account, error) {
	var acct Account

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts/%s/", rls.BaseURL(), rls.AccountID()), nil)
	if err != nil {
		return acct, err
	}

	err = rls.sendRequest(req, &acct)
	if err != nil {
		return acct, err
	}
	return acct, nil
}
