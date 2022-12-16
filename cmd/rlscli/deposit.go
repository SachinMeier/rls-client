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
