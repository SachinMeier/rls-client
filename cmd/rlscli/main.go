package main

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli"
)

const (
	flagAmt           = "amt"
	flagLabel         = "label"
	flagNetwork       = "network"
	flagInvoice       = "invoice"
	flagFeeLimit      = "fee_limit"
	flagWithdrawalID  = "withdrawal_id"
	flagDepositID     = "deposit_id"
	flagCurrency      = "currency"
	flagURL           = "url"
	flagLimit         = "limit"
	flagNextTimestamp = "next"
	flagTLSPath       = "tlspath"
	flagHeaders       = "headers"

	networkLN = "LN"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     flagHeaders,
			Usage:    "[Optional] add extra headers to the requests. Headers should be in format key:value,key:value...",
			Required: false,
		},
		cli.StringFlag{
			Name:     flagTLSPath,
			Usage:    "if set, loads TLS key and cert from <tlsPath>.key and <tlsPath>.cert and uses them in the HTTPS request",
			Required: false,
		},
	}
	app.Name = "rlscli"
	app.Usage = "River Financial's Enterprise Lightning API"
	app.Commands = []cli.Command{
		getAccount,
		newInvoice,
		getDeposit,
		listDeposits,
		newWithdrawal,
		getWithdrawal,
		newWebhook,
		getWebhook,
		rmWebhook,
		parseInvoice,
		estimateLightningFee,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Print(err.Error())
	}
}
