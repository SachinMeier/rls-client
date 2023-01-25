package main

import (
	"context"
	"fmt"
	"strconv"

	cli "github.com/urfave/cli"
)

var parseInvoice = cli.Command{
	Name:      "parseinvoice",
	Category:  "Lightning",
	Usage:     "Parses a BOLT-11 invoice",
	ArgsUsage: "invoice",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     flagInvoice,
			Usage:    "invoice to be parsed",
			Required: false,
		},
	},
	Description: `
	Parse a BOLT-11 invoice.
	`,
	Action: cliParseInvoice,
}

func cliParseInvoice(ctx *cli.Context) error {
	client, err := NewRLSClient(context.Background(), ctx)
	if err != nil {
		return err
	}

	args := ctx.Args()

	var invoice string

	if ctx.IsSet(flagInvoice) {
		invoice = ctx.String(flagInvoice)
	} else if args.Present() {
		invoice = args.First()
		args = args.Tail()
	} else {
		return fmt.Errorf("invoice must be set or passed as first argument")
	}

	decodedInvoice, err := client.DecodeInvoice(invoice)
	if err != nil {
		fmt.Printf("Error ParseInvoice: %s\n", err.Error())
		return err
	}
	printInvoice(decodedInvoice)
	return nil
}

var estimateLightningFee = cli.Command{
	Name:      "estimatefee",
	Category:  "Lightning",
	Usage:     "Estimates fee to send to an amount to a specific node",
	ArgsUsage: "invoice amount",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     flagInvoice,
			Usage:    "invoice to be parsed",
			Required: false,
		},
		cli.Int64Flag{
			Name:     flagAmt,
			Usage:    "invoice to be parsed",
			Required: false,
		},
	},
	Description: `
	Parse a BOLT-11 invoice.
	`,
	Action: cliEstimateLightningFee,
}

func cliEstimateLightningFee(ctx *cli.Context) error {
	client, err := NewRLSClient(context.Background(), ctx)
	if err != nil {
		return err
	}

	args := ctx.Args()

	var amount int64
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

	feeEstimate, err := client.EstimateLightningFee(invoice, amount)
	if err != nil {
		fmt.Printf("Error ParseInvoice: %s\n", err.Error())
		return err
	}
	printFeeEstimate(feeEstimate)
	return nil
}
