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

func cliGetAccount(ctx *cli.Context) error {
	client, err := NewRLSClient(context.Background())
	if err != nil {
		return err
	}

	acct, err := client.GetAccount()
	if err != nil {
		fmt.Printf("Error cliGetAccount: %s\n", err.Error())
		return err
	}
	printAccount(acct)
	return nil
}
