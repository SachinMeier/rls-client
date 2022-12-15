package main

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli"
)

const (
	flagAmt          = "amt"
	flagLabel        = "label"
	flagNetwork      = "network"
	flagInvoice      = "invoice"
	flagFeeLimit     = "fee_limit"
	flagWithdrawalID = "withdrawal_id"
	flagDepositID    = "deposit_id"
	flagCurrency     = "currency"
	flagURL          = "url"

	networkLN = "LN"
)

func main() {
	app := cli.NewApp()
	app.Name = "rlscli"
	app.Usage = "River Financial's Enterprise Lightning API"
	app.Commands = []cli.Command{
		getAccountSummary,
		newDepositInvoice,
		getDeposit,
		initiateWithdrawal,
		getWithdrawal,
		newWebhook,
		rmWebhook,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Print(err.Error())
	}
}

var getAccountSummary = cli.Command{
	Name:     "getaccount",
	Category: "Account",
	Usage:    "Returns a summary of your account",
	Description: `
	Returns a summary of your account.`,
	Action: cliGetAccountSummary,
}

var newDepositInvoice = cli.Command{
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
	Action: cliCreateDepositInvoice,
}

var initiateWithdrawal = cli.Command{
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
	Action: cliInitiateWithdrawal,
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

var newWebhook = cli.Command{
	Name:      "newwebhook",
	Category:  "Webhooks",
	Usage:     "Registers a new webhook if none exists",
	ArgsUsage: flagURL,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     flagURL,
			Usage:    "Webhook URL",
			Required: false,
		},
	},
	Description: `
	Registers a new webhook if none exists.
	`,
	Action: cliNewWebhook,
}

var rmWebhook = cli.Command{
	Name:      "rmwebhook",
	Category:  "Webhooks",
	Usage:     "Deletes the existing webhook",
	ArgsUsage: flagURL,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     flagURL,
			Usage:    "Webhook URL",
			Required: false,
		},
	},
	Description: `
	Deletes the existing webhook.
	`,
	Action: cliDeleteWebhook,
}
