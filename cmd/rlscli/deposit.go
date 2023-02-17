package main

import (
	"context"
	"fmt"
	"strconv"

	cli "github.com/urfave/cli"
)

var newInvoice = cli.Command{
	Name:      "newinvoice",
	Category:  "Deposits",
	Usage:     "Requests a new invoice from RLS",
	ArgsUsage: "amt [label] [network]",
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:     flagAmt,
			Usage:    "Amount of intended deposit in sats.",
			Required: false,
		},
		cli.StringFlag{
			Name:     flagLabel,
			Usage:    "Label (aka memo) for the deposit invoice.",
			Required: false,
		},
		cli.StringFlag{
			Name:     flagNetwork,
			Usage:    "Network (defaults to LN)",
			Required: false,
		},
	},
	Description: `
	Requests a new invoice from RLS.`,
	Action: cliNewInvoice,
}

func cliNewInvoice(ctx *cli.Context) {
	client, err := NewRLSClient(context.Background(), ctx)
	if err != nil {
		errFailedToCreateRLSClient(err)
		return
	}

	args := ctx.Args()

	var amount int64
	var label, network string

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
		fmt.Printf("amount in sats (--%s) must be provided\n", flagAmt)
		return
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
		return
	}
	printDepositInvoice(invoice)
}

var getInvoice = cli.Command{
	Name:      "getinvoice",
	Category:  "Deposits",
	Usage:     "Queries an invoice based on the invoice_id",
	ArgsUsage: flagInvoiceID,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     flagInvoiceID,
			Usage:    "Invoice ID to Query.",
			Required: false,
		},
	},
	Description: `
	Queries an invoice based on the invoice_id.
	`,
	Action: cliGetInvoice,
}

func cliGetInvoice(ctx *cli.Context) {
	client, err := NewRLSClient(context.Background(), ctx)
	if err != nil {
		errFailedToCreateRLSClient(err)
		return
	}

	args := ctx.Args()

	var invID string

	if ctx.IsSet(flagInvoiceID) {
		invID = ctx.String(flagInvoiceID)
	} else if args.Present() {
		invID = args.First()
		if invID == "" {
			fmt.Printf("%s must be set\n", flagInvoiceID)
			return
		}
	}

	invoice, err := client.GetInvoice(invID)
	if err != nil {
		fmt.Printf("Error GetInvoice: %s\n", err.Error())
		return
	}
	if invoice == nil {
		fmt.Printf("Error GetInvoice: invoice not returned\n")
		return
	}
	printDepositInvoice(invoice)
}

var getDeposit = cli.Command{
	Name:      "getdeposit",
	Category:  "Deposits",
	Usage:     "Queries a deposit based on the deposit_id",
	ArgsUsage: flagDepositID,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     flagDepositID,
			Usage:    "Deposit ID to Query.",
			Required: false,
		},
	},
	Description: `
	Queries a payment based on the deposit_id.
	`,
	Action: cliGetDeposit,
}

func cliGetDeposit(ctx *cli.Context) {
	client, err := NewRLSClient(context.Background(), ctx)
	if err != nil {
		errFailedToCreateRLSClient(err)
		return
	}

	args := ctx.Args()

	var depID string

	if ctx.IsSet(flagDepositID) {
		depID = ctx.String(flagDepositID)
	} else if args.Present() {
		depID = args.First()
		if depID == "" {
			fmt.Printf("deposit_id must be set\n")
			return
		}
	}

	dep, err := client.GetDeposit(depID)
	if err != nil {
		fmt.Printf("Error GetDeposit: %s\n", err.Error())
		return
	}
	if dep == nil {
		fmt.Printf("Error GetDeposit: deposit not returned\n")
		return
	}
	printDeposit(dep)
}

var listDeposits = cli.Command{
	Name:      "listdeposits",
	Category:  "Deposits",
	Usage:     "Queries list of settled deposits",
	ArgsUsage: fmt.Sprintf("%s %s", flagLimit, flagNextTimestamp),
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:     flagLimit,
			Usage:    "Number of results to return per page (1..25)",
			Required: false,
			Value:    25,
		},
		cli.Int64Flag{
			Name:     flagNextTimestamp,
			Usage:    "UNIX timestamp of next deposit to return",
			Required: false,
		},
	},
	Description: `
	Queries a payment based on the deposit_id.
	`,
	Action: cliListDeposits,
}

func cliListDeposits(ctx *cli.Context) {
	client, err := NewRLSClient(context.Background(), ctx)
	if err != nil {
		errFailedToCreateRLSClient(err)
		return
	}

	var limit int64 = 25
	var nextTimestamp int64

	args := ctx.Args()

	if ctx.IsSet(flagLimit) {
		limit = ctx.Int64(flagLimit)
	} else if args.Present() {
		limit, err = strconv.ParseInt(args.First(), 10, 64)
		fmt.Printf("read limit from args: %d", limit)
		if err != nil {
			fmt.Printf("unable to parse limit as int64: %s\n", err.Error())
			return
		}
		args = args.Tail()
	}

	if ctx.IsSet(flagNextTimestamp) {
		nextTimestamp = ctx.Int64(flagNextTimestamp)
	} else if args.Present() {
		nextTimestamp, err = strconv.ParseInt(args.First(), 10, 64)
		if err != nil {
			fmt.Printf("unable to parse nextTimestamp as int64: %s\n", err.Error())
			return
		}
	}

	deps, err := client.GetDeposits(limit, nextTimestamp)
	if err != nil {
		fmt.Printf("failed to list deposits: %s\n", err.Error())
		return
	}
	printDepositList(deps)
}
