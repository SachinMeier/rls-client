package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/SachinMeier/rls-client"
	cli "github.com/urfave/cli"
)

func cliGetAccountSummary(ctx *cli.Context) error {
	client, err := NewRLSClient(context.Background())
	if err != nil {
		return err
	}

	acct, err := client.GetAccount()
	if err != nil {
		fmt.Printf("Error cliGetAccountSummary: %s\n", err.Error())
		return err
	}
	printAccountSummary(acct)
	return nil
}

func cliCreateDepositInvoice(ctx *cli.Context) error {
	client, err := NewRLSClient(context.Background())
	if err != nil {
		return err
	}

	args := ctx.Args()

	var amount int64
	var label, network string

	if ctx.IsSet(flagAmt) {
		amount = ctx.Int64(flagAmt)
	} else if args.Present() {
		amount, err = strconv.ParseInt(args.First(), 10, 64)
		if err != nil {
			return fmt.Errorf("invalid amount : %w", err)
		}
		args = args.Tail()
	} else {
		return fmt.Errorf("amount in sats (--%s) must be provided", flagAmt)
	}

	if ctx.IsSet(flagLabel) {
		label = ctx.String(flagLabel)
	}

	if ctx.IsSet(flagNetwork) {
		network = ctx.String(flagNetwork)
	} else {
		network = networkLN
	}

	invoice, err := client.NewInvoice(amount, label, network)
	if err != nil {
		fmt.Printf("Error NewInvoice: %s\n", err.Error())
		return err
	}
	printDepositInvoice(invoice)
	return nil
}

func cliInitiateWithdrawal(ctx *cli.Context) error {
	client, err := NewRLSClient(context.Background())
	if err != nil {
		return err
	}

	args := ctx.Args()

	var amount, feeLimit int64
	var invoice string

	if ctx.IsSet(flagInvoice) {
		invoice = ctx.String(flagInvoice)
	} else if args.Present() {
		invoice = args.First()
		args = args.Tail()
	} else {
		return fmt.Errorf("invoice must be set or passed as first argument")
	}

	if ctx.IsSet(flagAmt) {
		amount = ctx.Int64(flagAmt)
	} else if args.Present() {
		amount, err = strconv.ParseInt(args.First(), 10, 64)
		if err != nil {
			return fmt.Errorf("invalid amount : %w", err)
		}
		args = args.Tail()
	} else {
		return fmt.Errorf("amount in sats must be provided")
	}

	if ctx.IsSet(flagFeeLimit) {
		feeLimit = ctx.Int64(flagFeeLimit)
	} else if args.Present() {
		feeLimit, err = strconv.ParseInt(args.First(), 10, 64)
		if err != nil {
			return fmt.Errorf("invalid fee_limit : %w", err)
		}
		args = args.Tail()
	} else {
		feeLimit = rls.DefaultFeeLimit
	}

	wd := rls.NewWithdrawalWithFeeLimit(amount, invoice, feeLimit)

	withdrawal, err := client.NewWithdrawal(wd)
	if err != nil {
		fmt.Printf("Error cliInitiateWithdrawal: %s\n", err.Error())
		return err
	}
	printWithdrawal(withdrawal)
	return nil
}

func cliGetWithdrawal(ctx *cli.Context) error {
	client, err := NewRLSClient(context.Background())
	if err != nil {
		return err
	}

	var wdID string

	if ctx.IsSet(flagWithdrawalID) {
		wdID = ctx.String(flagWithdrawalID)
	} else {
		wdID = ctx.Args().First()
		if wdID == "" {
			return fmt.Errorf("withdrawal_id must be set")
		}
	}

	wd, err := client.GetWithdrawal(wdID)
	if err != nil {
		fmt.Printf("Error cliGetWithdrawal: %s\n", err.Error())
		return err
	}
	printWithdrawal(wd)
	return nil
}

func cliGetDeposit(ctx *cli.Context) error {
	client, err := NewRLSClient(context.Background())
	if err != nil {
		return err
	}

	args := ctx.Args()

	var depID string

	if ctx.IsSet(flagDepositID) {
		depID = ctx.String(flagDepositID)
	} else if args.Present() {
		depID = args.First()
		if depID == "" {
			return fmt.Errorf("deposit_id must be set")
		}
	}

	dep, err := client.GetDeposit(depID)
	if err != nil {
		fmt.Printf("Error cliGetDeposit: %s\n", err.Error())
		return err
	}
	if dep == nil {
		fmt.Printf("deposit not returned")
		return nil
	}
	printDeposit(dep)
	return nil
}

func cliNewWebhook(ctx *cli.Context) error {
	client, err := NewRLSClient(context.Background())
	if err != nil {
		return err
	}

	args := ctx.Args()

	var url string

	if ctx.IsSet(flagURL) {
		url = ctx.String(flagURL)
	} else if args.Present() {
		url = args.First()
	} else {
		return fmt.Errorf("url flag must be set or argument must be passed")
	}

	wh, err := client.SubscribeToWebhook(url)
	if err != nil {
		return fmt.Errorf("failed to subscribe to webhook : %w", err)
	}
	fmt.Printf("Subscribed to webhook!\nURL: %s\nSecret: %s\n", wh.URL, wh.Secret)
	return nil
}

func cliDeleteWebhook(ctx *cli.Context) error {
	client, err := NewRLSClient(context.Background())
	if err != nil {
		return err
	}

	args := ctx.Args()

	var url string

	if ctx.IsSet(flagURL) {
		url = ctx.String(flagURL)
	} else if args.Present() {
		url = args.First()
	} else {
		return fmt.Errorf("url flag must be set or argument must be passed")
	}

	err = client.DeleteWebhook(url)
	if err != nil {
		return fmt.Errorf("failed to subscribe to webhook : %w", err)
	}
	fmt.Printf("successfully deleted webhook\n")
	return nil
}
