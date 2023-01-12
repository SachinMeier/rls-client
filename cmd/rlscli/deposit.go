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

func cliNewInvoice(ctx *cli.Context) error {
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

func cliListDeposits(ctx *cli.Context) error {
	client, err := NewRLSClient(context.Background())
	if err != nil {
		return err
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
			return fmt.Errorf("unable to parse limit as int64 : %w", err)
		}
		args = args.Tail()
	}

	if ctx.IsSet(flagNextTimestamp) {
		nextTimestamp = ctx.Int64(flagNextTimestamp)
	} else if args.Present() {
		nextTimestamp, err = strconv.ParseInt(args.First(), 10, 64)
		if err != nil {
			return fmt.Errorf("unable to parse nextTimestamp as int64 : %w", err)
		}
	}

	deps, err := client.GetDeposits(limit, nextTimestamp)
	if err != nil {
		return fmt.Errorf("failed to list deposits : %w", err)
	}
	printDepositList(deps)
	return nil
}
