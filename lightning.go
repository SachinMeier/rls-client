package rls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// DecodedInvoice contains the result of a call to decode invoice
type DecodedInvoice struct {
	Amount  int64  `json:"amount"`
	Memo    string `json:"memo"`
	NodeID  string `json:"node_id"`
	Invoice string `json:"destination"`
}

// FeeEstimate contains the result of a call to EstimateFee
type FeeEstimate struct {
	Amount  int64  `json:"amount"`
	Invoice string `json:"destination"`
	Fee     int64  `json:"fee"`
}

// DecodeInvoice decodes a Lightning Invoice using RLS using `lncli decodepayreq`
func (rls *RLSClient) DecodeInvoice(invoice string) (DecodedInvoice, error) {
	data := map[string]string{
		"destination": invoice,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return DecodedInvoice{}, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/lightning/parse_invoice", rls.BaseURL()),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return DecodedInvoice{}, err
	}

	var decodedInvoice DecodedInvoice
	err = rls.sendRequest(req, &decodedInvoice)
	if err != nil {
		return DecodedInvoice{}, err
	}
	return decodedInvoice, nil
}

// EstimateLightningFee estimates Lightning Fee of an invoice using `lncli`
func (rls *RLSClient) EstimateLightningFee(invoice string, amount int64) (FeeEstimate, error) {
	data := map[string]string{
		"destination": invoice,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return FeeEstimate{}, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/lightning/estimate_fee/", rls.BaseURL()),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return FeeEstimate{}, err
	}

	var feeEstimate FeeEstimate
	err = rls.sendRequest(req, &feeEstimate)
	if err != nil {
		return FeeEstimate{}, err
	}
	return feeEstimate, nil
}
