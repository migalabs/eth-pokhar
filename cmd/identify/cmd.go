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
		&cli.BoolFlag{
			Name:  "recreate-table",
			Usage: "Recreate the t_identified_validators table, meant to be used when one of the methodologies of identification changes",
		},
		// &cli.IntFlag{
		// 	Name:        "init-slot",
		// 	Usage:       "Slot from where to start the backfill",
		// 	EnvVars:     []string{"ANALYZER_INIT_SLOT"},
		// 	DefaultText: "0",
		// },
		// &cli.IntFlag{
		// 	Name:        "final-slot",
		// 	Usage:       "Slot from where to finish the backfill",
		// 	EnvVars:     []string{"ANALYZER_FINAL_SLOT"},
		// 	DefaultText: "0",
		// },
		// &cli.IntFlag{
		// 	Name:        "workers-num",
		// 	Usage:       "Number of workers to process validators",
		// 	EnvVars:     []string{"ANALYZER_WORKER_NUM"},
		// 	DefaultText: "4",
		// },
		// &cli.StringFlag{
		// 	Name:        "download-mode",
		// 	Usage:       "Either backfill specified slots or follow the chain head example: hybrid,historical,finalized",
		// 	EnvVars:     []string{"ANALYZER_DOWNLOAD_MODE"},
		// 	DefaultText: "finalized",
		// },
		// &cli.StringFlag{
		// 	Name:        "metrics",
		// 	Usage:       "Metrics to be persisted to the database: epoch,block,rewards,transactions",
		// 	EnvVars:     []string{"ANALYZER_METRICS"},
		// 	DefaultText: "epoch,block",
		// },
		// &cli.IntFlag{
		// 	Name:        "prometheus-port",
		// 	Usage:       "Port on which to expose prometheus metrics",
		// 	EnvVars:     []string{"ANALYZER_PROMETHEUS_PORT"},
		// 	DefaultText: "9080",
		// }
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
		identifyRunner.Close()

	case <-procDoneC:
		logCmdIdentify.Info("Process successfully finished!")
	}
	close(sigtermC)
	close(procDoneC)

	return nil
}
