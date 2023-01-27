package main

import (
	"context"
	"fmt"

	cli "github.com/urfave/cli"
)

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

func cliNewWebhook(ctx *cli.Context) {
	client, err := NewRLSClient(context.Background(), ctx)
	if err != nil {
		errFailedToCreateRLSClient(err)
		return
	}

	args := ctx.Args()

	var url string

	if ctx.IsSet(flagURL) {
		url = ctx.String(flagURL)
	} else if args.Present() {
		url = args.First()
	} else {
		fmt.Printf("url flag must be set or argument must be passed\n")
	}

	webhook, err := client.SubscribeToWebhook(url)
	if err != nil {
		fmt.Printf("failed to subscribe to webhook: %s\n", err.Error())
		return
	}
	printWebhook(webhook)
}

var getWebhook = cli.Command{
	Name:     "getwebhook",
	Category: "Webhooks",
	Usage:    "Fetches current webhook subscription if one exists",
	Description: `
	Fetches current webhook subscription if one exists.
	`,
	Action: cliGetWebhook,
}

func cliGetWebhook(ctx *cli.Context) {
	client, err := NewRLSClient(context.Background(), ctx)
	if err != nil {
		errFailedToCreateRLSClient(err)
		return
	}

	webhook, err := client.GetSubscribedWebhook()
	if err != nil {
		fmt.Printf("failed to get webhook: %s\n", err.Error())
		return
	}
	printWebhook(webhook)
}

var rmWebhook = cli.Command{
	Name:      "rmwebhook",
	Category:  "Webhooks",
	Usage:     "Deletes the existing webhook",
	ArgsUsage: flagURL,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     flagTLSPath,
			Usage:    "if set, loads TLS key and cert from <tlsPath>.key and <tlsPath>.cert and uses them in the HTTPS request",
			Required: false,
		},
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

func cliDeleteWebhook(ctx *cli.Context) {
	client, err := NewRLSClient(context.Background(), ctx)
	if err != nil {
		errFailedToCreateRLSClient(err)
		return
	}

	args := ctx.Args()

	var url string

	if ctx.IsSet(flagURL) {
		url = ctx.String(flagURL)
	} else if args.Present() {
		url = args.First()
	} else {
		fmt.Printf("url flag must be set or argument must be passed\n")
		return
	}

	err = client.DeleteWebhook(url)
	if err != nil {
		fmt.Printf("failed to delete webhook: %s\n", err.Error())
		return
	}
	fmt.Printf("successfully deleted webhook\n")
}
