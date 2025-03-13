package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/dapplink-labs/fee-estimator/common/cliapp"
	"github.com/dapplink-labs/fee-estimator/config"
	flags2 "github.com/dapplink-labs/fee-estimator/flags"
	"github.com/dapplink-labs/fee-estimator/services"
)

func runRpc(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	fmt.Println("running grpc services...")
	cfg := config.NewConfig(ctx)
	return services.NewFeeRpcService(&cfg)
}

func NewCli() *cli.App {
	flags := flags2.Flags
	return &cli.App{
		Version:              "0.0.1",
		Description:          "An fee estimator services with rpc",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "rpc",
				Flags:       flags,
				Description: "Run fee rpc services",
				Action:      cliapp.LifecycleCmd(runRpc),
			},
			{
				Name:        "version",
				Description: "Show project version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
