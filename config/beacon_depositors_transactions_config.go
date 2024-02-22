package config

import (
	cli "github.com/urfave/cli/v2"
)

type BeaconDepositorsTransactionsConfig struct {
	LogLevel     string `json:"log-level"`
	ElEndpoint   string `json:"el-endpoint"`
	DBUrl        string `json:"db-url"`
	Workers      int    `json:"workers-num"`
	AlchemyURL   string `json:"alchemy-url"`
	OnlyDeposits bool   `json:"only-deposits"`
}

func NewBeaconDepositorsTransactionsConfig() *BeaconDepositorsTransactionsConfig {
	// Return Default values for the ethereum configuration
	return &BeaconDepositorsTransactionsConfig{
		LogLevel:   DefaultLogLevel,
		DBUrl:      DefaultDBUrl,
		ElEndpoint: DefaultElEndpoint,
		Workers:    DefaultWorkers,
		AlchemyURL: DefaultAlchemyURL}
}

func (c *BeaconDepositorsTransactionsConfig) Apply(ctx *cli.Context) {
	// apply to the existing Default configuration the set flags
	// log level
	if ctx.IsSet("log-level") {
		c.LogLevel = ctx.String("log-level")
	}
	// el url
	if ctx.IsSet("el-endpoint") {
		c.ElEndpoint = ctx.String("el-endpoint")
	}
	// db url
	if ctx.IsSet("db-url") {
		c.DBUrl = ctx.String("db-url")
	}
	if ctx.IsSet("workers-num") {
		c.Workers = ctx.Int("workers-num")
	}
	if ctx.IsSet("alchemy-url") {
		c.AlchemyURL = ctx.String("alchemy-url")
	}

	if ctx.IsSet("only-deposits") {
		c.OnlyDeposits = true
	}

}
