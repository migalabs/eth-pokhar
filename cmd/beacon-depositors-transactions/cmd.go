package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	beaconDepTx "github.com/migalabs/eth-pokhar/beacon-depositors-transactions"
	"github.com/migalabs/eth-pokhar/config"
	"github.com/migalabs/eth-pokhar/utils"
	"github.com/sirupsen/logrus"

	cli "github.com/urfave/cli/v2"
)

var BeaconDepositorsTransactionsCommand = &cli.Command{
	Name:   "beacon_depositors_transactions",
	Usage:  "fetches the transactions of the depositors of the beaconchain contract",
	Action: LaunchBeaconDepositorsTransactions,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "el-endpoint",
			Usage:       "Execution node endpoint",
			EnvVars:     []string{"EL_ENDPOINT"},
			DefaultText: "http://localhost:8545",
		},
		&cli.StringFlag{
			Name:        "db-url",
			Usage:       "Database where to store transactions",
			EnvVars:     []string{"DB_URL"},
			DefaultText: "postgres://user:password@localhost:5432/goteth",
		},
		&cli.StringFlag{
			Name:        "log-level",
			Usage:       "Log level: debug, warn, info, error",
			EnvVars:     []string{"LOG_LEVEL"},
			DefaultText: "info",
		},
		&cli.IntFlag{
			Name:        "db-workers-num",
			Usage:       "Number of workers to process database operations",
			EnvVars:     []string{"DB_WORKER_NUM"},
			DefaultText: "4",
		},
		&cli.StringFlag{
			Name:        "alchemy-url",
			Usage:       "Alchemy url",
			EnvVars:     []string{"ALCHEMY_URL"},
			DefaultText: "https://eth-mainnet.g.alchemy.com/v2/KEY",
		},
	},
}

var logCmdBeaconDepTx = logrus.WithField(
	"module", "beaconDepositorsTransactionsCommand",
)

var QueryTimeout = 90 * time.Second

func LaunchBeaconDepositorsTransactions(c *cli.Context) error {

	conf := config.NewBeaconDepositorsTransactionsConfig()
	conf.Apply(c)

	logrus.SetLevel(utils.ParseLogLevel(conf.LogLevel))

	// generate the beaconDepositorsTransactions
	beaconDepositorsTransactions, err := beaconDepTx.NewBeaconDepositorsTransactions(c.Context, conf)
	if err != nil {
		return err
	}

	procDoneC := make(chan struct{})
	sigtermC := make(chan os.Signal, 1)

	signal.Notify(sigtermC, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)

	go func() {
		beaconDepositorsTransactions.Run()
		procDoneC <- struct{}{}
	}()

	select {
	case <-sigtermC:
		logCmdBeaconDepTx.Info("Sudden shutdown detected, controlled shutdown of the cli triggered")
		beaconDepositorsTransactions.Stop()

	case <-procDoneC:
		logCmdBeaconDepTx.Info("Process successfully finished!")
	}
	close(sigtermC)
	close(procDoneC)

	return nil
}
