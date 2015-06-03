package main

import (
	"github.com/codegangsta/cli"
	"github.com/gloob/go-telegram/tg"

	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var (
	// cli.App for define our command line application.
	app = cli.NewApp()

	// Global configuration.
	globalConfig tg.Config

	// Channel for handling signals.
	sigs = make(chan os.Signal, 1)
)

func main() {
	// Load main configuration.
	tg.LoadConfig("./config.toml", &globalConfig)
	fmt.Println(globalConfig)

	// Set up the command line application.
	app.Name = tg.AppName
	app.Version = tg.AppVersion
	app.Authors = tg.AppAuthors
	app.Usage = tg.AppDescription

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "loglevel, l",
			Value: 10,
			Usage: "Log level.",
		},
		cli.StringFlag{
			Name:  "logname, L",
			Value: "",
			Usage: "Log filename.",
		},
	}
	app.Action = func(c *cli.Context) {
		println("Starting actions on context...")

		// Create, setup and run Terminal.
		term := tg.NewTerminal()
		defer term.Destroy()

		term.Init()

		term.Loop()
	}

	// Capture and manage signals.
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-sigs
		//cleanup()
		fmt.Println("Signal interruption.")
		fmt.Println(sig)
		os.Exit(1)
	}()

	// Run main app loop.
	fmt.Println("App Start.")
	app.Run(os.Args)
}
