package main

import (
	"fmt"

	"github.com/SachinMeier/rls-client"
)

func printAccount(acct *rls.Account) {
	fmt.Printf("--- Account: %s ---\n", acct.ID)
	fmt.Printf("  Total Balance:     %d sats\n", acct.Balance)
	fmt.Printf("  Available Balance: %d sats\n", acct.AvailableBalance)
	fmt.Printf("  Reserved Balance:  %d sats\n", acct.GetReservedBalance())
	fmt.Printf("-----------------------------\n")
}

func printDepositInvoice(inv *rls.Invoice) {
	fmt.Printf("--- Deposit Invoice: %s ---\n", inv.ID)
	fmt.Printf("  Network:    %s\n", inv.Network)
	fmt.Printf("  Timestamp:  %d\n", inv.Timestamp)
	fmt.Printf("  Invoice: %s\n", inv.Invoice)
	fmt.Printf("-------------------------------------\n")
}

func printWithdrawal(wd *rls.Withdrawal) {
	fmt.Printf("----- Withdrawal: %s -----\n", wd.ID)
	fmt.Printf("  Currency/Network: %s/%s\n", wd.Currency, wd.Network())
	fmt.Printf("  State:            %s\n", wd.State)
	fmt.Printf("  Invoice: %s\n", wd.Invoice())
	fmt.Printf("  Fee Limit: %d\n", wd.FeeLimit())
	fmt.Printf("  Fee Paid: %d\n", wd.FeePaid)
	fmt.Printf("  Timestamp: %d\n", wd.Timestamp)
	fmt.Printf("-------------------------------------\n")
}

func printWithdrawalList(wds *rls.WithdrawalList) {
	for _, withdrawal := range wds.Withdrawals {
		printWithdrawal(&withdrawal)
	}
	fmt.Printf("Next Timestamp: %d\n", wds.NextTimestamp)
	fmt.Printf("-------------------------------------\n")
}

func printDeposit(dep *rls.Deposit) {
	fmt.Printf("--- Deposit: %s ---\n", dep.ID)
	fmt.Printf("  Amount:     %d\n", dep.Amount)
	fmt.Printf("  State:      %s\n", dep.State)
	fmt.Printf("  Network:    %s\n", dep.Detail.Network)
	fmt.Printf("  Timestamp:  %d\n", dep.Timestamp)
	fmt.Printf("  Invoice ID: %s\n", dep.Invoice.ID)
	fmt.Printf("  Invoice:    %s\n", dep.Invoice.Invoice)
	fmt.Printf("-------------------------------------\n")
}

func printDepositList(deps *rls.DepositList) {
	for _, deposit := range deps.Deposits {
		printDeposit(&deposit)
	}
	fmt.Printf("Next Timestamp: %d\n", deps.NextTimestamp)
	fmt.Printf("-------------------------------------\n")
}

func printWebhook(wh *rls.Webhook) {
	fmt.Printf("--- Webhook ---\n")
	fmt.Printf("  Enabled: %v\n", wh.Enabled)
	fmt.Printf("  URL:     %s\n", wh.URL)
	if wh.Secret != "" {
		fmt.Printf("  Secret:  %s\n", wh.Secret)
	}
	fmt.Printf("---------------\n")
}

func printInvoice(invoice *rls.DecodedInvoice) {
	fmt.Printf("--- Invoice ---\n")
	fmt.Printf("  Amount: %d\n", invoice.Amount)
	fmt.Printf("  Destination: %s\n", invoice.NodeID)
	fmt.Printf("  Memo: %s\n", invoice.Memo)
	// fmt.Printf("  Invoice: %s\n", invoice.Invoice)
	fmt.Printf("---------------\n")
}

func printFeeEstimate(feeEstimate *rls.FeeEstimate) {
	fmt.Printf("--- Fee Estimate ---\n")
	fmt.Printf("  Fee Estimate: %d\n", feeEstimate.Fee)
	fmt.Printf("  Amount: %d\n", feeEstimate.Amount)
	// fmt.Printf("  Invoice: %s\n", feeEstimate.Invoice)
	fmt.Printf("---------------\n")
}

func errFailedToCreateRLSClient(err error) {
	fmt.Printf("failed to load RLS client: %s\n", err.Error())
}
