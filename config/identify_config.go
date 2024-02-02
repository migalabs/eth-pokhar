package config

import (
	cli "github.com/urfave/cli/v2"
)

type IdentifyConfig struct {
	LogLevel      string `json:"log-level"`
	ElEndpoint    string `json:"el-endpoint"`
	DBUrl         string `json:"db-url"`
	DBWorkers     int    `json:"db-workers-num"`
	AlchemyURL    string `json:"alchemy-url"`
	RecreateTable bool   `json:"recreate-table"`
}

func NewIdentifyConfig() *IdentifyConfig {
	// Return Default values for the ethereum configuration
	return &IdentifyConfig{
		LogLevel:   DefaultLogLevel,
		DBUrl:      DefaultDBUrl,
		ElEndpoint: DefaultElEndpoint,
		DBWorkers:  DefaultDBWorkers,
		AlchemyURL: DefaultAlchemyURL}
}

func (c *IdentifyConfig) Apply(ctx *cli.Context) {
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
	if ctx.IsSet("db-workers-num") {
		c.DBWorkers = ctx.Int("db-workers-num")
	}
	if ctx.IsSet("alchemy-url") {
		c.AlchemyURL = ctx.String("alchemy-url")
	}

	if ctx.Bool("recreate-table") {
		c.RecreateTable = true
	}

}
