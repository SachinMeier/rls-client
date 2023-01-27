package main

import (
	"context"
	"fmt"

	cli "github.com/urfave/cli"
)

var getAccount = cli.Command{
	Name:     "getaccount",
	Category: "Account",
	Usage:    "Returns a summary of your account",
	Description: `
	Returns a summary of your account.`,
	Action: cliGetAccount,
}

func cliGetAccount(ctx *cli.Context) {
	client, err := NewRLSClient(context.Background(), ctx)
	if err != nil {
		errFailedToCreateRLSClient(err)
		return
	}

	acct, err := client.GetAccount()
	if err != nil {
		fmt.Printf("Error GetAccount: %s\n", err.Error())
		return
	}
	printAccount(acct)
}
