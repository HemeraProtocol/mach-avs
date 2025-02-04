package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/urfave/cli"

	"github.com/HemeraProtocol/avs/legacy/aggregator"
	"github.com/HemeraProtocol/avs/legacy/core/config"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string
)

func main() {
	app := cli.NewApp()
	app.Flags = config.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "mach-aggregator"
	app.Usage = "Mach Aggregator"
	app.Description = "Service that sends alerts by operator nodes."

	app.Action = aggregatorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func aggregatorMain(cliCtx *cli.Context) error {
	log.Println("Initializing Aggregator")
	config, err := config.NewConfig(cliCtx)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}
	fmt.Println("Config:", string(configJson))

	agg, err := aggregator.NewAggregator(config)
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	err = agg.Start(ctx, wg)
	if err != nil {
		log.Fatalln("Aggregator run failed", "err", err)
	}

	wg.Wait()

	return nil
}
