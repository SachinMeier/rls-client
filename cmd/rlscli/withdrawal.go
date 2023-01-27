package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/SachinMeier/rls-client"
	cli "github.com/urfave/cli"
)

var newWithdrawal = cli.Command{
	Name:      "newwithdrawal",
	Category:  "Withdrawals",
	Usage:     "Requests a payment to the specified invoice from RLS",
	ArgsUsage: "amt [label] [network]",
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:     flagAmt,
			Usage:    "Amount of intended withdrawal in sats.",
			Required: true,
		},
		cli.StringFlag{
			Name:     flagInvoice,
			Usage:    "BOLT-11 Invoice for RLS to pay",
			Required: true,
		},
		cli.StringFlag{
			Name:     flagFeeLimit,
			Usage:    "Fee Limit for the withdrawal.",
			Required: true,
		},
		cli.StringFlag{
			Name:     flagNetwork,
			Usage:    "Network (defaults to LN)",
			Required: false,
		},
		cli.StringFlag{
			Name:     flagCurrency,
			Usage:    "Currency (defaults to BTC)",
			Required: false,
		},
	},
	Description: `
	Requests a payment to the specified invoice from RLS.
	`,
	Action: cliNewWithdrawal,
}

func cliNewWithdrawal(ctx *cli.Context) {
	client, err := NewRLSClient(context.Background(), ctx)
	if err != nil {
		errFailedToCreateRLSClient(err)
		return
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
		fmt.Printf("invoice must be set or passed as first argument\n")
		return
	}

	if ctx.IsSet(flagAmt) {
		amount = ctx.Int64(flagAmt)
	} else if args.Present() {
		amount, err = strconv.ParseInt(args.First(), 10, 64)
		if err != nil {
			fmt.Printf("invalid amount: %s\n", err.Error())
			return
		}
		args = args.Tail()
	} else {
		fmt.Printf("amount in sats must be provided\n")
		return
	}

	if ctx.IsSet(flagFeeLimit) {
		feeLimit = ctx.Int64(flagFeeLimit)
	} else if args.Present() {
		feeLimit, err = strconv.ParseInt(args.First(), 10, 64)
		if err != nil {
			fmt.Printf("invalid fee_limit: %s\n", err.Error())
			return
		}
		args = args.Tail()
	} else {
		feeLimit = rls.DefaultFeeLimit
	}

	wd := rls.NewWithdrawalWithFeeLimit(amount, invoice, feeLimit)

	withdrawal, err := client.NewWithdrawal(wd)
	if err != nil {
		fmt.Printf("Error cliInitiateWithdrawal: %s\n", err.Error())
		return
	}
	printWithdrawal(withdrawal)
}

var getWithdrawal = cli.Command{
	Name:      "getwithdrawal",
	Category:  "Withdrawals",
	Usage:     "Queries a payment based on the withdrawal_id",
	ArgsUsage: flagWithdrawalID,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     flagWithdrawalID,
			Usage:    "Withdrawal ID to Query.",
			Required: false,
		},
	},
	Description: `
	Queries a payment based on the withdrawal_id.
	`,
	Action: cliGetWithdrawal,
}

func cliGetWithdrawal(ctx *cli.Context) {
	client, err := NewRLSClient(context.Background(), ctx)
	if err != nil {
		errFailedToCreateRLSClient(err)
		return
	}

	var wdID string

	if ctx.IsSet(flagWithdrawalID) {
		wdID = ctx.String(flagWithdrawalID)
	} else {
		wdID = ctx.Args().First()
		if wdID == "" {
			fmt.Printf("withdrawal_id must be set\n")
			return
		}
	}

	wd, err := client.GetWithdrawal(wdID)
	if err != nil {
		fmt.Printf("Error cliGetWithdrawal: %s\n", err.Error())
		return
	}
	printWithdrawal(wd)
}
