package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/migalabs/eth-pokhar/config"
	"github.com/migalabs/eth-pokhar/identify"
	"github.com/migalabs/eth-pokhar/utils"
	"github.com/sirupsen/logrus"

	cli "github.com/urfave/cli/v2"
)

var IdentifyCommand = &cli.Command{
	Name:   "identify",
	Usage:  "Identify the pool in which validators are participating or the entity who operates the validators",
	Action: LaunchIdentify,
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
			DefaultText: "postgres://user:password@localhost:5432/dbName",
		},
		&cli.StringFlag{
			Name:        "log-level",
			Usage:       "Log level: debug, warn, info, error",
			EnvVars:     []string{"LOG_LEVEL"},
			DefaultText: "info",
		},
		&cli.StringFlag{
			Name:        "alchemy-url",
			Usage:       "Alchemy url",
			EnvVars:     []string{"ALCHEMY_URL"},
			DefaultText: "https://eth-mainnet.g.alchemy.com/v2/KEY",
		},
		&cli.IntFlag{
			Name:        "workers-num",
			Usage:       "Number of workers to process API requests",
			EnvVars:     []string{"WORKER_NUM"},
			DefaultText: "10",
		},
		&cli.IntFlag{
			Name:        "whale-threshold",
			Usage:       "Minimum number of validators to be considered a whale",
			EnvVars:     []string{"WHALE_THRESHOLD"},
			DefaultText: "100",
		},
		&cli.BoolFlag{
			Name:  "recreate-table",
			Usage: "Recreate the t_identified_validators table, meant to be used when one of the methodologies of identification changes",
		},
	},
}

var logCmdIdentify = logrus.WithField(
	"module", "identifyCommand",
)

var QueryTimeout = 90 * time.Second

func LaunchIdentify(c *cli.Context) error {

	conf := config.NewIdentifyConfig()
	conf.Apply(c)

	logrus.SetLevel(utils.ParseLogLevel(conf.LogLevel))

	// generate the identify
	identifyRunner, err := identify.NewIdentify(c.Context, conf)
	if err != nil {
		return err
	}

	procDoneC := make(chan struct{})
	sigtermC := make(chan os.Signal, 1)

	signal.Notify(sigtermC, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)

	go func() {
		identifyRunner.Run()
		procDoneC <- struct{}{}
	}()

	select {
	case <-sigtermC:
		logCmdIdentify.Info("Sudden shutdown detected, controlled shutdown of the cli triggered")
		identifyRunner.Stop()

	case <-procDoneC:
		logCmdIdentify.Info("Process successfully finished!")
	}
	close(sigtermC)
	close(procDoneC)

	return nil
}
