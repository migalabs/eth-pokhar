package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"

	cmd "github.com/migalabs/eth-pokhar/cmd/beacon-depositors-transactions"
	"github.com/migalabs/eth-pokhar/utils"
)

var (
	Version = "v0.0.1"
	CliName = "eth-pokhar"
	log     = logrus.WithField(
		"cli", "CliName",
	)
)

func main() {
	fmt.Println(CliName, Version)

	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true

	// Set the general log configurations for the entire tool
	// logrus.SetFormatter(utils.ParseLogFormatter("text"))
	logrus.SetFormatter(customFormatter)
	logrus.SetOutput(utils.ParseLogOutput("terminal"))
	logrus.SetLevel(utils.ParseLogLevel("info"))

	app := &cli.App{
		Name:  CliName,
		Usage: "Repository to identify Ethereum entities and their belonging validators",
		// UsageText: "eth [commands] [arguments...]",

		EnableBashCompletion: true,
		Commands: []*cli.Command{
			cmd.BeaconDepositorsTransactionsCommand,
		},
	}

	// generate the crawler
	if err := app.RunContext(context.Background(), os.Args); err != nil {
		log.Errorf("error: %v\n", err)
		os.Exit(1)
	}
}
