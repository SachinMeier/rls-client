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

	networkLN = "LN"
)

func main() {
	app := cli.NewApp()
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
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Print(err.Error())
	}
}
